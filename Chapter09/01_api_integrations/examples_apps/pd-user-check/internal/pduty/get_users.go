package pduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// User is the DTO for the user and forms part of the API for this package
type User struct {
	ID       string
	Name     string
	Email    string
	Teams    int
	EmailSet bool
	PhoneSet bool
	SMSSet   bool
	PushSet  bool
}

// UsersAPI contains the functions to call the Users API
type UsersAPI struct{}

// GetUsers will returns the mapping between PD user id and email
func (u *UsersAPI) GetUsers(apiKey string, search string) ([]*User, error) {
	rawUsers, err := u.fetchAllUsers(apiKey, search)
	if err != nil {
		return nil, err
	}

	return u.convertUsers(rawUsers), nil
}

func (u *UsersAPI) fetchAllUsers(apiKey string, search string) ([]*userOuter, error) {
	hasMore := true
	offset := 0
	var out []*userOuter

	for hasMore {
		resp, err := u.getUserBatch(apiKey, search, offset)
		if err != nil {
			return nil, err
		}

		out = append(out, resp.UsersOuter...)

		// prepare for next iteration
		offset = resp.Offset + resp.Limit
		hasMore = resp.More
	}

	return out, nil
}

func (u *UsersAPI) getUserBatch(apiKey string, search string, offset int) (*apiResponse, error) {
	fmt.Printf("Fetching users %d to %d\n", offset+1, offset+fetchLimit)
	req, err := u.buildRequest(apiKey, search, offset)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	decoder := json.NewDecoder(resp.Body)

	apiResp := &apiResponse{}
	err = decoder.Decode(apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp, nil
}

func (u *UsersAPI) buildRequest(apiKey string, search string, offset int) (*http.Request, error) {
	params := &url.Values{}
	params.Set("offset", strconv.Itoa(offset))
	params.Set("limit", strconv.Itoa(fetchLimit))
	params.Set("query", search)
	params.Set("include[]", "contact_methods")
	params.Add("include[]", "teams")

	req, err := http.NewRequest("GET", apiBaseURL+"/users?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token token="+apiKey)
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	return req, nil
}

func (u *UsersAPI) convertUsers(rawUsers []*userOuter) []*User {
	var out []*User

	for _, rawUser := range rawUsers {
		emailSet := false
		phoneSet := false
		smsSet := false
		pushSet := false

		for _, method := range rawUser.ContactMethods {
			switch method.Type {
			case "email_contact_method":
				emailSet = len(method.Address) > 0

			case "phone_contact_method":
				phoneSet = len(method.Address) > 0

			case "sms_contact_method":
				smsSet = len(method.Address) > 0

			case "push_notification_contact_method":
				pushSet = len(method.Address) > 0
			}
		}

		out = append(out, &User{
			ID:       rawUser.ID,
			Name:     rawUser.Name,
			Email:    rawUser.Email,
			Teams:    len(rawUser.Teams),
			EmailSet: emailSet,
			PhoneSet: phoneSet,
			SMSSet:   smsSet,
			PushSet:  pushSet,
		})
	}

	return out
}
