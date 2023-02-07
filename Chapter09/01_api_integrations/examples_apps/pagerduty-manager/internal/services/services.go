package services

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/pd"
	"go.uber.org/zap"
)

const (
	getURI    = "/services/%s"
	listURI   = "/services"
	addURI    = "/services"
	updateURI = "/services/%s"
)

var ErrNoSuchService = errors.New("no such service")

func New(cfg Config, logger *zap.Logger) *Manager {
	return &Manager{
		cfg:    cfg,
		logger: logger,
		api:    pd.New(cfg, logger),
	}
}

// Manager allows for loading and creating services
type Manager struct {
	cfg    Config
	logger *zap.Logger
	api    *pd.API
}

func (u *Manager) Get(ctx context.Context, serviceID string) (*Service, error) {
	uri := fmt.Sprintf(getURI, serviceID)

	services := &getServiceResponse{}

	err := u.api.Get(ctx, uri, nil, services)
	if err != nil {
		return nil, fmt.Errorf("failed to get service '%s' with err: %s", serviceID, err)
	}

	if services.Service == nil {
		return nil, ErrNoSuchService
	}

	return services.Service, nil
}

func (u *Manager) GetByName(ctx context.Context, name string) (*Service, error) {
	params := url.Values{}
	params.Set("query", name)
	params.Set("total", "false")
	params.Set("limit", "1")

	services := &getServicesResponse{}

	err := u.api.Get(ctx, listURI, params, services)
	if err != nil {
		return nil, fmt.Errorf("failed to get services '%s' with err: %s", name, err)
	}

	if len(services.Service) == 0 {
		return nil, ErrNoSuchService
	}

	return services.Service[0], nil
}

func (u *Manager) Add(ctx context.Context, service NewService, team NewTeam) (string, error) {
	reqDTO := u.buildAddPayload(service, team)

	respDTO := &addResponse{}

	err := u.api.Post(ctx, addURI, reqDTO, respDTO)
	if err != nil {
		return "", fmt.Errorf("failed to add service '%#v' with err: %s", reqDTO, err)
	}

	return respDTO.Service.ID, nil
}

func (u *Manager) buildAddPayload(service NewService, team NewTeam) *addRequest {
	return &addRequest{
		Service: &Service{
			Name:        service.GetName(),
			Description: service.GetDescription(),
			Status:      "active",
			EscalationPolicy: &EscalationPolicy{
				ID:   team.GetEscalationPolicyID(),
				Type: "escalation_policy_reference",
			},
			Teams: []*Team{
				{
					ID: team.GetTeamID(),
				},
			},
			IncidentUrgencyRule: &IncidentUrgency{
				Type:    "constant",
				Urgency: "high",
			},
			AlertCreation: "create_alerts_and_incidents",
			AlertGroupingParameters: &AlertGroupParameters{
				Type: "intelligent",
			},
		},
	}
}

func (u *Manager) Update(ctx context.Context, serviceID string, service NewService, team NewTeam) error {
	reqDTO := u.buildAddPayload(service, team)

	reqDTO.Service.ID = serviceID

	uri := fmt.Sprintf(updateURI, serviceID)

	err := u.api.Put(ctx, uri, reqDTO)
	if err != nil {
		return fmt.Errorf("failed to update service '%#v' with err: %s", reqDTO, err)
	}

	return nil
}

type NewService interface {
	GetName() string
	GetDescription() string
}

type NewTeam interface {
	GetTeamID() string
	GetEscalationPolicyID() string
}

type getServiceResponse struct {
	Service *Service `json:"service"`
}

type getServicesResponse struct {
	Service []*Service `json:"services"`
}

type Service struct {
	ID                      string                `json:"id"`
	Type                    string                `json:"type"`
	Name                    string                `json:"name"`
	Description             string                `json:"description"`
	Status                  string                `json:"status"`
	EscalationPolicy        *EscalationPolicy     `json:"escalation_policy"`
	Teams                   []*Team               `json:"teams"`
	IncidentUrgencyRule     *IncidentUrgency      `json:"incident_urgency_rule"`
	AlertCreation           string                `json:"alert_creation"`
	AlertGroupingParameters *AlertGroupParameters `json:"alert_grouping_parameters"`
}

type EscalationPolicy struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Team struct {
	ID string `json:"id"`
}

type IncidentUrgency struct {
	Type    string `json:"type"`
	Urgency string `json:"urgency"`
}

type AlertGroupParameters struct {
	Type string `json:"type"`
}

type addRequest struct {
	Service *Service `json:"service"`
}

type addResponse struct {
	Service *Service `json:"service"`
}

type Config interface {
	Debug() bool
	BaseURL() string
	AuthToken() string
}
