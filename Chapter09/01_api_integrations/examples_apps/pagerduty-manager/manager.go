package pdmanager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/escalations"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/schedules"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/services"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/teams"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/users"
	"go.uber.org/zap"
)

const (
	roleMember   = "member"
	roleObserver = "observer"
	roleLead     = "lead"
	roleDeptHead = "dept-head"
)

// map of our roles to PD user roles
var rolesToPDUserRoles = map[string]string{
	roleMember:   "limited_user",
	roleObserver: "limited_user",
	roleLead:     "user",
	roleDeptHead: "admin",
}

// map of our roles to PD user roles
var rolesToPDTeamRoles = map[string]string{
	roleMember:   "responder",
	roleObserver: "responder",
	roleLead:     "manager",
	roleDeptHead: "manager",
}

func New(cfg Config, logger *zap.Logger) *Manager {
	return &Manager{
		cfg:           cfg,
		logger:        logger,
		companyConfig: &companyConfig{},
	}
}

// Manager is the main entry point for this package/tool
type Manager struct {
	cfg    Config
	logger *zap.Logger

	companyConfig *companyConfig

	userManager       *users.Manager
	teamManager       *teams.Manager
	scheduleManager   *schedules.Manager
	escalationManager *escalations.Manager
	serviceManager    *services.Manager
}

// Parse attempts to parse the provide file into this manager
func (m *Manager) Parse(_ context.Context) error {
	m.logger.Debug("loading data from file", zap.String("file", m.cfg.Filename()))

	fileContents, err := ioutil.ReadFile(m.cfg.Filename())
	if err != nil {
		return fmt.Errorf("failed to read input file with err: %w", err)
	}

	err = json.Unmarshal(fileContents, m.companyConfig)
	if err != nil {
		return fmt.Errorf("failed to parse config JSON with err: %w", err)
	}

	return m.validate()
}

func (m *Manager) validate() error {
	if len(m.companyConfig.Teams) == 0 {
		return errors.New("no teams found in the JSON")
	}

	for _, thisTeam := range m.companyConfig.Teams {
		for _, thisMember := range thisTeam.Members {
			_, ok := rolesToPDUserRoles[thisMember.Role]
			if !ok {
				return fmt.Errorf("invalid role 'value in: %v", thisMember)
			}
		}
	}

	return nil
}

// Sync calls all of the Sync Methods in the correct order
func (m *Manager) Sync(ctx context.Context) error {
	err := m.SyncUsers(ctx)
	if err != nil {
		return err
	}

	err = m.SyncTeams(ctx)
	if err != nil {
		return err
	}

	err = m.SyncSchedules(ctx)
	if err != nil {
		return err
	}

	err = m.SyncEscalation(ctx)
	if err != nil {
		return err
	}

	err = m.SyncServices(ctx)
	if err != nil {
		return err
	}

	return nil
}

// SyncUsers attempts to download the existing users and create any that do not yet exist.
// Note: existing data will not be modified in any way.
func (m *Manager) SyncUsers(ctx context.Context) error {
	var members []*Member
	for _, team := range m.companyConfig.Teams {
		members = append(members, team.Members...)
	}

	m.userManager = users.New(m.cfg, m.logger)

	for _, member := range members {
		fetchedUser, err := m.userManager.GetByEmail(ctx, member.Email)
		if err == nil {
			// user exists
			member.ID = fetchedUser.ID
			continue
		}

		if !errors.Is(err, users.ErrNoSuchUser) {
			m.logger.Error("failed to sync users - fetch user failed", zap.Error(err))
			return err
		}

		member.ID, err = m.userManager.Add(ctx, member, m.companyConfig.DefaultTimezone)
		if err != nil {
			m.logger.Error("failed to sync users - add user failed", zap.Error(err))
			return err
		}
	}

	return nil
}

// SyncTeams attempts to download the existing teams and create any that do not yet exist.
// Note: existing data will not be modified in any way.
// Note: creating a team also creates a matching service so we can have an `@oncall-[team]` slack alias
func (m *Manager) SyncTeams(ctx context.Context) error {
	m.teamManager = teams.New(m.cfg, m.logger)

	for _, team := range m.companyConfig.Teams {
		fetchedTeam, err := m.teamManager.GetByName(ctx, team.Name)
		if err == nil {
			// team exists
			team.ID = fetchedTeam.ID

			err = m.syncTeamMembers(ctx, team)
			if err != nil {
				return err
			}

			continue
		}

		if !errors.Is(err, teams.ErrNoSuchTeam) {
			m.logger.Error("failed to sync teams - fetch team failed", zap.Error(err))
			return err
		}

		team.ID, err = m.teamManager.Add(ctx, team.Name, team.Description)
		if err != nil {
			m.logger.Error("failed to sync teams - add team failed", zap.Error(err))
			return err
		}

		err = m.syncTeamMembers(ctx, team)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) syncTeamMembers(ctx context.Context, team *Team) error {
	for _, member := range team.Members {
		err := m.teamManager.AddMember(ctx, team.ID, member)
		if err != nil {
			m.logger.Error("failed to sync teams - add team member failed", zap.Error(err))
			return err
		}
	}

	return nil
}

// SyncSchedules attempts to download the existing schedules and create any that do not yet exist.
// Note: existing data will not be modified in any way.
func (m *Manager) SyncSchedules(ctx context.Context) error {
	m.scheduleManager = schedules.New(m.cfg, m.logger)

	for _, team := range m.companyConfig.Teams {
		fetchedSchedule, err := m.scheduleManager.GetByName(ctx, team.Name)
		if err == nil {
			team.ScheduleID = fetchedSchedule.ID

			err = m.scheduleManager.Update(ctx, fetchedSchedule.ID, team, m.companyConfig.DefaultTimezone)
			if err != nil {
				m.logger.Error("failed to sync schedule - update schedule failed", zap.Error(err))
				return err
			}

			continue
		}

		if !errors.Is(err, schedules.ErrNoSuchSchedule) {
			m.logger.Error("failed to sync schedule - fetch schedule failed", zap.Error(err))
			return err
		}

		team.ScheduleID, err = m.scheduleManager.Add(ctx, team, m.companyConfig.DefaultTimezone)
		if err != nil {
			m.logger.Error("failed to sync schedule - add schedule failed", zap.Error(err))
			return err
		}
	}

	return nil
}

// SyncEscalation attempts to download the existing escalation policies and create any that do not yet exist.
// Note: existing data will not be modified in any way.
func (m *Manager) SyncEscalation(ctx context.Context) error {
	m.escalationManager = escalations.New(m.cfg, m.logger)

	for _, team := range m.companyConfig.Teams {
		fetchedEscalation, err := m.escalationManager.GetByName(ctx, team.Name)
		if err == nil {
			team.PolicyID = fetchedEscalation.ID

			err = m.escalationManager.Update(ctx, fetchedEscalation.ID, team)
			if err != nil {
				m.logger.Error("failed to sync escalation - update escalation failed", zap.Error(err))
				return err
			}

			continue
		}

		if !errors.Is(err, escalations.ErrNoSuchPolicy) {
			m.logger.Error("failed to sync escalation - fetch escalation failed", zap.Error(err))
			return err
		}

		team.PolicyID, err = m.escalationManager.Add(ctx, team)
		if err != nil {
			m.logger.Error("failed to sync escalation - add escalation failed", zap.Error(err))
			return err
		}
	}

	return nil
}

// SyncServices attempts to download the existing services and create any that do not yet exist.
// Note: existing data will not be modified in any way.
func (m *Manager) SyncServices(ctx context.Context) error {
	m.serviceManager = services.New(m.cfg, m.logger)

	for _, team := range m.companyConfig.Teams {
		for _, service := range team.Services {
			err := m.upsertService(ctx, service, team)
			if err != nil {
				return err
			}
		}

		// fake service for the team (to make @oncall-[team]
		teamService := &Service{
			Name:      team.Name,
			Dashboard: team.Description,
		}

		err := m.upsertService(ctx, teamService, team)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) upsertService(ctx context.Context, service *Service, team *Team) error {
	fetchedService, err := m.serviceManager.GetByName(ctx, service.Name)
	if err == nil {
		err = m.serviceManager.Update(ctx, fetchedService.ID, service, team)
		if err != nil {
			m.logger.Error("failed to sync service - update service failed", zap.Error(err))
			return err
		}

		return nil
	}

	if !errors.Is(err, services.ErrNoSuchService) {
		m.logger.Error("failed to sync service - fetch service failed", zap.Error(err))
		return err
	}

	_, err = m.serviceManager.Add(ctx, service, team)
	if err != nil {
		m.logger.Error("failed to sync service - add service failed", zap.Error(err))
		return err
	}

	return nil
}

// Config is the config for this package
type Config interface {
	Debug() bool
	Filename() string
	BaseURL() string
	AuthToken() string
}

type companyConfig struct {
	Teams           []*Team `json:"teams"`
	DefaultTimezone string  `json:"default_timezone"`
}

type Team struct {
	ID          string     `json:"-"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Slack       string     `json:"slack"`
	Members     []*Member  `json:"members"`
	Services    []*Service `json:"services"`
	ScheduleID  string     `json:"-"`
	PolicyID    string     `json:"-"`
}

func (t *Team) GetEscalationPolicyID() string {
	return t.PolicyID
}

func (t *Team) GetScheduleID() string {
	return t.ScheduleID
}

func (t *Team) GetDeptHeadsIDs() []string {
	var ids []string

	for _, member := range t.Members {
		if member.Role == roleDeptHead {
			ids = append(ids, member.ID)
		}
	}

	return ids
}

func (t *Team) GetResponderIDs() []string {
	var ids []string

	for _, member := range t.Members {
		if member.Role == roleMember {
			ids = append(ids, member.ID)
		}
	}

	return ids
}

func (t *Team) GetLeadIDs() []string {
	var ids []string

	for _, member := range t.Members {
		if member.Role == roleLead {
			ids = append(ids, member.ID)
		}
	}

	return ids
}

func (t *Team) GetTeamName() string {
	return t.Name
}

func (t *Team) GetDescription() string {
	return t.Description
}

func (t *Team) GetTimeZone() string {
	return ""
}

func (t *Team) GetTeamID() string {
	return t.ID
}

type Member struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
	Role     string `json:"role"`
}

func (m *Member) GetUserID() string {
	return m.ID
}

func (m *Member) GetName() string {
	return m.Name
}

func (m *Member) GetEmail() string {
	return m.Email
}

func (m *Member) GetTimeZone() string {
	return m.Timezone
}

func (m *Member) GetUserRole() string {
	return rolesToPDUserRoles[m.Role]
}

func (m *Member) GetTeamRole() string {
	return rolesToPDTeamRoles[m.Role]
}

type Service struct {
	Name      string `json:"name"`
	Dashboard string `json:"dashboard"`
}

func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) GetDescription() string {
	return s.Dashboard
}
