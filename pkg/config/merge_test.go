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

	"github.com/TerraTech/go-MasterPassword/pkg/config"
	"github.com/TerraTech/go-MasterPassword/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

// TestMergeGood tests that the merge occurred correctly.
//
// This needs to be run in the 'config_test' public context to avoid an import cycle
func TestMergeGood(t *testing.T) {
	m := &crypto.MasterPW{
		Config: &config.MPConfig{}, // simulate what toml.Unmarshal will do to MPConfig on missing config items
	}
	c := &config.MPConfig{
		MasterPasswordSeed: "masterpasswordseed",
		PasswordType:       "passwordtype",
		Fullname:           "fullname",
		Password:           "password",
		Site:               "site",
		Counter:            69,
	}

	m.Config.Merge(c)
	assert.Equal(t, m.Config, c)
}
