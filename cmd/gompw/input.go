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

package main

import (
	"fmt"
	"os"

	"futurequest.net/FQgolibs/FQterm"
)

func readInput(prompt string, ssp bool) (string, error) {
	var input string
	var err error

	fmt.Fprint(os.Stderr, prompt)
	if ssp {
		input, err = FQterm.ReadPassword(os.Stdin)
		fmt.Fprintln(os.Stderr)
	} else {
		input, err = FQterm.ReadInput(os.Stdin)
	}

	return input, err
}
