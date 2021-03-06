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

package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"futurequest.net/FQgolibs/FQfile"
	"github.com/TerraTech/go-MasterPassword/pkg/common"
	"github.com/pelletier/go-toml"
)

// Gcfn generates standard locations of configFile filepaths
//
// Precedence:
//   1) ./gompw.toml
//   2) $HOME/.gompw.toml
//   3) /etc/gompw.toml
func Gcfn(f, home string, abort <-chan struct{}) <-chan string {
	ch := make(chan string)
	// get file name
	gfn := func(d, dot string) string {
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

// LoadConfig will load and toml.Unmarshal the given configFile
func (c *MPConfig) LoadConfig(configFile string) error {
	var t []byte
	var err error

	if configFile != "" {
		// validate given configFile
		if !FQfile.IsFile(configFile) {
			return fmt.Errorf("given gompw config file does not exist")
		}
	} else {
		// walk through the standard gompw configFile(s)
		abort := make(chan struct{})
		ch := Gcfn(common.DefaultConfigFilename, os.Getenv("HOME"), abort)
		for cf := range ch {
			if FQfile.IsFile(cf) {
				configFile = cf
				close(abort)
				break
			}
		}
	}

	if configFile != "" {
		t, err = ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
	}

	if len(t) == 0 {
		// just return empty on a zero-sized file with defaults set
		c.Counter = DefaultCounter
		c.PasswordType = DefaultPasswordType
		return nil
	}

	// stuff away c.dump since Unmarshal will clobber it
	doDump := c.dump

	// Needs pelletier/go-toml >= 4a000a21a414d139727f616a8bb97f847b1b310b
	err = toml.Unmarshal(t, c)
	if err != nil {
		return err
	}

	// stuff away the configFile for Dump() usage
	c.ConfigFile = configFile

	// dump trigger set?
	if doDump {
		err = c.Dump()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	// Set the necessary defaults, since the fields will be nil on 'omitempty'
	if c.Counter == 0 {
		c.Counter = DefaultCounter
	}

	if c.PasswordType == "" {
		c.PasswordType = DefaultPasswordType
	}

	return nil
}
