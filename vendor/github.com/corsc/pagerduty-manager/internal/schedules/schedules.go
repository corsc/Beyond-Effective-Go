package schedules

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/corsc/pagerduty-manager/internal/pd"

	"go.uber.org/zap"
)

const (
	getURI    = "/schedules/%s"
	listURI   = "/schedules"
	addURI    = "/schedules?overflow=true"
	updateURI = "/schedules/%s?overflow=true"
)

var (
	ErrNoSuchSchedule = errors.New("no such schedule")

	rotationLengthSeconds = 60 * 60 * 24 * 7
)

func New(cfg Config, logger *zap.Logger) *Manager {
	return &Manager{
		cfg:    cfg,
		logger: logger,
		api:    pd.New(cfg, logger),
	}
}

// Manager allows for loading and creating schedules
type Manager struct {
	cfg    Config
	logger *zap.Logger
	api    *pd.API
}

func (u *Manager) Get(ctx context.Context, scheduleID string) (*Schedule, error) {
	uri := fmt.Sprintf(getURI, scheduleID)

	schedules := &getServiceResponse{}

	err := u.api.Get(ctx, uri, nil, schedules)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule '%s' with err: %s", scheduleID, err)
	}

	if schedules.Schedule == nil {
		return nil, ErrNoSuchSchedule
	}

	return schedules.Schedule, nil
}

func (u *Manager) GetByName(ctx context.Context, name string) (*Schedule, error) {
	params := url.Values{}
	params.Set("query", name)
	params.Set("total", "false")
	params.Set("limit", "1")

	schedules := &getScheduleResponse{}

	err := u.api.Get(ctx, listURI, params, schedules)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedules '%s' with err: %s", name, err)
	}

	if len(schedules.Schedules) == 0 {
		return nil, ErrNoSuchSchedule
	}

	return schedules.Schedules[0], nil
}

func (u *Manager) Add(ctx context.Context, schedule ReqSchedule, defaultTimeZone string) (string, error) {
	location, err := time.LoadLocation(defaultTimeZone)
	if err != nil {
		return "", fmt.Errorf("failed to determine location with err: %w", err)
	}

	if len(schedule.GetResponderIDs()) == 0 {
		return "", fmt.Errorf("cannot create schedule with no responders")
	}

	reqDTO := &addRequest{
		Schedule: &Schedule{
			Name:        schedule.GetTeamName() + " Schedule",
			Description: schedule.GetDescription(),
			TimeZone:    defaultTimeZone,
			Teams: []*team{
				{
					ID: schedule.GetTeamID(),
				},
			},
			ScheduleLayers: []*scheduleLayer{
				buildMemberLayer(schedule, location),
			},
		},
	}

	// start and virtual start the same
	reqDTO.Schedule.ScheduleLayers[0].RotationVirtualStart = time.Date(2021, time.July, 5, 11, 0, 0, 0, location)

	respDTO := &addResponse{}

	err = u.api.Post(ctx, addURI, reqDTO, respDTO)
	if err != nil {
		return "", fmt.Errorf("failed to add schedule '%#v' with err: %s", reqDTO, err)
	}

	return respDTO.Schedule.ID, nil
}

func (u *Manager) Update(ctx context.Context, scheduleID string, schedule ReqSchedule, defaultTimeZone string) error {
	scheduleToUpdate, err := u.Get(ctx, scheduleID)
	if err != nil {
		return fmt.Errorf("failed to update schedule with err: %w", err)
	}

	location, err := time.LoadLocation(defaultTimeZone)
	if err != nil {
		return fmt.Errorf("failed to determine location with err: %w", err)
	}

	if len(schedule.GetResponderIDs()) == 0 {
		return fmt.Errorf("cannot update schedule with no responders")
	}

	updateLayer(scheduleToUpdate, schedule, defaultTimeZone, location)

	updateMembers(schedule, scheduleToUpdate)

	uri := fmt.Sprintf(updateURI, scheduleID)

	err = u.api.Put(ctx, uri, scheduleToUpdate)
	if err != nil {
		return fmt.Errorf("failed to add schedule '%#v' with err: %s", scheduleToUpdate, err)
	}

	return nil
}

func updateLayer(scheduleToUpdate *Schedule, schedule ReqSchedule, defaultTimeZone string, location *time.Location) {
	scheduleToUpdate.Name = schedule.GetTeamName() + " Schedule"
	scheduleToUpdate.Description = schedule.GetDescription()
	scheduleToUpdate.TimeZone = defaultTimeZone
	scheduleToUpdate.Teams = []*team{
		{
			ID: schedule.GetTeamID(),
		},
	}

	scheduleToUpdate.ScheduleLayers[0].Name = "Layer 1"
	scheduleToUpdate.ScheduleLayers[0].Start = time.Date(2021, time.July, 5, 11, 0, 0, 0, location)
	scheduleToUpdate.ScheduleLayers[0].RotationTurnLengthSeconds = rotationLengthSeconds

	// virtual start is the next Monday
	now := time.Now()
	virtualStart := time.Date(now.Year(), now.Month(), now.Day(), 11, 0, 0, 0, location)

	if time.Now().Weekday() == time.Sunday {
		virtualStart = virtualStart.Add(24 * time.Hour)
	} else {
		virtualStart = virtualStart.Add(24 * time.Hour * time.Duration(8-time.Now().Weekday()))
	}

	scheduleToUpdate.ScheduleLayers[0].RotationVirtualStart = virtualStart
}

func updateMembers(schedule ReqSchedule, scheduleToUpdate *Schedule) {
	var members []*user

	for _, userID := range schedule.GetResponderIDs() {
		members = append(members, &user{
			ID:   userID,
			Type: "user",
		})
	}

	for _, userID := range schedule.GetLeadIDs() {
		members = append(members, &user{
			ID:   userID,
			Type: "user",
		})
	}

	scheduleToUpdate.ScheduleLayers[0].Users = members
}

func buildMemberLayer(schedule ReqSchedule, location *time.Location) *scheduleLayer {
	var members []*user

	for _, userID := range schedule.GetResponderIDs() {
		members = append(members, &user{
			ID:   userID,
			Type: "user",
		})
	}

	for _, userID := range schedule.GetLeadIDs() {
		members = append(members, &user{
			ID:   userID,
			Type: "user",
		})
	}

	return &scheduleLayer{
		Name:                      "Layer 1",
		Start:                     time.Date(2021, time.July, 5, 11, 0, 0, 0, location),
		RotationTurnLengthSeconds: rotationLengthSeconds,
		Users:                     members,
	}
}

type ReqSchedule interface {
	GetTeamName() string
	GetDescription() string
	GetResponderIDs() []string
	GetLeadIDs() []string
	GetTeamID() string
}

type getServiceResponse struct {
	Schedule *Schedule `json:"schedule"`
}

type getScheduleResponse struct {
	Schedules []*Schedule `json:"schedules"`
}

type Schedule struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	TimeZone       string           `json:"time_zone"`
	Teams          []*team          `json:"teams"`
	ScheduleLayers []*scheduleLayer `json:"schedule_layers"`
}

type user struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type team struct {
	ID string `json:"id"`
}

type scheduleLayer struct {
	ID                        string    `json:"id"`
	Name                      string    `json:"name"`
	Start                     time.Time `json:"start"`
	RotationVirtualStart      time.Time `json:"rotation_virtual_start"`
	RotationTurnLengthSeconds int       `json:"rotation_turn_length_seconds"`
	Users                     []*user   `json:"users"`
}

type addRequest struct {
	Schedule *Schedule `json:"schedule"`
}

type addResponse struct {
	Schedule *Schedule `json:"schedule"`
}

type Config interface {
	Debug() bool
	BaseURL() string
	AuthToken() string
}
