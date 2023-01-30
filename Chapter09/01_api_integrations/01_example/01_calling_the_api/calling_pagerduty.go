package _1_calling_the_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// User is the DTO for the user and forms part of the API for this package
type User struct {
	// User Details
	ID    string
	Name  string
	Email string

	// Current settings
	EmailIsSet bool
	PhoneIsSet bool
	SMSIsSet   bool
	PushIsSet  bool
}

func NewUsersAPI(apiBaseURL string, apiKey string) *UsersAPI {
	return &UsersAPI{
		apiBaseURL: apiBaseURL,
		apiKey:     apiKey,
	}
}

// UsersAPI contains the functions to call the Users API
type UsersAPI struct {
	apiBaseURL string
	apiKey     string
}

// GetUsers will returns the mapping between PD user id and email
func (u *UsersAPI) GetUsers(ctx context.Context) ([]*User, error) {
	rawUsers, err := u.fetchAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return u.convertUsers(rawUsers), nil
}

func (u *UsersAPI) fetchAllUsers(ctx context.Context) ([]*users, error) {
	req, err := u.buildRequest(ctx)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	apiResp := &apiResponse{}
	err = decoder.Decode(apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp.Users, nil
}

func (u *UsersAPI) buildRequest(ctx context.Context) (*http.Request, error) {
	params := &url.Values{}
	params.Set("include[]", "contact_methods")

	req, err := http.NewRequestWithContext(ctx, "GET", u.apiBaseURL+"/users?"+params.Encode(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token token="+u.apiKey)
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (u *UsersAPI) convertUsers(rawUsers []*users) []*User {
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
			ID:         rawUser.ID,
			Name:       rawUser.Name,
			Email:      rawUser.Email,
			EmailIsSet: emailSet,
			PhoneIsSet: phoneSet,
			SMSIsSet:   smsSet,
			PushIsSet:  pushSet,
		})
	}

	return out
}

type apiResponse struct {
	Users []*users `json:"users"`
}

type users struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	ContactMethods []*contactMethod `json:"contact_methods"`
}

type contactMethod struct {
	Type    string
	Address string
}
