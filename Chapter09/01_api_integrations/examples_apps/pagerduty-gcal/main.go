package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/conflict"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/gcal"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/pduty"
)

// NOTES:
// * For this tool to work it requires a public calendar event with the word "out" in the title.
// 	 Using the "out of office" event type in the google calendar UI will achieve this.

const (
	timeFormat = "2006-01-02 15:04"
)

var (
	scheduleID        string
	startAsString     string
	days              int64
	daysBetweenShifts int64
)

func main() {
	// these are inputs that should come from command line or env
	apiKey, found := os.LookupEnv("PD_API_KEY")
	if !found {
		panic("PD_API_KEY must be set")
	}
	flag.StringVar(&scheduleID, "schedule", "", "schedule id (see README.md) for more info")
	flag.StringVar(&startAsString, "start", "", "start of the schedule")
	flag.Int64Var(&days, "days", 30, "days to add to start to define the schedule")
	flag.Int64Var(&daysBetweenShifts, "between", 3, "minimum number of days between shifts")
	flag.Parse()

	periodStart, err := time.Parse("2006-01-02", startAsString)
	if err != nil {
		fmt.Printf("failed to parse start with err: %s\n", err)
		flag.PrintDefaults()
		return
	}

	now := time.Now()
	if periodStart.Before(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)) {
		fmt.Print("sorry you cannot re-write the past\n")
		return
	}

	end := periodStart.Add(time.Duration(days) * 24 * time.Hour)
	credentialsFile := "credentials.json"
	tokenFile := "token.json"

	// actual logic
	fmt.Printf("Loading schedule for %s to %s\n", periodStart.Format(timeFormat), end.Format(timeFormat))
	scheduleStart := periodStart.Add(time.Duration(-daysBetweenShifts*24) * time.Hour)
	schedule, err := (&pduty.ScheduleAPI{}).GetSchedule(apiKey, scheduleID, scheduleStart, end)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Loading scheduled user details\n")
	participants, err := (&pduty.UserAPI{}).GetUsers(apiKey, schedule.Entries)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Loading calendars for scheduled users\n")
	calendars, err := (&gcal.CalendarAPI{}).GetCalendars(credentialsFile, tokenFile, participants, periodStart, end)
	if err != nil {
		fmt.Print(err)
		return
	}

	conflicts := checkForConflicts(schedule, calendars, daysBetweenShifts)
	if len(conflicts) == 0 {
		return
	}

	_ = findSwaps(periodStart, schedule, conflicts, calendars)
}

func checkForConflicts(schedule *pduty.Schedule, calendars map[string]*gcal.Calendar, daysBetweenShifts int64) []*pduty.ScheduleEntry {
	fmt.Printf("Checking for conflicts\n")
	conflictsOrdered, err := (&conflict.CheckerAPI{}).Check(schedule, calendars, daysBetweenShifts)
	if err != nil {
		panic(err)
	}

	// output result
	if len(conflictsOrdered) == 0 {
		log.Printf("No conflicts found")
		return nil
	}

	fmt.Printf("Conflict (slot : user)\n")
	for _, scheduleEntry := range conflictsOrdered {
		fmt.Printf("%s to %s : %s\n", scheduleEntry.Start.Format(timeFormat), scheduleEntry.End.Format(timeFormat), scheduleEntry.User.Name)
	}

	return conflictsOrdered
}

func findSwaps(periodStart time.Time, schedule *pduty.Schedule, conflicts []*pduty.ScheduleEntry, calendars map[string]*gcal.Calendar) map[*pduty.ScheduleEntry]*pduty.ScheduleEntry {
	fmt.Printf("\nPotential Swaps (slot - user -> slot - user)\n")
	swapAPI := &conflict.SwapAPI{}
	swaps := map[*pduty.ScheduleEntry]*pduty.ScheduleEntry{}

	for _, conflict := range conflicts {
		swap := swapAPI.FindSwap(periodStart, schedule, conflict, calendars)
		if swap != nil {
			fmt.Printf("%s - %s - %s", conflict.Start.Format(timeFormat), conflict.End.Format(timeFormat), conflict.User.Name)
			fmt.Printf(" -> %s - %s - %s\n", swap.Start.Format(timeFormat), swap.End.Format(timeFormat), swap.User.Name)
			swaps[conflict] = swap
			continue
		}

		fmt.Fprintf(os.Stderr, "\n ==> SWAP NOT FOUND FOR %s - %s - %s <==\n\n", conflict.Start.Format(timeFormat), conflict.End.Format(timeFormat), conflict.User.Name)
	}

	return swaps
}
