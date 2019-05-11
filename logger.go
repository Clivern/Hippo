// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"go.uber.org/zap"
)

// Logger struct
type Logger struct {
}

// Create returns a logger instance
func (L *Logger) Create() (*zap.Logger, error) {
	return zap.NewProduction()
}

// Info logs info message
func (L *Logger) Info(v ...interface{}) {
	logger, _ := L.Create()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info(v...)
}
