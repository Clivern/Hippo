// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP struct
type HTTP struct {
}

// Get http call
func (h *HTTP) Get(url string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)

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
func (h *HTTP) Delete(url string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("DELETE", url, nil)

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

// ToString response body to string
func (h *HTTP) ToString(response *http.Response) (string, error) {

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
	respObj, _ := http.Get("https://httpbin.org/get", map[string]string{"X-AUTH-Token": "123"})
	fmt.Println(http.GetStatusCode(respObj))
	// Output: 200
}

// ExampleDelete example for Delete method
func ExampleDelete() {
	http := HTTP{}
	respObj, _ := http.Delete("https://httpbin.org/delete", map[string]string{"X-AUTH-Token": "123"})
	fmt.Println(http.GetStatusCode(respObj))
	// Output: 200
}
