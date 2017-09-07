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
	"testing"

	"futurequest.net/FQgolibs/FQtesting"
	"github.com/stretchr/testify/assert"
)

var D = FQtesting.D
var ane = FQtesting.ANE

func TestGcfn(t *testing.T) {
	expected := []string{"gompw.toml", "TESTHOME/.gompw.toml", "/etc/gompw.toml"}

	abort := make(chan struct{})
	defer close(abort)
	ch := Gcfn(DefaultConfigFilename, "TESTHOME", abort)
	var i = 0
	for cf := range ch {
		if !assert.Equal(t, expected[i], cf) {
			return
		}
		i++
	}
}

func TestLoadConfig(t *testing.T) {
	var c Config

	expected := Config{
		Fullname:     "TestUser",
		Password:     "liveLifeToTheEdge",
		PasswordType: "maximum",
		Site:         "FutureQuest.net",
		Counter:      69,
	}

	err := c.LoadConfig("../../files/gompw.toml")
	ane(t, err)
	assert.Equal(t, expected, c)

	// test 'Counter' and 'PasswordType' defaults when 'omitempty'
	expected.Counter = 1
	expected.PasswordType = "long"

	c = NewConfig()
	err = c.LoadConfig("../../files/gompw-omitempty.toml")
	ane(t, err)
	assert.Equal(t, expected, c)

	// test against empty gompw.toml
	expected = Config{
		PasswordType: "long",
		Counter:      1,
	}
	c = NewConfig()
	err = c.LoadConfig("../../files/gompw-empty.toml")
	ane(t, err)
	assert.Equal(t, expected, c)
}

func TestMerge(t *testing.T) {
	var c Config

	expected := Config{
		Fullname:     "TestUser",
		Password:     "liveLifeToTheEdge",
		PasswordType: "long",
		Counter:      1,
	}

	err := c.LoadConfig("../../files/gompw-merge.toml")
	ane(t, err)
	assert.Equal(t, expected, c)
}
