package pduty

const (
	// this is the base URL for all API calls
	apiBaseURL = "https://api.pagerduty.com/"
)

type apiResponse struct {
	ScheduleOuter *scheduleOuter `json:"schedule"`
	UserOuter     *userOuter     `json:"user"`
}
