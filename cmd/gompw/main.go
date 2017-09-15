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
	"os"
	"path"

	"futurequest.net/FQgolibs/FQversion"
	"github.com/TerraTech/go-MasterPassword/pkg/config"
	"github.com/TerraTech/go-MasterPassword/pkg/crypto"
)

var (
	// PROG is used for building the version string
	PROG = path.Base(os.Args[0])
	// VERSION is filled by Makefile and used for building the version string
	VERSION string
	// BUILD is filled by Makefile and used for building the version string
	BUILD = FQversion.GetBUILD()
	// BUILDHOST is filled by Makefile and used for building the version string
	BUILDHOST string
)

type mpw struct {
	*crypto.MasterPW
	cu     *config.MPConfig // (MP)Config User, loaded from .toml files
	fd     uint
	pwFile string
	ssp    bool
}

func main() {
	mpw := &mpw{
		MasterPW: crypto.NewMasterPassword(),
		cu:       &config.MPConfig{},
	}

	handleFlags(mpw)

	mPassword, err := mpw.MasterPassword()
	if err != nil {
		fatal(err.Error())
	}

	printPassword(mpw, mPassword)
}
