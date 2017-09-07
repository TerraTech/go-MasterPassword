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

package FQfile

import (
	"os"

	//crickets: "github.com/phayes/permbits"
	"github.com/goruha/permbits"
)

// https://godoc.org/github.com/phayes/permbits
func Chmod(filepath string, b permbits.PermissionBits) error {
	return permbits.Chmod(filepath, b)
}

// https://godoc.org/github.com/phayes/permbits
func FileMode(fm os.FileMode) permbits.PermissionBits {
	return permbits.FileMode(fm)
}

// https://godoc.org/github.com/phayes/permbits
func Stat(filepath string) (permbits.PermissionBits, error) {
	return permbits.Stat(filepath)
}

// https://godoc.org/github.com/phayes/permbits
func UpdateFileMode(fm *os.FileMode, b permbits.PermissionBits) {
	permbits.UpdateFileMode(fm, b)
}
