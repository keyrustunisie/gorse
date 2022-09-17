// Copyright 2020 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"
	std_errors "errors"
	"github.com/juju/errors"
	"io"
	"os"
	"path/filepath"
)

type MetaCache struct {
	ServerName string `json:"server_name"`
}

// LocalCache is local cache for the server node.
type LocalCache struct {
	folderPath string
	metaPath   string
	meta       MetaCache
}

// LoadLocalCache loads local cache from a file.
func LoadLocalCache(path string) (*LocalCache, error) {
	state := &LocalCache{
		folderPath: path,
		metaPath:   filepath.Join(path, "meta.json"),
	}
	// check if file exists
	if _, err := os.Stat(path); err != nil {
		if std_errors.Is(err, os.ErrNotExist) {
			return state, errors.NotFoundf("local cache file %s", path)
		}
		return state, errors.Trace(err)
	}
	// open file
	f, err := os.Open(state.metaPath)
	if err != nil {
		return state, errors.Trace(err)
	}
	defer f.Close()
	metaData, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if err = json.Unmarshal(metaData, &state.meta); err != nil {
		return nil, errors.Trace(err)
	}
	return state, nil
}

// WriteLocalCache writes local cache to a file.
func (s *LocalCache) WriteLocalCache() error {
	// create folder if not exists
	if _, err := os.Stat(s.folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(s.folderPath, os.ModePerm)
		if err != nil {
			return errors.Trace(err)
		}
	}
	// create file
	f, err := os.Create(s.metaPath)
	if err != nil {
		return errors.Trace(err)
	}
	defer f.Close()
	// write file
	metaData, err := json.Marshal(s.meta)
	if err != nil {
		return errors.Trace(err)
	}
	_, err = f.Write(metaData)
	return errors.Trace(err)
}
