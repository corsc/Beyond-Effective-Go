package pduty

const (
	// this is the base URL for all API calls
	apiBaseURL = "https://api.pagerduty.com/"

	fetchLimit = 50
)

type apiResponse struct {
	UsersOuter []*userOuter `json:"users"`
	Limit      int
	Offset     int
	More       bool
}

type userOuter struct {
	ID             string
	Name           string
	Email          string
	Teams          []*team
	ContactMethods []*contactMethod `json:"contact_methods"`
}

type team struct {
	ID   string
	Name string
}

type contactMethod struct {
	ID      string
	Type    string
	Address string
}
