// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// ServiceUp const
	ServiceUp = "UP"
	// ServiceDown const
	ServiceDown = "DOWN"
	// ServiceUnknown const
	ServiceUnknown = "UNKNOWN"
)

// Check struct
type Check struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Error    string `json:"error"`
	Result   bool   `json:"result"`
	callable func() (bool, error)
}

// Health struct
type Health struct {
	Status string
	Checks []*Check
}

// NewHealthChecker initializes a new health checker
func NewHealthChecker() *Health {
	return &Health{}
}

// IsUnknown returns true if Status is Unknown
func (h *Health) IsUnknown() bool {
	return h.Status == ServiceUnknown
}

// IsUp returns true if Status is Up
func (h *Health) IsUp() bool {
	return h.Status == ServiceUp
}

// IsDown returns true if Status is Down
func (h *Health) IsDown() bool {
	return h.Status == ServiceDown
}

// Down set the Status to Down
func (h *Health) Down() *Health {
	h.Status = ServiceDown
	return h
}

// Unknown set the Status to Unknown
func (h *Health) Unknown() *Health {
	h.Status = ServiceUnknown
	return h
}

// Up set the Status to Up
func (h *Health) Up() *Health {
	h.Status = ServiceUp
	return h
}

// ChecksStatus get checks Status
func (h *Health) ChecksStatus() string {
	return h.Status
}

// ChecksReport get checks Status
func (h *Health) ChecksReport() (string, error) {
	bytes, err := json.Marshal(h.Checks)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// AddCheck adds a new check
func (h *Health) AddCheck(ID string, callable func() (bool, error)) {
	check := &Check{
		ID:       ID,
		Status:   ServiceUnknown,
		callable: callable,
	}
	h.Checks = append(h.Checks, check)
}

// RunChecks runs all health checks
func (h *Health) RunChecks() {
	upCount := 0
	downCount := 0
	var err error
	for _, check := range h.Checks {
		check.Result, err = check.callable()
		if err != nil {
			check.Error = err.Error()
		}
		if check.Result {
			check.Status = ServiceUp
			upCount++
		} else {
			check.Status = ServiceDown
			downCount++
		}
	}
	if downCount > 0 {
		h.Down()
	} else {
		h.Up()
	}
}

// HTTPCheck do HTTP health check
func HTTPCheck(ctx context.Context, serviceName, URL string, parameters map[string]string, headers map[string]string) (bool, error) {
	httpClient := NewHTTPClient()
	response, error := httpClient.Get(
		ctx,
		URL,
		parameters,
		headers,
	)

	if error != nil {
		return false, error
	}

	if httpClient.GetStatusCode(response) == http.StatusServiceUnavailable {
		return false, fmt.Errorf("Service %s is unavailable", serviceName)
	}

	return true, nil
}

// RedisCheck do a redis health check
func RedisCheck(serviceName string, addr string, password string, db int) (bool, error) {
	redisDriver := NewRedisDriver(addr, password, db)
	_, err := redisDriver.Connect()

	if err != nil {
		return false, fmt.Errorf("Error while connecting %s: %s", serviceName, err.Error())
	}

	status, err := redisDriver.Ping()

	if err != nil {
		return false, fmt.Errorf("Error while connecting %s: %s", serviceName, err.Error())
	}

	return status, nil
}
