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

	"github.com/TerraTech/go-MasterPassword/pkg/common"
)

const DefaultPasswordPurpose = common.DefaultPasswordPurpose

const (
	PasswordPurposeUnSet PasswordPurpose = iota
	PasswordPurposeAuthentication
	PasswordPurposeIdentification
	PasswordPurposeRecovery
)

var (
	ErrPasswordPurposeEmpty   = errors.New("Site password purpose must be set")
	ErrPasswordPurposeInvalid = errors.New("Invalid site password purpose")
)

var (
	ppmap = map[string]PasswordPurpose{
		"a":     PasswordPurposeAuthentication,
		"auth":  PasswordPurposeAuthentication,
		"i":     PasswordPurposeIdentification,
		"ident": PasswordPurposeIdentification,
		"r":     PasswordPurposeRecovery,
		"rec":   PasswordPurposeRecovery,
	}

	pampp = map[PasswordPurpose]string{
		PasswordPurposeAuthentication: "auth",
		PasswordPurposeIdentification: "ident",
		PasswordPurposeRecovery:       "rec",
	}
)

// PasswordPurpose allows for different generated passwords for same Site depending on its intended purpose.
//
// Given the same {fullname, password, site}, specifying a different 'purpose' will perturb the final generated password
// in a consistent manner.
//
//   0) Authentication
//   1) Identification
//   2) Recovery
type PasswordPurpose int

func (p *PasswordPurpose) String() string {
	switch *p {
	case PasswordPurposeAuthentication:
		return "Authentication"
	case PasswordPurposeIdentification:
		return "Identification"
	case PasswordPurposeRecovery:
		return "Recovery"
	}

	return ""
}

// SetPurpose sets the MasterPassword's generated password purpose
func (mpw *MasterPW) SetPasswordPurpose(purpose string) (err error) {
	if err = ValidatePasswordPurpose(purpose); err == nil {
		mpw.passwordPurpose = ppmap[purpose]
	}
	return
}

// ValidatePasswordPurpose will test if MasterPW password purpose is valid
func (mpw *MasterPW) ValidatePasswordPurpose() error {
	return mpw.passwordPurpose.Validate()
}

// Validate will test if password purpose is valid
func (pp *PasswordPurpose) Validate() error {
	if _, ok := pampp[*pp]; !ok {
		return ErrPasswordPurposeInvalid
	}

	return nil
}

// ValidatePasswordPurpose will test if given purpose is valid
func ValidatePasswordPurpose(purpose string) error {
	if purpose == "" {
		return ErrPasswordPurposeEmpty
	}
	if _, ok := ppmap[purpose]; !ok {
		return ErrPasswordPurposeInvalid
	}

	return nil
}

// purpose returns the string used to munge MasterPassword's seed
func (mpw *MasterPW) purpose() string {
	switch mpw.passwordPurpose {
	case PasswordPurposeAuthentication:
		return ""
	case PasswordPurposeIdentification:
		return ".login"
	case PasswordPurposeRecovery:
		return ".answer"
	}

	panic(ErrPasswordPurposeInvalid.Error())
}
