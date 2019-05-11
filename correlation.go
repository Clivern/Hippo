// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"github.com/satori/go.uuid"
)

// Correlation struct
type Correlation struct {
}

// UUIDv4 create a UUID version 4
func (c *Correlation) UUIDv4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}

// ExampleUUIDv4 example for UUIDv4 method
func ExampleUUIDv4() {
	c := Correlation{}
	fmt.Println(c.UUIDv4() != "")
	// Output: true
}
