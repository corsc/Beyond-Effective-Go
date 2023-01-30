package pduty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// User is a response DTO
type User struct {
	ID   string
	Name string `json:"summary"`
}

// ScheduleEntry is a response DTO
type ScheduleEntry struct {
	Start time.Time
	End   time.Time
	User  *User
}

func (s *ScheduleEntry) String() string {
	return fmt.Sprintf("User: %s - Start: %s - End: %s", s.User.Name, s.Start.Format(time.RFC3339), s.End.Format(time.RFC3339))
}

// Schedule is a response DTO
type Schedule struct {
	Name    string
	Entries []*ScheduleEntry `json:"rendered_schedule_entries"`
}

// ScheduleAPI contains the functions to call the schedule APIs
type ScheduleAPI struct{}

// GetSchedule will return the schedule for the supplied id
func (s *ScheduleAPI) GetSchedule(apiKey string, scheduleID string, start time.Time, end time.Time) (*Schedule, error) {
	req, err := s.buildRequest(apiKey, scheduleID, start, end)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	apiResp := &apiResponse{}
	err = decoder.Decode(apiResp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response to JSON with err: %s", err)
	}

	return apiResp.ScheduleOuter.Schedule, nil
}

func (s *ScheduleAPI) buildRequest(apiKey string, scheduleID string, start time.Time, end time.Time) (*http.Request, error) {
	params := &url.Values{}
	params.Set("time_zone", "UTC")
	params.Set("since", start.Format(time.RFC3339))
	params.Set("until", end.Format(time.RFC3339))

	req, err := http.NewRequest("GET", apiBaseURL+"/schedules/"+scheduleID+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token token="+apiKey)
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	return req, nil
}

type scheduleOuter struct {
	Schedule *Schedule `json:"final_schedule"`
}
