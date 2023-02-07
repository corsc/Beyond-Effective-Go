package escalations

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/pd"
	"go.uber.org/zap"
)

const (
	getURI    = "/escalation_policies/%s"
	listURI   = "/escalation_policies"
	addURI    = "/escalation_policies"
	updateURI = "/escalation_policies/%s"
)

var ErrNoSuchPolicy = errors.New("no such escalation policy")

func New(cfg Config, logger *zap.Logger) *Manager {
	return &Manager{
		cfg:    cfg,
		logger: logger,
		api:    pd.New(cfg, logger),
	}
}

// Manager allows for loading and creating escalations
type Manager struct {
	cfg    Config
	logger *zap.Logger
	api    *pd.API
}

func (u *Manager) Get(ctx context.Context, policyID string) (*EscalationPolicy, error) {
	uri := fmt.Sprintf(getURI, policyID)

	escalations := &getEscalationsResponse{}

	err := u.api.Get(ctx, uri, nil, escalations)
	if err != nil {
		return nil, fmt.Errorf("failed to get escalation policy '%s' with err: %s", policyID, err)
	}

	if escalations.Policy == nil {
		return nil, ErrNoSuchPolicy
	}

	return escalations.Policy, nil
}

func (u *Manager) GetByName(ctx context.Context, name string) (*EscalationPolicy, error) {
	params := url.Values{}
	params.Set("query", name)
	params.Set("total", "false")
	params.Set("limit", "1")

	escalations := &getEscalationPolicyResponse{}

	err := u.api.Get(ctx, listURI, params, escalations)
	if err != nil {
		return nil, fmt.Errorf("failed to get escalations '%s' with err: %s", name, err)
	}

	if len(escalations.Policies) == 0 {
		return nil, ErrNoSuchPolicy
	}

	return escalations.Policies[0], nil
}

func (u *Manager) Add(ctx context.Context, policy NewPolicy) (string, error) {
	reqDTO := buildAddRequest(policy)
	addLeads(policy, reqDTO)
	addDeptHeads(policy, reqDTO)

	respDTO := &addResponse{}

	err := u.api.Post(ctx, addURI, reqDTO, respDTO)
	if err != nil {
		return "", fmt.Errorf("failed to add policy '%#v' with err: %s", reqDTO, err)
	}

	return respDTO.Policy.ID, nil
}

func (u *Manager) Update(ctx context.Context, escalationID string, policy NewPolicy) error {
	prevPolicy, err := u.Get(ctx, escalationID)
	if err != nil {
		return fmt.Errorf("failed to update escalation with err: %w", err)
	}

	reqDTO := buildAddRequest(policy)
	addLeads(policy, reqDTO)
	addDeptHeads(policy, reqDTO)

	updateIDs(reqDTO, prevPolicy)

	uri := fmt.Sprintf(updateURI, escalationID)

	err = u.api.Put(ctx, uri, reqDTO)
	if err != nil {
		return fmt.Errorf("failed to update policy '%#v' with err: %s", reqDTO, err)
	}

	return nil
}

func updateIDs(reqDTO *addRequest, prevPolicy *EscalationPolicy) {
	reqDTO.Policy.ID = prevPolicy.ID
}

func buildAddRequest(policy NewPolicy) *addRequest {
	return &addRequest{
		Policy: &EscalationPolicy{
			Name: policy.GetTeamName() + " Escalation",
			EscalationRules: []*escalationRule{
				{
					EscalationDelayInMinutes: 5,
					Targets: []*escalationTarget{
						{
							ID:   policy.GetScheduleID(),
							Type: "schedule_reference",
						},
					},
				},
			},
			NumLoops: 9,
			Teams: []*team{
				{
					ID:   policy.GetTeamID(),
					Type: "team_reference",
				},
			},
			OnCallHandoffNotifications: "always",
			Description:                policy.GetDescription(),
		},
	}
}

func addLeads(policy NewPolicy, reqDTO *addRequest) {
	if len(policy.GetLeadIDs()) == 0 {
		return
	}

	rule := &escalationRule{
		EscalationDelayInMinutes: 5,
	}

	for _, userID := range policy.GetLeadIDs() {
		rule.Targets = append(rule.Targets, &escalationTarget{
			ID:   userID,
			Type: "user_reference",
		})
	}

	reqDTO.Policy.EscalationRules = append(reqDTO.Policy.EscalationRules, rule)
}

func addDeptHeads(policy NewPolicy, reqDTO *addRequest) {
	if len(policy.GetDeptHeadsIDs()) == 0 {
		return
	}

	rule := &escalationRule{
		EscalationDelayInMinutes: 5,
	}

	for _, userID := range policy.GetDeptHeadsIDs() {
		rule.Targets = append(rule.Targets, &escalationTarget{
			ID:   userID,
			Type: "user_reference",
		})
	}

	reqDTO.Policy.EscalationRules = append(reqDTO.Policy.EscalationRules, rule)
}

type NewPolicy interface {
	GetTeamName() string
	GetDescription() string
	GetScheduleID() string
	GetTeamID() string
	GetLeadIDs() []string
	GetDeptHeadsIDs() []string
}

type getEscalationsResponse struct {
	Policy *EscalationPolicy `json:"escalation_policy"`
}

type getEscalationPolicyResponse struct {
	Policies []*EscalationPolicy `json:"escalation_policies"`
}

type EscalationPolicy struct {
	ID                         string            `json:"id"`
	Name                       string            `json:"name"`
	EscalationRules            []*escalationRule `json:"escalation_rules"`
	NumLoops                   int               `json:"num_loops"`
	Teams                      []*team           `json:"teams"`
	OnCallHandoffNotifications string            `json:"on_call_handoff_notifications"`
	Description                string            `json:"description"`
}

type escalationRule struct {
	EscalationDelayInMinutes int                 `json:"escalation_delay_in_minutes"`
	Targets                  []*escalationTarget `json:"targets"`
}

type escalationTarget struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type team struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type addRequest struct {
	Policy *EscalationPolicy `json:"escalation_policy"`
}

type addResponse struct {
	Policy *EscalationPolicy `json:"escalation_policy"`
}

type Config interface {
	Debug() bool
	BaseURL() string
	AuthToken() string
}
