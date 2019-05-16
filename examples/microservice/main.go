// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/clivern/hippo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	correlation := hippo.NewCorrelation()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "pong",
			"correlation": correlation.UUIDv4(),
		})
	})
	r.Run()
}
