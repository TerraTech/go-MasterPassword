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

package FQdebug

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

// D is wrapper around spew.Dump with DisableMethods == true
func D(args ...interface{}) {
	s := spew.NewDefaultConfig()
	s.DisableMethods = true
	s.Dump(args...)
}

// DE is D() + os.Exit(1)
func DE(args ...interface{}) {
	D(args)
	os.Exit(1)
}

// DF will dump to spew.dump file
func DF(args ...interface{}) {
	w, err := os.OpenFile("spew.dump", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		return
	}
	defer w.Close()

	spew.Fdump(w, args...)
}

// Dump is wrapper around spew.Dump with default config
func Dump(args ...interface{}) {
	spew.Dump(args...)
}
