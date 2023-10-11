package gcal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// CalendarItem is an output DTO
type CalendarItem struct {
	Start time.Time
	End   time.Time
}

// Calendar is an output DTO
type Calendar struct {
	Items []*CalendarItem
}

// CalendarAPI contains the functions to call the calendar APIs
type CalendarAPI struct{}

// GetCalendars returns the calendars for the emails (map values) provided
func (c *CalendarAPI) GetCalendars(credentialsFile, tokenFile string, users map[string]string, start time.Time, end time.Time) (map[string]*Calendar, error) {
	api, err := c.getAPI(credentialsFile, tokenFile)
	if err != nil {
		return nil, err
	}

	out := map[string]*Calendar{}

	for id, email := range users {
		calendar := &Calendar{}

		// check for "Out of Office" events
		err = c.getCalendar(api, calendar, "out", email, start, end)
		if err != nil {
			return nil, err
		}

		// check for "No On-Call" events
		err = c.getCalendar(api, calendar, "xoncall", email, start, end)
		if err != nil {
			return nil, err
		}

		out[id] = calendar
	}

	return out, nil
}

// will return the calendar for the supplied email address
// (taken from API example)
func (c *CalendarAPI) getCalendar(api *calendar.Service, out *Calendar, searchTerm string, email string, start time.Time, end time.Time) error {
	settings, err := api.Settings.Get("timezone").Do()
	if err != nil {
		return err
	}

	userTimezone := settings.Value
	location, err := time.LoadLocation(userTimezone)
	if err != nil {
		return err
	}

	events, err := api.Events.List(email).
		AlwaysIncludeEmail(false).
		ShowDeleted(false).
		SingleEvents(true).
		// Expand the search to ensure we get everything
		TimeMin(start.AddDate(0, 0, -1).Format(time.RFC3339)).
		TimeMax(end.AddDate(0, 0, 1).Format(time.RFC3339)).
		MaxResults(100).
		Q(searchTerm).
		Do()

	if err != nil {
		return err
	}

	if len(events.Items) == 0 {
		return nil
	}

	for _, item := range events.Items {
		startTime, err := c.getTime(item.Start.DateTime, item.Start.Date, location)
		if err != nil {
			return err
		}

		endTime, err := c.getTime(item.End.DateTime, item.End.Date, location)
		if err != nil {
			return err
		}

		out.Items = append(out.Items, &CalendarItem{Start: startTime, End: endTime})
	}

	return nil
}

func (x *CalendarAPI) getTime(dateTime string, date string, location *time.Location) (time.Time, error) {
	if dateTime != "" {
		out, err := time.ParseInLocation(time.RFC3339, dateTime, location)
		if err != nil {
			return time.Time{}, err
		}

		return out, nil
	}

	out, err := time.ParseInLocation("2006-01-02", date, location)
	if err != nil {
		return time.Time{}, err
	}

	return out, nil
}

func (c *CalendarAPI) getAPI(credsFile, tokFile string) (*calendar.Service, error) {
	b, err := os.ReadFile(credsFile)
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}
	client := getClient(tokFile, config)

	return calendar.New(client)
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(tokFile string, config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Fprintf(os.Stderr, "Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
