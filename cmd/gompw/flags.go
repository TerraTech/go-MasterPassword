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
	"log"
	"os"
	"strconv"

	"github.com/TerraTech/go-MasterPassword/pkg/common"
	flag "github.com/spf13/pflag"
)

func handleFlags(m *MPW) {
	var config Config
	var configFile string
	var err error
	var ignoreConfigFile bool
	var pwBytes []byte
	var pwInput io.Reader

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] site\n", PROG)
		flag.PrintDefaults()
		fmt.Println("\n==Environment Variables==")
		fmt.Println("  MP_CONFIGFILE   | The user configuration file (see -C)")
		//             MP_DEBUG
		fmt.Println("  MP_FULLNAME     | The full name of the user (see -u)")
		fmt.Println("  MP_PWTYPE       | The password type (see -t)")
		fmt.Println("  MP_SEED         | The master password seed (see -S)")
		fmt.Println("  MP_SITECOUNTER  | The default counter value (see -c)")

		fmt.Println("\n==User Config file location search order==")
		fmt.Println("  1) ./gompw.toml")
		fmt.Println("  2) $HOME/.gompw.toml")
		fmt.Println("  3) /etc/gompw.toml")
	}

	default_pwType := m.PasswordType // stuff away default PasswordType
	var flagShowVersion bool
	var flagListPasswordTypes bool

	// "-v" reserved for '--verbose' if implemented
	flag.BoolVarP(&flagListPasswordTypes, "listPasswordTypes", "l", false, "List valid Password Types")
	flag.BoolVarP(&flagShowVersion, "version", "V", false, "Show version")
	flag.BoolVarP(&ignoreConfigFile, "ignoreUserConfig", "I", false, "Ignore user configuration file")
	flag.BoolVar(&m.ssp, "ssp", false, "Shoulder Surfing Prevention by not echoing any terminal input")
	flag.StringVarP(&configFile, "config", "C", "", "User configuration file override")
	flag.StringVarP(&m.Fullname, "fullname", "u", os.Getenv("MP_FULLNAME"), "Fullname")
	flag.StringVarP(&m.MasterPasswordSeed, "mpseed", "S", os.Getenv("MP_SEED"), "Override the Master Password Seed")
	flag.StringVarP(&m.PasswordType, "pwtype", "t", os.Getenv("MP_PWTYPE"), flagHelp("t"))
	flag.StringVarP(&m.pwFile, "file", "f", "", "Read user's master password from given filename")
	flag.Uint32VarP(&m.Counter, "counter", "c", 1, "Site password counter value")
	flag.UintVarP(&m.fd, "fd", "d", 0, "Read user's master password from given file descriptor")

	flag.Parse()

	if os.Getenv("MP_DEBUG") != "" {
		MP_DEBUG = true
	}

	if flagShowVersion {
		showVersion()
		os.Exit(0)
	}

	if flagListPasswordTypes {
		listPasswordTypes(m)
		os.Exit(0)
	}

	// -d and -f are mutually exclusive
	if flag.ShorthandLookup("d").Changed && flag.ShorthandLookup("f").Changed {
		fatal("-d and -f are mutually exclusive.")
	}

	if !ignoreConfigFile {
		err := config.LoadConfig(configFile)
		if err != nil {
			fatal(err.Error())
		}

		// prime MasterPW struct with user configFile settings
		config.Merge(m)
	}

	getResponse := func(prompt, errMsg string) string {
		input, err := readInput(prompt, m.ssp)
		if err != nil {
			fatal(err.Error())
		}
		if input == "" {
			fatal(errMsg)
		}

		return input
	}

	if m.Fullname == "" {
		m.Fullname = getResponse("Your full name: ", "Fullname must be specified")
	}

	// read password from io.Reader
	// Priority:
	// 1) -f
	// 2) -d
	// 3) stdin
	var errNoPassword = "Password must be specified"
	if flag.ShorthandLookup("f").Changed || flag.ShorthandLookup("d").Changed {
		if flag.ShorthandLookup("f").Changed {
			DEBUG("pwInput: file")
			pwInput, err = os.Open(m.pwFile)
		} else if flag.ShorthandLookup("d").Changed {
			DEBUG("pwInput: fd")
			pwInput = os.NewFile(uintptr(m.fd), "")
		}
		if err != nil {
			fatal(err.Error())
		}

		if pwInput == nil {
			fatal("Cannot create io.Reader for password input")
		}

		pwBytes, err = ioutil.ReadAll(pwInput)
		if err != nil {
			fatal(err.Error())
		}

		m.Password = string(bytes.TrimSpace(pwBytes))

		if m.Password == "" {
			fatal(errNoPassword)
		}
	} else {
		DEBUG("pwInput: stdin")
		m.Password = getResponse("Your master password: ", errNoPassword)
	}

	m.Site = flag.Arg(0)
	if m.Site == "" {
		m.Site = getResponse("Site name: ", "Site must be specified")
	}

	// handle pwType
	if m.PasswordType == "" {
		m.PasswordType = default_pwType
	}

	// handle site counter
	if !flag.ShorthandLookup("c").Changed {
		// Check to see if defined via envariable
		siteCounter := os.Getenv("MP_SITECOUNTER")
		if siteCounter != "" {
			mc, err := strconv.Atoi(siteCounter)
			if err != nil {
				log.Print("Invalid value specified for MP_SITECOUNTER")
				log.Fatal(err.Error())
			}
			m.Counter = uint32(mc)
		}

	}
	if err := common.ValidateSiteCounter(m.Counter); err != nil {
		log.Fatal(err.Error())
	}
}
