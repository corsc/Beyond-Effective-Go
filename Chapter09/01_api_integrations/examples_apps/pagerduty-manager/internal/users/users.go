package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/pd"
	"go.uber.org/zap"
	"net/url"
)

const listURI = "/users"

var ErrNoSuchUser = errors.New("no such user")

func New(cfg Config, logger *zap.Logger) *Manager {
	return &Manager{
		cfg:    cfg,
		logger: logger,
		api:    pd.New(cfg, logger),
	}
}

// Manager allows for loading and creating users
type Manager struct {
	cfg    Config
	logger *zap.Logger
	api    *pd.API
}

func (u *Manager) GetByEmail(ctx context.Context, email string) (*User, error) {
	params := url.Values{}
	params.Set("query", email)
	params.Set("total", "false")
	params.Set("limit", "1")

	users := &listResponse{}

	err := u.api.Get(ctx, listURI, params, users)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email '%s` with err: %s", email, err)
	}

	if len(users.Users) == 0 {
		return nil, ErrNoSuchUser
	}

	return users.Users[0], nil
}

func (u *Manager) Add(ctx context.Context, user NewUser, defaultTimeZone string) (string, error) {
	reqDTO := newNewUserRequest(user, defaultTimeZone)

	respDTO := &newUserResponse{}

	err := u.api.Post(ctx, listURI, reqDTO, respDTO)
	if err != nil {
		return "", fmt.Errorf("failed to add user '%#v` with err: %s", user, err)
	}

	return respDTO.User.ID, nil
}

type NewUser interface {
	GetName() string
	GetEmail() string
	GetTimeZone() string
	GetUserRole() string
}

func newNewUserRequest(user NewUser, defaultTimeZone string) *newUserRequest {
	timeZone := user.GetTimeZone()
	if timeZone == "" {
		timeZone = defaultTimeZone
	}

	out := &newUserRequest{
		User: &userFormat{
			Type:     "user",
			Name:     user.GetName(),
			Email:    user.GetEmail(),
			TimeZone: timeZone,
			Role:     user.GetUserRole(),
		},
	}

	return out
}

type newUserRequest struct {
	User *userFormat `json:"user"`
}

type newUserResponse struct {
	User userFormat `json:"user"`
}

type userFormat struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	TimeZone string `json:"time_zone"`
	Role     string `json:"role"`
}

type listResponse struct {
	Users []*User `json:"users"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Teams []Team `json:"teams"`
}

type Team struct {
	ID      string `json:"id"`
	Summary string `json:"summary"`
}

type Config interface {
	Debug() bool
	BaseURL() string
	AuthToken() string
}
