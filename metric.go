// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"net"
	"time"
)

// Graphite is a struct that defines the relevant properties of a graphite
type Graphite struct {
	Host     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Prefix   string
	conn     net.Conn
	nop      bool
}

// defaultTimeout is the default number of seconds that we're willing to wait
const defaultTimeout = 5

// IsNop is a getter for *graphite.Graphite.nop
func (graphite *Graphite) IsNop() bool {
	if graphite.nop {
		return true
	}
	return false
}
