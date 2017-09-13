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
	"strings"
)

const passwordTypeHelpIndent = 28

var passwordTypeHelp = `Specify the password's type template
Defaults to 'long'
    x, maximum  | 20 characters, contains symbols
    l, long     | Copy-friendly, 14 characters, symbols
    m, medium   | Copy-friendly, 8 characters, symbols
    b, basic    | 8 characters, no symbols
    s, short    | Copy-friendly, 4 characters, no symbols
    i, pin      | 4 numbers
    n, name     | 9 letter name
    p, phrase   | 20 character sentence`

func flagHelp(opt string) string {
	var help string
	switch opt {
	case "t":
		indention := strings.Repeat(" ", passwordTypeHelpIndent)
		help = strings.Replace(passwordTypeHelp, "\n", "\n"+indention, -1)
	}

	return help
}

func printPassword(mpw *MPW, pw string) {
	if !mpw.ssp && isaTTY(os.Stdout.Fd()) {
		fmt.Printf("%s's password for %s:\n", mpw.Config.Fullname, mpw.Config.Site)
	}
	fmt.Println(pw)
}
