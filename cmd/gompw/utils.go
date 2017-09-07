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
	"log"
	"strings"

	"futurequest.net/FQgolibs/FQdebug"
	"futurequest.net/FQgolibs/FQversion"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	D        = FQdebug.D
	MP_DEBUG bool
)

func DEBUG(msg string) {
	if MP_DEBUG {
		log.Printf("[DEBUG] %s", msg)
	}
}

func fatal(msg string) {
	log.Fatalf("[Fatal] %s", msg)
}

func isaTTY(fd uintptr) bool {
	return terminal.IsTerminal(int(fd))
}

func listPasswordTypes(m *MPW) {
	fmt.Println("=Valid Password Types=")
	fmt.Println(strings.Join(m.GetPasswordTypes(), "\n"))
}

func readPassword(fd uintptr) ([]byte, error) {
	return terminal.ReadPassword(int(fd))
}

func showVersion() {
	if BUILDHOST == "" {
		// If built via Makefile
		fmt.Printf("%s\n", FQversion.ShowVersionsAligned(PROG, VERSION, BUILD))
	} else {
		// via go get and friends
		// FIXME: FQversion >= v5.0.0 will handle BUILDHOST natively
		fmt.Printf("%s.(%s)\n", FQversion.ShowVersionsAligned(PROG, VERSION, BUILD), BUILDHOST)
	}
}
