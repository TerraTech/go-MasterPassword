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
	"path"

	"futurequest.net/FQgolibs/FQversion"
	"github.com/TerraTech/go-MasterPassword/pkg/crypto"
)

var (
	PROG string = path.Base(os.Args[0])
	// VERSION follows the Major.Minor of mpw cli, however .Patch is incremented for changes to gompw
	VERSION   string = "2.6.0"
	BUILD     string = FQversion.GetBUILD()
	BUILDHOST string // Filled via Makefile
)

type MPW struct {
	*crypto.MasterPW
	fd     uint
	pwFile string
}

func main() {
	mpw := &MPW{
		MasterPW: crypto.NewMasterPassword(),
	}

	handleFlags(mpw)

	mPassword, err := mpw.MasterPassword()
	if err != nil {
		fatal(err.Error())
	}

	fmt.Println(mPassword)
}
