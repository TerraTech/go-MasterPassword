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
	"io/ioutil"
	"os"
	"path/filepath"

	"futurequest.net/FQgolibs/FQfile"
	"github.com/TerraTech/go-MasterPassword/pkg/crypto"
	"github.com/pelletier/go-toml"
)

const DefaultConfigFilename = "gompw.toml"

type Config crypto.MasterPW

// generate configfile names
// Precedence:
//   1) ./gompw.toml
//   2) $HOME/.gompw.toml
//   3) /etc/gompw.toml
func Gcfn(f, home string, abort <-chan struct{}) <-chan string {
	ch := make(chan string)
	// get file name
	gfn := func(d, dot string) string {
		if d == "" {
			//d = "."
		}
		return filepath.Join(d, dot+f)
	}

	go func() {
		defer close(ch)
		cfns := make([]string, 0, 3)
		cfns = append(cfns, gfn(".", ""))
		if home != "" {
			cfns = append(cfns, gfn(home, "."))
		}
		cfns = append(cfns, gfn("/etc", ""))
		for _, cf := range cfns {
			select {
			case ch <- cf:
			case <-abort:
				return
			}
		}
	}()

	return ch
}

func (c *Config) LoadConfig(configFile string) error {
	if configFile != "" {
		// validate given configFile
		if !FQfile.IsFile(configFile) {
			return fmt.Errorf("given gompw config file does not exist")
		}
	} else {
		// walk through the standard gompw configFile(s)
		abort := make(chan struct{})
		ch := Gcfn(DefaultConfigFilename, os.Getenv("HOME"), abort)
		for cf := range ch {
			if FQfile.IsFile(cf) {
				configFile = cf
				close(abort)
				break
			}
		}
	}

	t, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	if len(t) == 0 {
		fmt.Println("!!! zero sized")
		// just return empty on a zero-sized file
		return nil
	}

	// Needs pelletier/go-toml >= 4a000a21a414d139727f616a8bb97f847b1b310b
	err = toml.Unmarshal(t, c)

	return err
}
