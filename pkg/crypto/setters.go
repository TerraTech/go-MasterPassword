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

// SetCounter is a setter for MasterPW.counter
func (mpw *MasterPW) SetCounter(counter uint32) (err error) {
	if err = ValidateCounter(counter); err == nil {
		mpw.counter = counter
	}
	return
}

// SetFullname is a setter for MasterPW.fullname
func (mpw *MasterPW) SetFullname(fullname string) (err error) {
	if err = ValidateFullname(fullname); err == nil {
		mpw.fullname = fullname
	}
	return
}

// SetMasterPasswordSeed is a setter for MasterPW.masterPasswordSeed
func (mpw *MasterPW) SetMasterPasswordSeed(seed string) (err error) {
	if err = ValidateMasterPasswordSeed(seed); err == nil {
		mpw.masterPasswordSeed = seed
	}
	return
}

// SetPassword is a setter for MasterPW.password
func (mpw *MasterPW) SetPassword(password string) (err error) {
	if err = ValidatePassword(password); err == nil {
		mpw.password = password
	}
	return
}

/*
 * SetPasswordPurpose(): passwordPurpose.go
 */

// SetPasswordType is a setter for MasterPW.passwordType
func (mpw *MasterPW) SetPasswordType(pwtype string) (err error) {
	if err = ValidatePasswordType(pwtype); err == nil {
		mpw.passwordType = pwtype
	}
	return
}

// SetSite is a setter for MasterPW.site
func (mpw *MasterPW) SetSite(site string) (err error) {
	if err = ValidateSite(site); err == nil {
		mpw.site = site
	}
	return
}
