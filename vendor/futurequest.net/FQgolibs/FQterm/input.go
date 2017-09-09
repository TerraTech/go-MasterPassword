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

package FQterm

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadInput reads terminal input
func ReadInput(f *os.File) (string, error) {
	reader := bufio.NewReader(f)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

// ReadInputFD reads terminal input from given file descriptor
func ReadInputFD(fd uintptr) (string, error) {
	return ReadInput(os.NewFile(fd, ""))
}

// ReadPassword reads terminal input with echoing disabled
func ReadPassword(f *os.File) (string, error) {
	return ReadPasswordFD(f.Fd())
}

// ReadPasswordFD reads terminal input from given file descriptor with echoing disabled
func ReadPasswordFD(fd uintptr) (string, error) {
	fdi := int(fd)

	if IsaTTY(fd) {
		oldState, err := terminal.GetState(fdi)
		if err != nil {
			return "", err
		}
		// In case user presses Ctrl-C during input
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			terminal.Restore(fdi, oldState)
			fmt.Println()
			os.Exit(1)
		}()
	}

	input, err := terminal.ReadPassword(fdi)

	return strings.TrimSpace(string(input)), err
}
