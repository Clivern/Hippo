// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"fmt"
	"github.com/nbio/st"
	"testing"
	"time"
)

// TestRedis test cases
func TestRedis(t *testing.T) {

	driver := NewRedisDriver("10.11.120.60:7001", "", 0)

	ok, err := driver.Connect()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.Ping()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	// Do Clean
	driver.Del("app_name")
	driver.HTruncate("configs")

	count, err := driver.Del("app_name")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)

	ok, err = driver.Set("app_name", "Hippo", 0)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.Exists("app_name")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	value, err := driver.Get("app_name")
	st.Expect(t, value, "Hippo")
	st.Expect(t, err, nil)

	count, err = driver.HDel("configs", "app_name")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)

	ok, err = driver.HSet("configs", "app_name", "Hippo")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.HExists("configs", "app_name")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	value, err = driver.HGet("configs", "app_name")
	st.Expect(t, value, "Hippo")
	st.Expect(t, err, nil)

	count, err = driver.HLen("configs")
	st.Expect(t, int(count), 1)
	st.Expect(t, err, nil)

	count, err = driver.HDel("configs", "app_name")
	st.Expect(t, int(count), 1)
	st.Expect(t, err, nil)

	count, err = driver.HTruncate("configs")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)

	c := make(chan string)

	go func() {
		c <- "Hello World"
		driver.Subscribe("hippo", func(message Message) error {
			t.Log(message.Channel)
			t.Log(message.Payload)
			st.Expect(t, "hippo", message.Channel)
			st.Expect(t, "Hello World", message.Payload)
			return fmt.Errorf("Terminate listener")
		})
	}()

	msg := <-c
	time.Sleep(4 * time.Second)
	driver.Publish("hippo", msg)
}
