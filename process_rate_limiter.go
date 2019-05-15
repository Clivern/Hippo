// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"sync"
	"time"
)

// ProcessLimiter interface
type ProcessLimiter interface {
	// Take should block to make sure that the RPS is met.
	Take() time.Time
}

// Clock Type
type Clock struct {
}

// Now get current time
func (c *Clock) Now() time.Time {
	return time.Now()
}

// Sleep sleeps for a time
func (c *Clock) Sleep(d time.Duration) {
	time.Sleep(d)
}

// processLimiter type
type processLimiter struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
	maxSlack   time.Duration
	clock      Clock
}

// NewProcessLimiter create a new process rate limiter
func NewProcessLimiter(rate int) ProcessLimiter {
	l := &processLimiter{
		perRequest: time.Second / time.Duration(rate),
		maxSlack:   -10 * time.Second / time.Duration(rate),
	}
	l.clock = Clock{}

	return l
}

// Take sleep for time to limit requests
func (t *processLimiter) Take() time.Time {
	t.Lock()
	defer t.Unlock()

	now := t.clock.Now()

	// If this is our first request, then we allow it.
	if t.last.IsZero() {
		t.last = now
		return t.last
	}

	// sleepFor calculates how much time we should sleep based on
	// the perRequest budget and how long the last request took.
	// Since the request may take longer than the budget, this number
	// can get negative, and is summed across requests.
	t.sleepFor += t.perRequest - now.Sub(t.last)

	// We shouldn't allow sleepFor to get too negative, since it would mean that
	// a service that slowed down a lot for a short period of time would get
	// a much higher RPS following that.
	if t.sleepFor < t.maxSlack {
		t.sleepFor = t.maxSlack
	}

	// If sleepFor is positive, then we should sleep now.
	if t.sleepFor > 0 {
		t.clock.Sleep(t.sleepFor)
		t.last = now.Add(t.sleepFor)
		t.sleepFor = 0
	} else {
		t.last = now
	}

	return t.last
}
