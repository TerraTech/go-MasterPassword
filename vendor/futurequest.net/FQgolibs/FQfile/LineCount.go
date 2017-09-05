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
	"bufio"
	"os"
)

// LineCount returns the number of lines in given file
//   it is a 'wc -l <file>' workalike
// returns -1 if there is a processing error
func LineCount(name string) int {
	f, err := os.Open(name)
	if err != nil {
		return -1
	}
	var wc int
	// use a scanner as I don't want to slurp in large files
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wc++
	}

	return wc
}
