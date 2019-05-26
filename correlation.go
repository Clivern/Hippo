// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"github.com/satori/go.uuid"
)

// Correlation interface
type Correlation interface {
	UUIDv4() string
}

// Correlation struct
type correlation struct {
}

// NewCorrelation creates an instance of correlation struct
func NewCorrelation() Correlation {
	c := &correlation{}
	return c
}

// UUIDv4 create a UUID version 4
func (c *correlation) UUIDv4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}
