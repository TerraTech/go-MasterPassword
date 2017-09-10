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

package crypto

// MasterPW contains all relevant items for MasterPassword to act upon.
type MPConfig struct {
	MasterPasswordSeed string `toml:"masterPasswordSeed,omitempty"`
	PasswordType       string `toml:"passwordType,omitempty"`
	Fullname           string `toml:"fullname,omitempty"`
	Password           string `toml:"password,omitempty"`
	Site               string `toml:"site,omitempty"`
	Counter            uint32 `toml:"counter,omitempty"` // Counter >= 1
}

// NewMPConfig returns a new MPConfig with defaults set
func NewMPConfig() *MPConfig {
	return &MPConfig{
		MasterPasswordSeed: MasterPasswordSeed,
		PasswordType:       "long",
		Counter:            1,
	}
}
