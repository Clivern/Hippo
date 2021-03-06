// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// ConsulConfig struct
type ConsulConfig struct {
	URL     string
	Version string
}

// ConsulStatus struct
type ConsulStatus struct {
	Config ConsulConfig
}

// ConsulKv struct
type ConsulKv struct {
	Config ConsulConfig
}

// GetRaftLeader returns the Raft leader for the datacenter in which the agent is running
func (c *ConsulStatus) GetRaftLeader(ctx context.Context, parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/status/leader",
		strings.TrimSuffix(c.Config.URL, "/"),
		c.Config.Version,
	)

	httpClient := NewHTTPClient()

	response, err := httpClient.Get(ctx, endpoint, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) != http.StatusOK {
		return "", fmt.Errorf("Error: Invalid HTTP status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}

// ListRaftPeers retrieves the Raft peers for the datacenter in which the the agent is running
func (c *ConsulStatus) ListRaftPeers(ctx context.Context, parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/status/peers",
		strings.TrimSuffix(c.Config.URL, "/"),
		c.Config.Version,
	)

	httpClient := NewHTTPClient()

	response, err := httpClient.Get(ctx, endpoint, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) != http.StatusOK {
		return "", fmt.Errorf("Error: Invalid HTTP status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}

// Read gets a kv
func (c *ConsulKv) Read(ctx context.Context, key string, parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/kv/%s",
		strings.TrimSuffix(c.Config.URL, "/"),
		c.Config.Version,
		key,
	)

	httpClient := NewHTTPClient()

	response, err := httpClient.Get(ctx, endpoint, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) == http.StatusNotFound {
		return "", fmt.Errorf("Error: Key [%s] does not exist", key)
	}

	if httpClient.GetStatusCode(response) != http.StatusOK {
		return "", fmt.Errorf("Error: Invalid HTTP status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}

// Update update or create a kv
func (c *ConsulKv) Update(ctx context.Context, key string, value string, parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/kv/%s",
		strings.TrimSuffix(c.Config.URL, "/"),
		c.Config.Version,
		key,
	)

	httpClient := NewHTTPClient()

	response, err := httpClient.Post(ctx, endpoint, value, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) != http.StatusOK {
		return "", fmt.Errorf("Error: Invalid HTTP status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}

// Delete deletes a kv
func (c *ConsulKv) Delete(ctx context.Context, key string, parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/kv/%s",
		strings.TrimSuffix(c.Config.URL, "/"),
		c.Config.Version,
		key,
	)

	httpClient := NewHTTPClient()

	response, err := httpClient.Get(ctx, endpoint, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) == http.StatusNotFound {
		return "", fmt.Errorf("Error: Key [%s] does not exist", key)
	}

	if httpClient.GetStatusCode(response) != http.StatusOK {
		return "", fmt.Errorf("Error: Invalid HTTP status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}
