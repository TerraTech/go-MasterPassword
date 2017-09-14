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

import (
	"github.com/TerraTech/go-MasterPassword/pkg/config"
)

// MergeConfig will transfer and validate data from Config to MasterPW for any nil values.
func (mpw *MasterPW) MergeConfig() error {
	return mpw.MergeConfigEX(mpw.Config)
}

// MergeConfigEX will transfer and validate data from given MPConfig to MasterPW for any nil values.
func (mpw *MasterPW) MergeConfigEX(c *config.MPConfig) error {
	if mpw.masterPasswordSeed == "" {
		mpw.masterPasswordSeed = c.MasterPasswordSeed
	}
	if mpw.passwordPurpose == PasswordPurposeUnSet {
		mpw.SetPasswordPurpose(c.PasswordPurpose)
	}
	if mpw.passwordType == "" {
		mpw.passwordType = c.PasswordType
	}
	if mpw.fullname == "" {
		mpw.fullname = c.Fullname
	}
	if mpw.password == "" {
		mpw.password = c.Password
	}
	if mpw.site == "" {
		mpw.site = c.Site
	}
	if mpw.counter == 0 {
		mpw.counter = c.Counter
	}

	return mpw.Validate()
}
