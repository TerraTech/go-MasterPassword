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

package crypto_test

import (
	"testing"

	"github.com/TerraTech/go-MasterPassword/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestValidateFullname(t *testing.T) {
	// good
	assert.NoError(t, crypto.ValidateFullname("Terra"))

	// bad
	assert.Error(t, crypto.ErrFullnameEmpty, crypto.ValidateFullname(""))
}

// ValidateMasterPasswordSeed validates that seed is not empty
func TestValidateMasterPasswordSeed(t *testing.T) {
	// good
	assert.NoError(t, crypto.ValidateMasterPasswordSeed("liveLifeBeyondAllYourTomorrrows"))

	// bad
	assert.Error(t, crypto.ErrMasterPasswordSeedEmpty, crypto.ValidateMasterPasswordSeed(""))
}

// ValidatePassword verifies that password is not empty
func TestValidatePassword(t *testing.T) {
	// good
	assert.NoError(t, crypto.ValidatePassword("imapassword"))

	// bad
	assert.Error(t, crypto.ErrPasswordEmpty, crypto.ValidatePassword(""))
}

// ValidatePasswordType verifies that passwordType is not empty and a valid type
func TestValidatePasswordType(t *testing.T) {
	// good
	assert.NoError(t, crypto.ValidatePasswordType("maximum"))

	// bad
	assert.Error(t, crypto.ErrPasswordTypeInvalid, crypto.ValidatePasswordType("overdrive"))
	assert.Error(t, crypto.ErrPasswordTypeEmpty, crypto.ValidatePasswordType(""))
}

func TestValidateSite(t *testing.T) {
	// good
	assert.NoError(t, crypto.ValidateSite("futurequest.net"))

	// bad
	assert.Error(t, crypto.ErrSiteEmpty, crypto.ValidateSite(""))
}
