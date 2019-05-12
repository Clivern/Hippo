// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"github.com/nbio/st"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestStatusGetRaftLeader test cases
func TestStatusGetRaftLeader(t *testing.T) {
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

// TestStatusListRaftPeers test cases
func TestStatusListRaftPeers(t *testing.T) {
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

// TestKvRead test cases
func TestKvRead(t *testing.T) {
	version := "v1"
	endpoint := "kv"
	item := "key"

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == fmt.Sprintf("/%s/%s/%s", version, endpoint, item) {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`true`))
			}
		}),
	)
	defer ts.Close()

	t.Log(fmt.Sprintf("%s/%s/%s/%s", ts.URL, version, endpoint, item))

	config := ConsulConfig{
		URL:     ts.URL,
		Version: version,
	}

	kv := ConsulKv{
		Config: config,
	}

	body, error := kv.Read("key", map[string]string{})

	t.Log(body)
	t.Log(error)
	st.Expect(t, "true", body)
	st.Expect(t, nil, error)
}

// TestKvUpdate test cases
func TestKvUpdate(t *testing.T) {
	version := "v1"
	endpoint := "kv"
	item := "key"

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == fmt.Sprintf("/%s/%s/%s", version, endpoint, item) {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				body, _ := ioutil.ReadAll(r.Body)
				w.Write([]byte(string(body)))
			}
		}),
	)
	defer ts.Close()

	t.Log(fmt.Sprintf("%s/%s/%s/%s", ts.URL, version, endpoint, item))

	config := ConsulConfig{
		URL:     ts.URL,
		Version: version,
	}

	kv := ConsulKv{
		Config: config,
	}

	body, error := kv.Update("key", "value", map[string]string{})

	t.Log(body)
	t.Log(error)
	st.Expect(t, "value", body)
	st.Expect(t, nil, error)
}

// TestKvDelete test cases
func TestKvDelete(t *testing.T) {
	version := "v1"
	endpoint := "kv"
	item := "key"

	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == fmt.Sprintf("/%s/%s/%s", version, endpoint, item) {
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`true`))
			}
		}),
	)
	defer ts.Close()

	t.Log(fmt.Sprintf("%s/%s/%s/%s", ts.URL, version, endpoint, item))

	config := ConsulConfig{
		URL:     ts.URL,
		Version: version,
	}

	kv := ConsulKv{
		Config: config,
	}

	body, error := kv.Delete("key", map[string]string{})

	t.Log(body)
	t.Log(error)
	st.Expect(t, "true", body)
	st.Expect(t, nil, error)
}
