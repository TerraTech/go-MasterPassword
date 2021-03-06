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
	"github.com/TerraTech/go-MasterPassword/pkg/crypto"

	flag "github.com/spf13/pflag"
)

func (mpw *mpw) getResponse(prompt, errMsg string) string {
	input, err := readInput(prompt, mpw.ssp)
	if err != nil {
		fatal(err.Error())
	}
	if input == "" {
		fatal(errMsg)
	}

	return input
}

func (mpw *mpw) handleFullname() {
	if mpw.Config.Fullname == "" {
		mpw.Config.Fullname = mpw.getResponse("Your full name: ", "Fullname must be specified")
	}
}

func (mpw *mpw) handlePassword() {
	var err error
	var pwBytes []byte
	var pwInput io.Reader

	// read password from io.Reader
	// Priority:
	// 1) -f
	// 2) -d
	// 3) stdin
	var errNoPassword = "Password must be specified"
	if flag.ShorthandLookup("f").Changed || flag.ShorthandLookup("d").Changed {
		if flag.ShorthandLookup("f").Changed {
			debug("pwInput: file")
			pwInput, err = os.Open(mpw.pwFile)
		} else if flag.ShorthandLookup("d").Changed {
			debug("pwInput: fd")
			pwInput = os.NewFile(uintptr(mpw.fd), "")
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

		mpw.Config.Password = string(bytes.TrimSpace(pwBytes))

		if mpw.Config.Password == "" {
			fatal(errNoPassword)
		}
	} else {
		debug("pwInput: stdin")
		mpw.Config.Password = mpw.getResponse("Your master password: ", errNoPassword)
	}
}

func (mpw *mpw) handleSite() {
	// handle site
	site := flagDefaults("", flag.Arg(0), os.Getenv("MP_SITE"))
	if site == "" {
		site = mpw.getResponse("Site name: ", "Site must be specified")
	}
	mpw.Config.Site = site
}

func (mpw *mpw) handleUserConfigLoading(configFile string, dump bool) {
	if dump {
		mpw.cu.SetDump(dump)
	}

	err := mpw.cu.LoadConfig(configFile)
	if err != nil {
		fatal(err.Error())
	}

	// prime MasterPW struct with user configFile settings
	mpw.Config.Merge(mpw.cu)
}

func handleFlags(mpw *mpw) {
	var configFile string
	var err error
	var flagDumpConfig bool
	var flagListPasswordTypes bool
	var flagShowVersion bool
	var ignoreConfigFile bool

	/* Flow of config for standard usage
	 *  (NOTE: MasterPW.(priv-members) has full range of setters for advanced usage
	 *
	 *   MasterPW.(priv-members) <= MasterPW.Config (flag set) <=merge== MPConfig (userConfig file)
	 */

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] site\n", PROG)
		flag.PrintDefaults()
		fmt.Println("\n==Environment Variables==")
		fmt.Println("  MP_CONFIGFILE   | The user configuration file (see -C)")
		//             MP_DEBUG
		//             MP_DUMP
		fmt.Println("  MP_FULLNAME     | The full name of the user (see -u)")
		fmt.Println("  MP_PWPURPOSE    | The password purpose (see -p)")
		fmt.Println("  MP_PWTYPE       | The password type (see -t)")
		fmt.Println("  MP_SEED         | The master password seed (see -S)")
		fmt.Println("  MP_SITE         | The site for generated password")
		fmt.Println("  MP_SITECOUNTER  | The default site counter value (see -c)")

		fmt.Println("\n==User Config file location search order==")
		fmt.Println("  1) ./gompw.toml")
		fmt.Println("  2) $HOME/.gompw.toml")
		fmt.Println("  3) /etc/gompw.toml")
	}

	// "-v" reserved for '--verbose' if implemented
	flag.BoolVarP(&flagDumpConfig, "dumpConfig", "D", false, "Dump the user configuration file and exit")
	flag.BoolVarP(&flagListPasswordTypes, "listPasswordTypes", "l", false, "List valid Password Types")
	flag.BoolVarP(&flagShowVersion, "version", "V", false, "Show version")
	flag.BoolVarP(&ignoreConfigFile, "ignoreUserConfig", "I", false, "Ignore user configuration file")
	flag.BoolVar(&mpw.ssp, "ssp", false, "Shoulder Surfing Prevention by not echoing any terminal input")
	flag.StringVarP(&configFile, "config", "C", "", "User configuration file override")
	flag.StringVarP(&mpw.Config.Fullname, "fullname", "u", os.Getenv("MP_FULLNAME"), "Fullname")
	flag.StringVarP(&mpw.Config.MasterPasswordSeed, "mpseed", "S", flagDefaults(common.DefaultMasterPasswordSeed, os.Getenv("MP_SEED")), "Override the Master Password Seed")
	flag.StringVarP(&mpw.Config.PasswordPurpose, "purpose", "p", flagDefaults(common.DefaultPasswordPurpose, os.Getenv("MP_PWPURPOSE")), flagHelp("p"))
	flag.StringVarP(&mpw.Config.PasswordType, "pwtype", "t", flagDefaults(common.DefaultPasswordType, os.Getenv("MP_PWTYPE")), flagHelp("t"))
	flag.StringVarP(&mpw.pwFile, "file", "f", "", "Read user's master password from given filename")
	flag.Uint32VarP(&mpw.Config.Counter, "counter", "c", flagDefaultCounter(os.Getenv("MP_SITECOUNTER")), "Site password counter value")
	flag.UintVarP(&mpw.fd, "fd", "d", 0, "Read user's master password from given file descriptor")

	flag.Parse()

	if flagShowVersion {
		showVersion()
		os.Exit(0)
	}

	if flagListPasswordTypes {
		listPasswordTypes(mpw)
		os.Exit(0)
	}

	// -d and -f are mutually exclusive
	if flag.ShorthandLookup("d").Changed && flag.ShorthandLookup("f").Changed {
		fatal("-d and -f are mutually exclusive.")
	}

	// -I and -C are mutually exclusive
	if flag.ShorthandLookup("I").Changed && flag.ShorthandLookup("C").Changed {
		fatal("-I and -C are mutually exclusive.")
	}

	// -c ( >1 ) does not work with -p [i,r] (or -p [!a])
	if mpw.Config.Counter > 1 && flag.ShorthandLookup("c").Changed && flag.ShorthandLookup("p").Changed {
		if err = crypto.ValidatePasswordPurpose(mpw.Config.PasswordPurpose); err != nil {
			fatal(err.Error())
		}

		var token crypto.PasswordPurpose
		token, err = crypto.PasswordPurposeToToken(mpw.Config.PasswordPurpose)
		if err != nil {
			fatal(err.Error())
		}
		if token != crypto.PasswordPurposeAuthentication {
			fatal(crypto.ErrPasswordPurposeCounterOutOfRange.Error())
		}
	}

	// prime the pump
	if !ignoreConfigFile {
		mpw.handleUserConfigLoading(configFile, flagDumpConfig)
	}

	mpw.handleFullname()
	mpw.handlePassword()
	mpw.handleSite()
}

func flagDefaults(_default string, overrides ...string) string {
	for _, override := range overrides {
		if override != "" {
			return override
		}
	}

	return _default
}

func flagDefaultCounter(override string) uint32 {
	if override != "" {
		mpsc, err := strconv.Atoi(override)
		if err != nil {
			log.Print("Invalid value specified for MP_SITECOUNTER")
			log.Fatal(err.Error())
		}
		return uint32(mpsc)
	}

	return common.DefaultCounter
}
