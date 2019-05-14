// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"go.uber.org/zap"
)

// StdoutDriver the stdout driver
const StdoutDriver = "STDOUT"

// FileDriver the file driver
const FileDriver = "FILE"

// ElkDriver the elk driver
const ElkDriver = "ELK"

// Logger struct
type Logger struct {
	Driver string
}

// New returns a logger instance
func (l *Logger) New() (*zap.Logger, error) {
	return zap.NewProduction()
}

// Info logs info message
func (l *Logger) Info(v ...interface{}) {
	logger, _ := l.New()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info(v...)
}
