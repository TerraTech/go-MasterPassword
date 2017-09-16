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

package config_test

import (
	"testing"

	"github.com/TerraTech/go-MasterPassword/pkg/common"
	"github.com/TerraTech/go-MasterPassword/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGcfn(t *testing.T) {
	expected := []string{"gompw.toml", "TESTHOME/.gompw.toml", "/etc/gompw.toml"}

	abort := make(chan struct{})
	defer close(abort)
	ch := config.Gcfn(common.DefaultConfigFilename, "TESTHOME", abort)
	var i = 0
	for cf := range ch {
		if !assert.Equal(t, expected[i], cf) {
			return
		}
		i++
	}
}

func TestLoadConfig(t *testing.T) {
	var c *config.MPConfig = &config.MPConfig{}

	expected := &config.MPConfig{
		MasterPasswordSeed: "overrideDefaultMPWseed",
		Fullname:           "TestUser",
		Password:           "liveLifeToTheEdge",
		PasswordType:       "maximum",
		Site:               "FutureQuest.net",
		Counter:            69,
	}

	cf := "../../files/gompw.toml"
	expected.ConfigFile = cf
	err := c.LoadConfig(cf)
	assert.NoError(t, err)
	assert.Equal(t, expected, c)

	// test 'Counter' and 'PasswordType' defaults when 'omitempty'
	expected.MasterPasswordSeed = ""
	expected.Counter = 1
	expected.PasswordType = "long"

	c = config.NewMPConfig()
	cf = "../../files/gompw-omitempty.toml"
	expected.ConfigFile = cf
	err = c.LoadConfig(cf)
	assert.NoError(t, err)
	assert.Equal(t, expected, c)
	expected.ConfigFile = ""

	// test against empty gompw.toml
	expected = &config.MPConfig{
		MasterPasswordSeed: common.DefaultMasterPasswordSeed,
		PasswordType:       "long",
		Counter:            1,
	}
	c = config.NewMPConfig()
	err = c.LoadConfig("../../files/gompw-empty.toml")
	assert.NoError(t, err)
	assert.Equal(t, expected, c)
}

func TestMerge(t *testing.T) {
	var c config.MPConfig

	expected := config.MPConfig{
		Fullname:     "TestUser",
		Password:     "liveLifeToTheEdge",
		PasswordType: "long",
		Counter:      1,
	}

	cf := "../../files/gompw-merge.toml"
	expected.ConfigFile = cf
	err := c.LoadConfig(cf)
	assert.NoError(t, err)
	assert.Equal(t, expected, c)
}
