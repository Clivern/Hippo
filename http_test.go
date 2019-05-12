// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"github.com/nbio/st"
	"net/http"
	"strings"
	"testing"
)

// TestGet test cases
func TestGet(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Get(
		"https://httpbin.org/get",
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusOK, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, true, strings.Contains(body, "value1"))
	st.Expect(t, true, strings.Contains(body, "arg1"))
	st.Expect(t, true, strings.Contains(body, "arg1=value1"))
	st.Expect(t, true, strings.Contains(body, "X-Auth"))
	st.Expect(t, true, strings.Contains(body, "hipp-123"))
	st.Expect(t, nil, error)
}

// TestDelete test cases
func TestDelete(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Delete(
		"https://httpbin.org/delete",
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusOK, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, true, strings.Contains(body, "value1"))
	st.Expect(t, true, strings.Contains(body, "arg1"))
	st.Expect(t, true, strings.Contains(body, "arg1=value1"))
	st.Expect(t, true, strings.Contains(body, "X-Auth"))
	st.Expect(t, true, strings.Contains(body, "hipp-123"))
	st.Expect(t, nil, error)
}

// TestPost test cases
func TestPost(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Post(
		"https://httpbin.org/post",
		`{"Username":"admin", "Password":"12345"}`,
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusOK, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, true, strings.Contains(body, `"12345"`))
	st.Expect(t, true, strings.Contains(body, `"Username"`))
	st.Expect(t, true, strings.Contains(body, `"admin"`))
	st.Expect(t, true, strings.Contains(body, `"Password"`))
	st.Expect(t, true, strings.Contains(body, "value1"))
	st.Expect(t, true, strings.Contains(body, "arg1"))
	st.Expect(t, true, strings.Contains(body, "arg1=value1"))
	st.Expect(t, true, strings.Contains(body, "X-Auth"))
	st.Expect(t, true, strings.Contains(body, "hipp-123"))
	st.Expect(t, nil, error)
}

// TestPut test cases
func TestPut(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Put(
		"https://httpbin.org/put",
		`{"Username":"admin", "Password":"12345"}`,
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusOK, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, true, strings.Contains(body, `"12345"`))
	st.Expect(t, true, strings.Contains(body, `"Username"`))
	st.Expect(t, true, strings.Contains(body, `"admin"`))
	st.Expect(t, true, strings.Contains(body, `"Password"`))
	st.Expect(t, true, strings.Contains(body, "value1"))
	st.Expect(t, true, strings.Contains(body, "arg1"))
	st.Expect(t, true, strings.Contains(body, "arg1=value1"))
	st.Expect(t, true, strings.Contains(body, "X-Auth"))
	st.Expect(t, true, strings.Contains(body, "hipp-123"))
	st.Expect(t, nil, error)
}

// TestGetStatusCode1 test cases
func TestGetStatusCode1(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Get(
		"https://httpbin.org/status/200",
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusOK, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, "", body)
	st.Expect(t, nil, error)
}

// TestGetStatusCode2 test cases
func TestGetStatusCode2(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Get(
		"https://httpbin.org/status/500",
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusInternalServerError, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, "", body)
	st.Expect(t, nil, error)
}

// TestGetStatusCode3 test cases
func TestGetStatusCode3(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Get(
		"https://httpbin.org/status/404",
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusNotFound, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, "", body)
	st.Expect(t, nil, error)
}

// TestGetStatusCode4 test cases
func TestGetStatusCode4(t *testing.T) {
	httpClient := HTTP{}
	response, error := httpClient.Get(
		"https://httpbin.org/status/201",
		map[string]string{"arg1": "value1"},
		map[string]string{"X-Auth": "hipp-123"},
	)
	t.Log(httpClient.GetStatusCode(response))
	st.Expect(t, http.StatusCreated, httpClient.GetStatusCode(response))
	st.Expect(t, nil, error)

	body, error := httpClient.ToString(response)
	t.Log(body)
	st.Expect(t, "", body)
	st.Expect(t, nil, error)
}
