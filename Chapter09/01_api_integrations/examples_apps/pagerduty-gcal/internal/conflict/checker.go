package conflict

import (
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/gcal"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/pduty"
)

// CheckerAPI will compare the schedule with the calendar and return any conflicts
type CheckerAPI struct{}

// Check is the main entry point for this struct
func (c *CheckerAPI) Check(schedule *pduty.Schedule, calendars map[string]*gcal.Calendar, daysBetweenShifts int64) ([]*pduty.ScheduleEntry, error) {
	var conflictsOrdered []*pduty.ScheduleEntry

	for _, scheduleEntry := range schedule.Entries {
		scheduleUserID := scheduleEntry.User.ID

		// check schedule for "minimum days between shifts" violations
		conflict := c.checkMinimumDays(schedule, scheduleEntry, daysBetweenShifts)
		if conflict {
			conflictsOrdered = append(conflictsOrdered, scheduleEntry)
			continue
		}

		// check scheduled "unavailable" time
		calendar := calendars[scheduleUserID]
		if calendar == nil {
			// not items
			continue
		}

		conflict = c.checkForConflict(scheduleEntry, calendar)
		if conflict {
			conflictsOrdered = append(conflictsOrdered, scheduleEntry)
			continue
		}
	}

	return conflictsOrdered, nil
}

func (c *CheckerAPI) checkForConflict(shift *pduty.ScheduleEntry, calendar *gcal.Calendar) bool {
	if calendar == nil {
		return false
	}

	for _, calendarEntry := range calendar.Items {
		if calendarEntry.Start.Equal(shift.Start) {
			return true
		}

		if calendarEntry.Start.After(shift.Start) {
			if calendarEntry.Start.Before(shift.End) {
				return true
			}
		}

		if calendarEntry.Start.Before(shift.Start) {
			if calendarEntry.End.After(shift.Start) {
				return true
			}
		}
	}
	return false
}

func (c *CheckerAPI) checkMinimumDays(schedule *pduty.Schedule, currentEntry *pduty.ScheduleEntry, daysBetweenShifts int64) bool {
	for _, previousEntry := range schedule.Entries {
		if previousEntry.Start.After(currentEntry.Start) || previousEntry == currentEntry {
			// don't look at the future or the same entry
			return false
		}

		if currentEntry.User.ID != previousEntry.User.ID {
			// don't compare different users
			continue
		}

		if currentEntry.Start.Before(previousEntry.Start.Add(time.Duration(daysBetweenShifts*24) * time.Hour)) {
			return true
		}
	}

	return false
}
