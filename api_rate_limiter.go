// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

// caller struct
type caller struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// callers list
var callers = make(map[string]*caller)

// mtx mutex
var mtx sync.Mutex

// NewCallerLimiter create a new rate limiter with an identifier
func NewCallerLimiter(identifier string, eventsRate rate.Limit, tokenBurst int) *rate.Limiter {
	mtx.Lock()
	v, exists := callers[identifier]
	if !exists {
		mtx.Unlock()
		return addCaller(identifier, eventsRate, tokenBurst)
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	mtx.Unlock()
	return v.limiter
}

// addCaller add a caller
func addCaller(identifier string, eventsRate rate.Limit, tokenBurst int) *rate.Limiter {
	limiter := rate.NewLimiter(eventsRate, tokenBurst)
	mtx.Lock()
	// Include the current time when creating a new visitor.
	callers[identifier] = &caller{limiter, time.Now()}
	mtx.Unlock()

	return limiter
}

// CleanupCallers cleans old clients
func CleanupCallers(cleanAfter time.Duration) {
	mtx.Lock()
	for identifier, v := range callers {
		if time.Now().Sub(v.lastSeen) > cleanAfter*time.Second {
			delete(callers, identifier)
		}
	}
	mtx.Unlock()
}
