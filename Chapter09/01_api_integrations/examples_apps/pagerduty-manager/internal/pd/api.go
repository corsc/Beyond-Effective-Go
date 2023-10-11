package pd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	"github.com/corsc/go-commons/iocloser"
)

func New(cfg Config, logger *zap.Logger) *API {
	return &API{
		cfg:    cfg,
		logger: logger,
		client: &http.Client{},
	}
}

// API encapsulates the REST calls to the API
type API struct {
	cfg    Config
	logger *zap.Logger
	client *http.Client
}

func (u *API) Get(ctx context.Context, uri string, params url.Values, respDTO interface{}) error {
	fullURI := u.buildURI(uri, params)

	u.logger.Debug("making HTTP GET request", zap.String("uri", fullURI))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURI, nil)
	if err != nil {
		return fmt.Errorf("failed to build GET request with err: %w", err)
	}

	req.Header.Set("Authorization", "Token token="+u.cfg.AuthToken())
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do GET request with err: %w", err)
	}

	defer iocloser.Close(resp.Body)

	if resp.StatusCode != http.StatusOK {
		payload, _ := io.ReadAll(resp.Body)
		u.logger.Debug("response", zap.ByteString("payload", payload))

		return fmt.Errorf("unexpected HTTP GET response code: %d", resp.StatusCode)
	}

	return u.parseResponse(resp, respDTO)
}

func (u *API) Put(ctx context.Context, uri string, reqDTO interface{}) error {
	fullURI := u.buildURI(uri, nil)

	payload, err := json.Marshal(reqDTO)
	if err != nil {
		return fmt.Errorf("failed to build PUT request payload with err: %w", err)
	}

	u.logger.Debug("making HTTP PUT request", zap.String("uri", fullURI))

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, fullURI, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to build PUT request with err: %w", err)
	}

	req.Header.Set("Authorization", "Token token="+u.cfg.AuthToken())
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req) //nolint:bodyclose
	if err != nil {
		return fmt.Errorf("failed to do PUT request with err: %w", err)
	}

	defer iocloser.Close(resp.Body)

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		payload, _ := io.ReadAll(resp.Body)
		u.logger.Debug("response", zap.ByteString("payload", payload))

		return fmt.Errorf("unexpected HTTP PUT response code: %d", resp.StatusCode)
	}

	return nil
}

func (u *API) Post(ctx context.Context, uri string, reqDTO, respDTO interface{}) error {
	fullURI := u.buildURI(uri, nil)

	payload, err := json.Marshal(reqDTO)
	if err != nil {
		return fmt.Errorf("failed to build POST request payload with err: %w", err)
	}

	u.logger.Debug("making HTTP POST request", zap.String("uri", fullURI))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURI, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to build PUT request with err: %w", err)
	}

	req.Header.Set("Authorization", "Token token="+u.cfg.AuthToken())
	req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Set("Content-Type", "application/json")

	resp, err := u.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do PUT request with err: %w", err)
	}

	defer iocloser.Close(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(resp.Body)
		u.logger.Debug("response", zap.ByteString("payload", payload))

		return fmt.Errorf("unexpected HTTP PUT response code: %d", resp.StatusCode)
	}

	return u.parseResponse(resp, respDTO)
}

func (u *API) buildURI(uri string, params url.Values) string {
	resultURI := u.cfg.BaseURL() + uri

	if len(params) != 0 {
		resultURI += "?" + params.Encode()
	}

	return resultURI
}

func (u *API) parseResponse(resp *http.Response, respDTO interface{}) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body with err: %w", err)
	}

	u.logger.Debug("response", zap.ByteString("payload", payload))

	err = json.Unmarshal(payload, respDTO)
	if err != nil {
		return fmt.Errorf("failed to read response JSON with err: %w", err)
	}

	return nil
}

type Config interface {
	Debug() bool
	BaseURL() string
	AuthToken() string
}
