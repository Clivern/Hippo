// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetRaftLeader test cases
func TestGetRaftLeader(t *testing.T) {
	version := "v1"
	endpoint := "status/leader"

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == fmt.Sprintf("/%s/%s", version, endpoint) {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`true`))
			}
		}),
	)
	defer ts.Close()

	t.Log(fmt.Sprintf("%s/%s/%s", ts.URL, version, endpoint))

	config := ConsulConfig{
		URL:     ts.URL,
		Version: version,
	}

	status := ConsulStatus{
		Config: config,
	}

	body, error := status.GetRaftLeader(map[string]string{})

	t.Log(body)
	t.Log(error)
	st.Expect(t, "true", body)
	st.Expect(t, nil, error)
}

// TestListRaftPeers test cases
func TestListRaftPeers(t *testing.T) {
	version := "v1"
	endpoint := "status/peers"

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == fmt.Sprintf("/%s/%s", version, endpoint) {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`true`))
			}
		}),
	)
	defer ts.Close()

	t.Log(fmt.Sprintf("%s/%s/%s", ts.URL, version, endpoint))

	config := ConsulConfig{
		URL:     ts.URL,
		Version: version,
	}

	status := ConsulStatus{
		Config: config,
	}

	body, error := status.ListRaftPeers(map[string]string{})

	t.Log(body)
	t.Log(error)
	st.Expect(t, "true", body)
	st.Expect(t, nil, error)
}
