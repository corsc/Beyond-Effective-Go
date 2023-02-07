package conflict

import (
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/gcal"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/pduty"
)

// SwapAPI will attempt to find a swap in the schedule
type SwapAPI struct {
	checker *CheckerAPI

	proposedSwaps []*pduty.ScheduleEntry
}

// FindSwap attempts to find a swap for the supplied conflict
func (s *SwapAPI) FindSwap(periodStart time.Time, schedule *pduty.Schedule, conflict *pduty.ScheduleEntry, calendars map[string]*gcal.Calendar) *pduty.ScheduleEntry {
	if s.checker == nil {
		s.checker = &CheckerAPI{}
	}

	for _, potentialSwap := range schedule.Entries {
		if potentialSwap.Start.Equal(conflict.Start) && potentialSwap.End.Equal(conflict.End) ||
			potentialSwap.User.ID == conflict.User.ID {
			// cant swap with the same slot or same user
			continue
		}

		if potentialSwap.Start.Before(periodStart) {
			// cant swap with slots in the prior to start of the schedule period
			continue
		}

		if !s.timeEqual(potentialSwap.Start, conflict.Start) || !s.timeEqual(potentialSwap.End, conflict.End) {
			continue
		}

		if s.checker.checkForConflict(conflict, calendars[potentialSwap.User.ID]) {
			// potential swap user cannot take the conflict shift
			continue
		}

		if s.checker.checkForConflict(potentialSwap, calendars[conflict.User.ID]) {
			// conflict user cannot take the potential swap's shift
			continue
		}

		if s.isAlreadySwapped(potentialSwap) {
			// ensure we have not already included this user/slot in a previous swap
			continue
		}

		s.proposedSwaps = append(s.proposedSwaps, potentialSwap)

		return potentialSwap
	}

	return nil
}

// compare the hour and minute only
func (s *SwapAPI) timeEqual(a time.Time, b time.Time) bool {
	return a.Hour() == b.Hour() && a.Minute() == b.Minute()
}

func (s *SwapAPI) isAlreadySwapped(potentialSwap *pduty.ScheduleEntry) bool {
	for _, thisSwap := range s.proposedSwaps {
		if thisSwap.Start.Equal(potentialSwap.Start) && thisSwap.End.Equal(potentialSwap.End) {
			return true
		}
	}

	return false
}
