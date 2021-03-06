// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugofs

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestLanguagFs(t *testing.T) {
	languages := map[string]bool{
		"sv": true,
	}
	base := filepath.FromSlash("/my/base")
	assert := require.New(t)
	m := afero.NewMemMapFs()
	bfs := afero.NewBasePathFs(m, base)
	lfs := NewLanguageFs("sv", languages, bfs)
	assert.NotNil(lfs)
	assert.Equal("sv", lfs.Lang())
	err := afero.WriteFile(lfs, filepath.FromSlash("sect/page.md"), []byte("abc"), 0777)
	assert.NoError(err)
	fi, err := lfs.Stat(filepath.FromSlash("sect/page.md"))
	assert.NoError(err)
	assert.Equal("__hugofs_sv_page.md", fi.Name())

	languager, ok := fi.(LanguageAnnouncer)
	assert.True(ok)

	assert.Equal("sv", languager.Lang())

	lfi, ok := fi.(*LanguageFileInfo)
	assert.True(ok)
	assert.Equal(filepath.FromSlash("/my/base/sect/page.md"), lfi.Filename())
	assert.Equal(filepath.FromSlash("sect/page.md"), lfi.Path())
	assert.Equal("page.sv.md", lfi.virtualName)
	assert.Equal("__hugofs_sv_page.md", lfi.Name())
	assert.Equal("page.md", lfi.RealName())

}
