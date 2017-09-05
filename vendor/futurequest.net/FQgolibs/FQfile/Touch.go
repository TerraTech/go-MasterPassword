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
	"time"
)

// Touch emulates the unix touch program
//   1) if file doesn't exist, create it
//   2) if file does exist, then update modtime to current time
func Touch(name string) error {
	var err error
	var f *os.File
	if !Exists(name) {
		f, err = os.Create(name)
		if err != nil {
			return err
		}
		f.Close()
		return nil
	}
	now := time.Now()
	err = os.Chtimes(name, now, now)
	if err != nil {
		return err
	}

	return nil
}

// TouchByReference creates and/or sets the target file mtime based on the source file
func TouchByReference(source, target string) error {
	err := Touch(target)
	if err != nil {
		return err
	}
	mtime, err := GetMtime(source)
	if err != nil {
		return err
	}
	err = os.Chtimes(target, mtime, mtime)
	if err != nil {
		return err
	}

	return nil
}
