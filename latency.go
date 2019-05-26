// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"time"
)

// Point struct
type Point struct {
	Start time.Time
	End   time.Time
}

// Latency struct
type Latency struct {
	Actions map[string][]Point
}

// NewLatencyTracker creates a new latency instance
func NewLatencyTracker() *Latency {
	return &Latency{}
}

// NewAction creates a new action tracking bucket
func (l *Latency) NewAction(name string) {
	l.Actions = make(map[string][]Point)
	l.Actions[name] = []Point{}
}

// SetPoint adds a new point
func (l *Latency) SetPoint(name string, start, end time.Time) {
	if _, ok := l.Actions[name]; !ok {
		l.NewAction(name)
	}
	l.Actions[name] = append(l.Actions[name], Point{Start: start, End: end})
}

// SetStart adds point start time
func (l *Latency) SetStart(name string, start time.Time) bool {
	if _, ok := l.Actions[name]; !ok {
		l.NewAction(name)
	}
	l.Actions[name] = append(l.Actions[name], Point{Start: start})

	return true
}

// SetEnd adds point end time
func (l *Latency) SetEnd(name string, end time.Time) bool {
	if _, ok := l.Actions[name]; !ok {
		l.NewAction(name)
	}

	length := len(l.Actions[name])

	if length <= 0 {
		return false
	}

	if l.Actions[name][length-1].End.String() == "" {
		return false
	}

	l.Actions[name][length-1].End = end

	return true
}
