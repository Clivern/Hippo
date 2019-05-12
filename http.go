// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HTTP struct
type HTTP struct {
}

// Get http call
func (h *HTTP) Get(endpoint string, parameters map[string]string, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.BuildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", endpoint, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Post http call
func (h *HTTP) Post(endpoint string, data string, parameters map[string]string, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.BuildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(data)))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Put http call
func (h *HTTP) Put(endpoint string, data string, parameters map[string]string, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.BuildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer([]byte(data)))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// Delete http call
func (h *HTTP) Delete(endpoint string, parameters map[string]string, headers map[string]string) (*http.Response, error) {

	endpoint, err := h.BuildParameters(endpoint, parameters)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", endpoint, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	return resp, err
}

// BuildParameters add parameters to URL
func (h *HTTP) BuildParameters(endpoint string, parameters map[string]string) (string, error) {
	u, err := url.Parse(endpoint)

	if err != nil {
		return "", err
	}

	q := u.Query()

	for k, v := range parameters {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// ToString response body to string
func (h *HTTP) ToString(response *http.Response) (string, error) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetStatusCode response status code
func (h *HTTP) GetStatusCode(response *http.Response) int {
	return response.StatusCode
}

// ExampleGet example for Get method
func ExampleGet() {
	http := HTTP{}
	respObj, _ := http.Get("https://httpbin.org/get", map[string]string{}, map[string]string{"X-AUTH-Token": "123"})
	fmt.Println(http.GetStatusCode(respObj))
	// Output: 200
}

// ExampleDelete example for Delete method
func ExampleDelete() {
	http := HTTP{}
	respObj, _ := http.Delete("https://httpbin.org/delete", map[string]string{}, map[string]string{"X-AUTH-Token": "123"})
	fmt.Println(http.GetStatusCode(respObj))
	// Output: 200
}
