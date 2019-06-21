// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
)

// NewLogger returns a logger instance
func NewLogger(level, encoding string, outputPaths []string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	rawJSON := []byte(fmt.Sprintf(`{
      		"level": "%s",
      		"encoding": "%s",
      		"outputPaths": []
    	}`, level, encoding))

	err := json.Unmarshal(rawJSON, &cfg)

	if err != nil {
		panic(err)
	}

	cfg.Encoding = encoding
	cfg.OutputPaths = outputPaths

	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	return logger, nil
}

// PathExists reports whether the path exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileExists reports whether the named file exists
func FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

// DirExists reports whether the dir exists
func DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

// EnsureDir ensures that directory exists
func EnsureDir(dirName string, mode int) (bool, error) {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return true, nil
	}
	return false, err
}
