// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"github.com/nbio/st"
	"testing"
)

// TestHealthCheck test cases
func TestHealthCheck(t *testing.T) {

	healthChecker := NewHealthChecker()
	healthChecker.AddCheck("ping_check", func() (bool, error) {
		return true, nil
	})
	healthChecker.AddCheck("db_check", func() (bool, error) {
		return false, fmt.Errorf("Database Down")
	})
	healthChecker.RunChecks()

	st.Expect(t, healthChecker.ChecksStatus(), "DOWN")

	report, err := healthChecker.ChecksReport()

	st.Expect(t, report, `[{"id":"ping_check","status":"UP","error":"","result":true},{"id":"db_check","status":"DOWN","error":"Database Down","result":false}]`)
	st.Expect(t, err, nil)

	st.Expect(t, healthChecker.IsDown(), true)
	st.Expect(t, healthChecker.IsUnknown(), false)
	st.Expect(t, healthChecker.IsUp(), false)

	healthChecker.Down()
	st.Expect(t, healthChecker.IsDown(), true)

	healthChecker.Unknown()
	st.Expect(t, healthChecker.IsUnknown(), true)

	healthChecker.Up()
	st.Expect(t, healthChecker.IsUp(), true)

}
