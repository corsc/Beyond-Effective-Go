package teams

import (
	"context"
	"errors"
	"fmt"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/pd"
	"go.uber.org/zap"
	"net/url"
)

const (
	getURI         = "/teams/%s"
	listURI        = "/teams"
	addURI         = "/teams"
	listMembersURI = "/teams/%s/members"
	addMemberURI   = "/teams/%s/users/%s"
)

var (
	ErrNoSuchTeam = errors.New("no such team")
	ErrNoMembers  = errors.New("no members")
)

func New(cfg Config, logger *zap.Logger) *Manager {
	return &Manager{
		cfg:    cfg,
		logger: logger,
		api:    pd.New(cfg, logger),
	}
}

// Manager allows for loading and creating teams
type Manager struct {
	cfg    Config
	logger *zap.Logger
	api    *pd.API
}

func (u *Manager) Get(ctx context.Context, teamID string) (*Team, error) {
	uri := fmt.Sprintf(getURI, teamID)

	teams := &getTeamResponse{}

	err := u.api.Get(ctx, uri, nil, teams)
	if err != nil {
		return nil, fmt.Errorf("failed to get team '%s' with err: %s", teamID, err)
	}

	if teams.Team == nil {
		return nil, ErrNoSuchTeam
	}

	return teams.Team, nil
}

func (u *Manager) GetByName(ctx context.Context, name string) (*Team, error) {
	params := url.Values{}
	params.Set("query", name)
	params.Set("total", "false")
	params.Set("limit", "1")

	teams := &getTeamsResponse{}

	err := u.api.Get(ctx, listURI, params, teams)
	if err != nil {
		return nil, fmt.Errorf("failed to get teams '%s' with err: %s", name, err)
	}

	if len(teams.Team) == 0 {
		return nil, ErrNoSuchTeam
	}

	return teams.Team[0], nil
}

func (u *Manager) GetMembers(ctx context.Context, teamID string) ([]*Member, error) {
	uri := fmt.Sprintf(listMembersURI, teamID)

	params := url.Values{}
	params.Set("total", "true")

	team := &getTeamMembersResponse{}

	err := u.api.Get(ctx, uri, params, team)
	if err != nil {
		return nil, fmt.Errorf("failed to get team members for team '%s' with err: %s", teamID, err)
	}

	if team.Members == nil {
		return nil, ErrNoMembers
	}

	members := make([]*Member, len(team.Members))

	for index, member := range team.Members {
		members[index] = &Member{
			ID:   member.User.ID,
			Role: member.Role,
		}
	}

	return members, nil
}

func (u *Manager) Add(ctx context.Context, name, description string) (string, error) {
	reqDTO := &addRequest{
		Team: Team{
			Name:        name,
			Description: description,
		},
	}

	respDTO := &addResponse{}

	err := u.api.Post(ctx, addURI, reqDTO, respDTO)
	if err != nil {
		return "", fmt.Errorf("failed to add team '%#v' with err: %s", reqDTO, err)
	}

	return respDTO.Team.ID, nil
}

func (u *Manager) AddMember(ctx context.Context, teamID string, user User) error {
	uri := fmt.Sprintf(addMemberURI, teamID, user.GetUserID())

	reqDTO := &addMemberRequest{
		Role: user.GetTeamRole(),
	}

	err := u.api.Put(ctx, uri, reqDTO)
	if err != nil {
		return fmt.Errorf("failed to add user '%#v' to team '%s' with err: %s", user, teamID, err)
	}

	return nil
}

type User interface {
	GetUserID() string
	GetTeamRole() string
}

type getTeamResponse struct {
	Team *Team `json:"team"`
}

type getTeamsResponse struct {
	Team []*Team `json:"teams"`
}

type Team struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Member struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type getTeamMembersResponse struct {
	Members []*member `json:"members"`
}

type member struct {
	User *user  `json:"user"`
	Role string `json:"role"`
}

type user struct {
	ID string `json:"id"`
}

type addRequest struct {
	Team Team `json:"team"`
}

type addResponse struct {
	Team Team `json:"team"`
}

type addMemberRequest struct {
	Role string `json:"role"`
}

type Config interface {
	Debug() bool
	BaseURL() string
	AuthToken() string
}
