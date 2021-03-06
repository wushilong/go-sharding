// Copyright 2019 The Gaea Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package source

import (
	"errors"
	"github.com/XiaoMi/Gaea/logging"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	defaultFilePath = "./etc/file"
)

// File source provider for configuration
type fileSource struct {
	Prefix string
}

// New constructor of etcdSource
func NewFileSource(path string) (*fileSource, error) {
	if strings.TrimSpace(path) == "" {
		path = defaultFilePath
	}
	if err := checkDir(path); err != nil {
		logging.DefaultLogger.Warnf("check file source directory failed, %v", err)
		return nil, err
	}
	return &fileSource{Prefix: path}, nil
}

func checkDir(path string) error {
	if strings.TrimSpace(path) == "" {
		return errors.New("invalid path")
	}
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return errors.New("invalid path, should be a directory")
	}

	return nil
}

// Close do nothing
func (c *fileSource) Close() error {
	return nil
}

// Create do nothing
func (c *fileSource) Create(path string, data []byte) error {
	return nil
}

// Update do nothing
func (c *fileSource) Update(path string, data []byte) error {
	return nil
}

// UpdateWithTTL update path with data and ttl
func (c *fileSource) UpdateWithTTL(path string, data []byte, ttl time.Duration) error {
	return nil
}

// Delete delete path
func (c *fileSource) Delete(path string) error {
	return nil
}

// Read read file data
func (c *fileSource) Read(file string) ([]byte, error) {
	value, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// List list path, return slice of all files
func (c *fileSource) List(path string) ([]string, error) {
	r := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return r, err
	}

	for _, f := range files {
		r = append(r, f.Name())
	}

	return r, nil
}

// BasePrefix return base prefix
func (c *fileSource) BasePrefix() string {
	return c.Prefix
}
