//==============================================================================
// This file is part of go-MasterPassword
// Copyright (c) 2017, TerraTech
// Development funded by FutureQuest, Inc.
//   https://www.FutureQuest.net
//
// go-MasterPassword is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-MasterPassword is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You can find a copy of the GNU General Public License in the
// LICENSE file.  Alternatively, see <http://www.gnu.org/licenses/>.
//==============================================================================

package debug

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureLogOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)

	return buf.String()
}

func Test_wantDebug(t *testing.T) {
	d := NewDebug()
	d.SetFilename("foo/bar.go")

	assert.True(t, d.wantDebug("foo/bar.go"))
	assert.True(t, d.wantDebug("a/b/c/foo/bar.go"))
	assert.False(t, d.wantDebug("bar.go"))
	assert.False(t, d.wantDebug("ar.go"))
}

func TestDbg(t *testing.T) {
	d := NewDebug()
	d.SetFilename("debug/debug_internal_test.go")

	expect := "ZtesTingZ"
	output := captureLogOutput(func() {
		d.Dbg(expect)
	})
	assert.Contains(t, output, expect)
}
