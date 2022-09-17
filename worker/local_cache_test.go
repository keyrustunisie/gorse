// Copyright 2021 gorse Project Authors
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

package worker

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLocalCache(t *testing.T) {
	// delete test file if exists
	path := filepath.Join(os.TempDir(), "TestLocalCache_Worker")
	_ = os.RemoveAll(path)
	// load non-existed file
	cache, err := LoadLocalCache(path)
	assert.Error(t, err)
	assert.Equal(t, path, cache.folderPath)
	assert.Equal(t, filepath.Join(path, "meta.json"), cache.metaPath)
	assert.Empty(t, cache.meta.WorkerName)
	// write and load
	cache.meta.WorkerName = "Worker"
	assert.NoError(t, cache.WriteLocalCache())
	read, err := LoadLocalCache(path)
	assert.NoError(t, err)
	assert.Equal(t, "Worker", read.meta.WorkerName)
	// delete test file
	assert.NoError(t, os.RemoveAll(path))
}
