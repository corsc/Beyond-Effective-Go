package conflict

import (
	"testing"
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/gcal"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-gcal/internal/pduty"
	"github.com/stretchr/testify/assert"
)

var (
	testUserFoo = "FOO"
	testUserBar = "BAR"
)

func TestCheckerAPI_Check(t *testing.T) {
	scenarios := []struct {
		desc              string
		inSchedule        *pduty.Schedule
		inCalendars       map[string]*gcal.Calendar
		inMinimumDays     int64
		expectedConflicts int
		expectErr         bool
	}{
		{
			desc:              "happy path - no inputs",
			inSchedule:        &pduty.Schedule{},
			inCalendars:       map[string]*gcal.Calendar{},
			expectedConflicts: 0,
			expectErr:         false,
		},
		{
			desc: "no conflict - ooo start and end before",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {
					Items: []*gcal.CalendarItem{
						{
							Start: time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC),
							End:   time.Date(2019, 01, 01, 2, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			expectedConflicts: 0,
			expectErr:         false,
		},
		{
			desc: "no conflict - ooo start and end after",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {
					Items: []*gcal.CalendarItem{
						{
							Start: time.Date(2019, 01, 03, 0, 0, 0, 0, time.UTC),
							End:   time.Date(2019, 01, 04, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			expectedConflicts: 0,
			expectErr:         false,
		},
		{
			desc: "conflict - ooo start and end equal",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {
					Items: []*gcal.CalendarItem{
						{
							Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
							End:   time.Date(2019, 01, 03, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			expectedConflicts: 1,
			expectErr:         false,
		},
		{
			desc: "conflict - ooo start before and end after",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {
					Items: []*gcal.CalendarItem{
						{
							Start: time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC),
							End:   time.Date(2019, 01, 04, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			expectedConflicts: 1,
			expectErr:         false,
		},
		{
			desc: "conflict - ooo start after and end before",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {
					Items: []*gcal.CalendarItem{
						{
							Start: time.Date(2019, 01, 02, 1, 0, 0, 0, time.UTC),
							End:   time.Date(2019, 01, 02, 2, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			expectedConflicts: 1,
			expectErr:         false,
		},
		{
			desc: "conflict - minimum days between violation",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 03, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 03, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {},
			},
			inMinimumDays:     7,
			expectedConflicts: 1,
			expectErr:         false,
		},
		{
			desc: "no conflict - minimum days between violation",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 10, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 10, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {},
			},
			inMinimumDays:     7,
			expectedConflicts: 0,
			expectErr:         false,
		},
		{
			desc: "no conflict - minimum days between violation (different user)",
			inSchedule: &pduty.Schedule{
				Entries: []*pduty.ScheduleEntry{
					{
						User: &pduty.User{
							ID: testUserBar,
						},
						Start: time.Date(2019, 01, 02, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 02, 8, 0, 0, 0, time.UTC),
					},
					{
						User: &pduty.User{
							ID: testUserFoo,
						},
						Start: time.Date(2019, 01, 03, 0, 0, 0, 0, time.UTC),
						End:   time.Date(2019, 01, 03, 8, 0, 0, 0, time.UTC),
					},
				},
			},
			inCalendars: map[string]*gcal.Calendar{
				testUserFoo: {},
			},
			inMinimumDays:     7,
			expectedConflicts: 0,
			expectErr:         false,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// call
			api := &CheckerAPI{}
			result, resultErr := api.Check(scenario.inSchedule, scenario.inCalendars, scenario.inMinimumDays)

			// validate
			assert.Equal(t, scenario.expectedConflicts, len(result), scenario.desc)
			assert.Equal(t, scenario.expectErr, resultErr != nil, scenario.desc)
		})
	}
}
