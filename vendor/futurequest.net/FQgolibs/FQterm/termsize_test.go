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

package FQterm_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"futurequest.net/FQgolibs/FQterm"

	"github.com/stretchr/testify/assert"
)

// OK - this is whack!
// FAILS: go test ./...
// WORKS: go test
// WORKS: go test -v ./...
// WORKS: go test -v
func sttySize() (cols int, rows int, err error) {
	cmd := exec.Command("/bin/stty", "size")
	// NOTE: This works when running go test in same directory, but fails
	//  when run via go test ./...  (e.g. make testv)
	cmd.Stdin = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	_, err = fmt.Sscanf(string(out), "%d %d\n", &cols, &rows)
	if err != nil {
		return 0, 0, err
	}

	return
}

func TestTermSize(t *testing.T) {
	sttyCols, sttyRows, err := sttySize()
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	//go test, stdin == /dev/null  :(
	FQterm.TStestFD = os.Stderr.Fd()
	cols, rows := FQterm.TermSize()
	for _, i := range []int{sttyCols, sttyRows, cols, rows} {
		assert.NotZero(t, i)
	}
	assert.Equal(t, sttyCols, cols)
	assert.Equal(t, sttyRows, rows)
}
