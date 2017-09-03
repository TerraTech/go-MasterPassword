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
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/TerraTech/go-MasterPassword/crypto"
	flag "github.com/spf13/pflag"
)

func handleFlags(m *MPW) {
	var err error
	var pwInput io.Reader

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] site\n", PROG)
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("==Environment Variables==")
		fmt.Println("  MP_FULLNAME")
		fmt.Println("  MP_PWTYPE")
	}

	default_pwType := m.PWtype // stuff away default PWtype
	var flagShowVersion bool
	var flagListPWtypes bool

	flagthelp := func(msg string) string {
		return fmt.Sprintf("%s [%s]", msg, crypto.Master_password_types)
	}

	// "-v" reserved for '--verbose' if implemented
	flag.UintVarP(&m.fd, "fd", "d", 0, "Read user's master password from given file descriptor.")
	flag.StringVarP(&m.pwFile, "file", "f", "", "Read user's master password from given file.")
	flag.StringVarP(&m.Fullname, "fullname", "u", os.Getenv("MP_FULLNAME"), "Fullname")
	flag.StringVarP(&m.PWtype, "pwtype", "t", os.Getenv("MP_PWTYPE"), flagthelp("Password Type"))
	flag.BoolVarP(&flagListPWtypes, "listPWtypes", "l", false, "List valid Password Types")
	flag.BoolVarP(&flagShowVersion, "version", "V", false, "Show version")

	flag.Parse()

	if flagShowVersion {
		showVersion()
		os.Exit(0)
	}

	if flagListPWtypes {
		listPWtypes(m)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// -d and -f are mutually exclusive
	if flag.ShorthandLookup("d").Changed && flag.ShorthandLookup("f").Changed {
		fatal("-d and -f are mutually exclusive.")
	}

	if m.Fullname == "" {
		fatal("Fullname must be specified")
	}

	m.Site = flag.Arg(0)
	if m.Site == "" {
		fatal("Site must be specified")
	}

	// read password from io.Reader
	// Priority:
	// 1) -f
	// 2) -d
	// 3) stdin
	if flag.ShorthandLookup("f").Changed {
		pwInput, err = os.Open(m.pwFile)
		if err != nil {
			fatal(err.Error())
		}
	} else if flag.ShorthandLookup("d").Changed {
		pwInput = os.NewFile(uintptr(m.fd), "")
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			pwInput = os.Stdin
		}
	}

	if pwInput == nil {
		fatal("Cannot create io.Reader for password input")
	}

	pwBytes, err := ioutil.ReadAll(pwInput)
	if err != nil {
		fatal(err.Error())
	}
	m.Password = string(bytes.TrimSpace(pwBytes))

	if m.Password == "" {
		fatal("Password must be specified")
	}

	// handle pwType
	if m.PWtype == "" {
		m.PWtype = default_pwType
	}
}
