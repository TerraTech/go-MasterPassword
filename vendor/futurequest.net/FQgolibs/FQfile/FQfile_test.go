//==============================================================================
// This file is part of FQgolibs
// Copyright (c) 2017, FutureQuest, Inc.
//   https://www.FutureQuest.net
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
//==============================================================================

package FQfile_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"futurequest.net/FQgolibs/FQfile"

	"github.com/stretchr/testify/assert"
)

const fileLineCount = 999
const testDir = "/tmp/test_dir"
const testFile = "/tmp/test_file.log"
const testFileSrc = "/tmp/test_touch_src.log"
const testFileTgt = "/tmp/test_touch_tgt.log"
const testFileTouch = "/tmp/test_touch.log"
const testSymlink = "/tmp/test_symlink.log"
const testFifo = "/tmp/test_fifo.fifo"

func begin() {
	f, err := os.Create(testFile)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= fileLineCount; i++ {
		// write lines for LineCount()
		f.WriteString(fmt.Sprintf("test:%d\n", i))
	}
	f.Close()
	err = os.Mkdir(testDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Symlink(testFile, testSymlink)
	if err != nil {
		log.Fatal(err)
	}
	f, err = os.Create(testFileSrc)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	time.Sleep(1 * time.Second)
	f, err = os.Create(testFileTgt)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	cmd := exec.Command("/usr/bin/mkfifo", testFifo)
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
func end() {
	for _, f := range []string{testFile, testDir, testSymlink, testFileSrc, testFileTgt, testFileTouch, testFifo} {
		err := os.Remove(f)
		if err != nil {
			log.Fatal(f, err)
		}
	}
}

func TestMain(m *testing.M) {
	begin()
	rc := m.Run()
	end()
	os.Exit(rc)
}

func TestExists(t *testing.T) {
	var ft = FQfile.Exists
	assert.True(t, ft(testFile))
	assert.True(t, ft(testDir))
	assert.True(t, ft(testSymlink))
	assert.False(t, ft("/tmp/foo_non-existent-file_bar"))
}

func TestIsDir(t *testing.T) {
	var ft = FQfile.IsDir
	assert.True(t, ft(testDir))
	assert.False(t, ft(testFile))
	assert.False(t, ft("/tmp/foo_non-existent-file_bar"))
}

func TestIsFile(t *testing.T) {
	var ft = FQfile.IsFile
	assert.True(t, ft(testFile))
	assert.False(t, ft(testDir))
	assert.False(t, ft("/tmp/foo_non-existent-file_bar"))
}

func TestIsSymlink(t *testing.T) {
	var ft = FQfile.IsSymlink
	assert.True(t, ft(testSymlink))
	assert.False(t, ft(testFile))
	assert.False(t, ft(testDir))
	assert.False(t, ft("/tmp/foo_non-existent-file_bar"))
}

func TestLineCount(t *testing.T) {
	var wc int
	wc = FQfile.LineCount(testFile)
	assert.Equal(t, fileLineCount, wc)
	wc = FQfile.LineCount("ZnonExistentFileZ")
	assert.Equal(t, -1, wc)
}

func TestTouch(t *testing.T) {
	assert.False(t, FQfile.Exists(testFileTouch))
	err := FQfile.Touch(testFileTouch)
	assert.NoError(t, err)
	assert.True(t, FQfile.Exists(testFileTouch))
}

func TestTouchByReference(t *testing.T) {
	mtimeSrc, mtimeTgt, err := myStat(testFileSrc, testFileTgt)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	if !assert.NotEqual(t, mtimeSrc, mtimeTgt) {
		t.Fatal(err)
	}

	err = FQfile.TouchByReference(testFileSrc, testFileTgt)
	mtimeSrc, mtimeTgt, err = myStat(testFileSrc, testFileTgt)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	if !assert.Equal(t, mtimeSrc, mtimeTgt) {
		t.Fatal(err)
	}
}

func TestIsFifo(t *testing.T) {
	var ft = FQfile.IsFifo
	assert.True(t, ft(testFifo))
	assert.False(t, ft(testFile))
	assert.False(t, ft(testDir))
	assert.False(t, ft("/tmp/foo_non-existent-file_bar"))
}

func myStat(src, tgt string) (mtimeSrc, mtimeTgt time.Time, err error) {
	fiSrc, err := os.Stat(src)
	if err != nil {
		return
	}

	fiTgt, err := os.Stat(tgt)
	if err != nil {
		return
	}

	mtimeSrc = fiSrc.ModTime()
	mtimeTgt = fiTgt.ModTime()

	return
}
