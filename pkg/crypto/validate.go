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
	"errors"
)

var (
	ErrCounter                 = errors.New("site password counter must be >= 1")
	ErrFullnameEmpty           = errors.New("Site fullname must be set")
	ErrMasterPasswordSeedEmpty = errors.New("MasterPassword seed must be set")
	ErrPasswordEmpty           = errors.New("Site password must be set")
	ErrPasswordTypeEmpty       = errors.New("Password type must be set")
	ErrPasswordTypeInvalid     = errors.New("Password type is invalid")
	ErrSiteEmpty               = errors.New("Site name must be set")
)

// Validate ensures that MasterPW is ready for MasterPassword().
//
//   1) masterPasswordSeed
//   2) passwordType
//   3) passwordPurpose
//   4) fullname
//   5) password
//   6) site
//   7) counter
func (mpw *MasterPW) Validate() error {
	if err := ValidateMasterPasswordSeed(mpw.masterPasswordSeed); err != nil {
		return err
	}
	if err := ValidatePasswordType(mpw.passwordType); err != nil {
		return err
	}
	if err := mpw.ValidatePasswordPurpose(); err != nil {
		return err
	}
	if err := ValidateFullname(mpw.fullname); err != nil {
		return err
	}
	if err := ValidatePassword(mpw.password); err != nil {
		return err
	}
	if err := ValidateSite(mpw.site); err != nil {
		return err
	}
	if err := ValidateCounter(mpw.counter); err != nil {
		return err
	}

	// Extra test to catch the following constraints:
	//   0 > auth >= 1
	//   0 > ident <= 1
	//   0 > rec   <= 1
	if mpw.passwordPurpose != PasswordPurposeAuthentication {
		if mpw.counter > 1 {
			return ErrPasswordPurposeCounterOutOfRange
		}
	}

	return nil
}

// ValidateCounter validates that the site counter value is >= 1
func ValidateCounter(counter uint32) error {
	if counter < 1 {
		return ErrCounter
	}

	return nil
}

// ValidateFullname validates that fullname is not empty
func ValidateFullname(fullname string) error {
	if fullname == "" {
		return ErrFullnameEmpty
	}

	return nil
}

// ValidateMasterPasswordSeed validates that seed is not empty
func ValidateMasterPasswordSeed(seed string) error {
	if seed == "" {
		return ErrMasterPasswordSeedEmpty
	}

	return nil
}

// ValidatePassword verifies that password is not empty
func ValidatePassword(password string) error {
	if password == "" {
		return ErrPasswordEmpty
	}

	return nil
}

/*
 * ValidatePasswordPurpose(): passwordPurpose.go
 */

// ValidatePasswordType verifies that passwordType is not empty and a valid type
func ValidatePasswordType(passwordType string) error {
	if passwordType == "" {
		return ErrPasswordTypeEmpty
	}

	_, exists := passwordTypeTemplates[passwordType]
	if !exists {
		return ErrPasswordTypeInvalid
	}

	return nil
}

// ValidateSite validates that site is not empty
func ValidateSite(site string) error {
	if site == "" {
		return ErrSiteEmpty
	}

	return nil
}
