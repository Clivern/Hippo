// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"strings"
)

// Consul struct
type Consul struct {
	URL     string
	Version string
}

// GetRaftLeader returns the Raft leader for the datacenter in which the agent is running
func (c *Consul) GetRaftLeader(parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/status/leader",
		strings.TrimSuffix(c.URL, "/"),
		c.Version,
	)

	httpClient := HTTP{}

	response, err := httpClient.Get(endpoint, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) != 200 {
		return "", fmt.Errorf("Invalid http status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}

// ListRaftPeers retrieves the Raft peers for the datacenter in which the the agent is running
func (c *Consul) ListRaftPeers(parameters map[string]string) (string, error) {
	endpoint := fmt.Sprintf(
		"%s/%s/status/peers",
		strings.TrimSuffix(c.URL, "/"),
		c.Version,
	)

	httpClient := HTTP{}

	response, err := httpClient.Get(endpoint, parameters, map[string]string{})

	if err != nil {
		return "", err
	}

	if httpClient.GetStatusCode(response) != 200 {
		return "", fmt.Errorf("Invalid http status code %d", httpClient.GetStatusCode(response))
	}

	body, err := httpClient.ToString(response)

	if err != nil {
		return "", err
	}

	return body, nil
}
