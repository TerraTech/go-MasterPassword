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
)

// IsFifo checks given filename to see if it is a Fifo (named pipe).
func IsFifo(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && (fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe)
}
