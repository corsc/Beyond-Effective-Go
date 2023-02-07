package conflict

import (
	"testing"
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/gcal"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/pduty"
	"github.com/stretchr/testify/assert"
)

var (
	periodStart = time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)

	sourceUserID      = "FOO"
	destinationUserID = "BAR"

	dayPastMorning   = time.Date(2018, 12, 30, 0, 0, 0, 0, time.UTC)
	dayPastAfternoon = time.Date(2018, 12, 30, 8, 0, 0, 0, time.UTC)

	day2Morning   = time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC)
	day2Afternoon = time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC)
	day2Evening   = time.Date(2019, 01, 02, 16, 0, 0, 0, time.UTC)

	day3Morning   = day2Morning.Add(24 * time.Hour)
	day3Afternoon = day2Afternoon.Add(24 * time.Hour)

	day2MorningSource = &pduty.ScheduleEntry{
		User: &pduty.User{
			ID: sourceUserID,
		},
		Start: day2Morning,
		End:   day2Afternoon,
	}

	day2AfternoonDestination = &pduty.ScheduleEntry{
		User: &pduty.User{
			ID: sourceUserID,
		},
		Start: day2Afternoon,
		End:   day2Evening,
	}

	day3MorningDestination = &pduty.ScheduleEntry{
		User: &pduty.User{
			ID: destinationUserID,
		},
		Start: day3Morning,
		End:   day3Afternoon,
	}

	dayPastMorningDestination = &pduty.ScheduleEntry{
		User: &pduty.User{
			ID: destinationUserID,
		},
		Start: dayPastMorning,
		End:   dayPastAfternoon,
	}
)

func TestSwapAPI_FindSwap(t *testing.T) {
	scenarios := []struct {
		desc             string
		inSchedule       *pduty.Schedule
		inConflict       *pduty.ScheduleEntry
		inCalendars      map[string]*gcal.Calendar
		inAlreadySwapped []*pduty.ScheduleEntry
		expected         *pduty.ScheduleEntry
	}{
		{
			desc: "swap available",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					day2MorningSource,
					day3MorningDestination,
				},
			},
			inConflict: day2MorningSource,
			inCalendars: map[string]*gcal.Calendar{
				sourceUserID:      {},
				destinationUserID: {},
			},
			expected: day3MorningDestination,
		},
		{
			desc: "swap not possible, in the past",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					day2MorningSource,
					dayPastMorningDestination,
				},
			},
			inConflict: day2MorningSource,
			inCalendars: map[string]*gcal.Calendar{
				sourceUserID:      {},
				destinationUserID: {},
			},
			expected: nil,
		},
		{
			desc: "already swapped",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					day2MorningSource,
					day3MorningDestination,
				},
			},
			inConflict: day2MorningSource,
			inCalendars: map[string]*gcal.Calendar{
				sourceUserID:      {},
				destinationUserID: {},
			},
			inAlreadySwapped: []*pduty.ScheduleEntry{
				day3MorningDestination,
			},
			expected: nil,
		},
		{
			desc: "swap not possible because destination is not available",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					day2MorningSource,
					day3MorningDestination,
				},
			},
			inConflict: day2MorningSource,
			inCalendars: map[string]*gcal.Calendar{
				sourceUserID: {},
				destinationUserID: {
					Items: []*gcal.CalendarItem{
						{
							Start: day2Morning,
							End:   day2Afternoon,
						},
					},
				},
			},
			expected: nil,
		},
		{
			desc: "swap not possible because original user cannot take the replacement's shift",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					day2MorningSource,
					day3MorningDestination,
				},
			},
			inConflict: day2MorningSource,
			inCalendars: map[string]*gcal.Calendar{
				sourceUserID: {
					Items: []*gcal.CalendarItem{
						{
							Start: day3Morning,
							End:   day3Afternoon,
						},
					},
				},
				destinationUserID: {},
			},
			expected: nil,
		},
		{
			desc: "cannot swap with yourself",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					day2MorningSource, day2MorningSource,
				},
			},
			inConflict: day2MorningSource,
			inCalendars: map[string]*gcal.Calendar{
				sourceUserID: {},
			},
			expected: nil,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// call
			api := &SwapAPI{
				proposedSwaps: scenario.inAlreadySwapped,
			}
			result := api.FindSwap(periodStart, scenario.inSchedule, scenario.inConflict, scenario.inCalendars)

			// validate
			assert.EqualValues(t, scenario.expected, result, scenario.desc)
		})
	}
}
