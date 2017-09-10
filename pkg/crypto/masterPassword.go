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
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// MpwSeries denotes the mpw cli client version compatibility.
const MpwSeries = "2.6"

// MasterPasswordSeed is the default seed and allows it to be compatible with
// http://masterpasswordapp.com/algorithm.html
const MasterPasswordSeed = "com.lyndir.masterpassword"

// MasterPW contains all relevant items for MasterPassword to act upon.
type MasterPW struct {
	*MPConfig
	masterPasswordSeed string
	passwordType       string
	fullname           string
	password           string
	site               string
	counter            uint32
}

// NewMasterPassword returns a new empty MasterPW struct with counter==1 and pwtype=="long"
func NewMasterPassword() *MasterPW {
	return &MasterPW{
		MPConfig:           NewMPConfig(),
		masterPasswordSeed: MasterPasswordSeed,
		counter:            1,
		passwordType:       "long",
	}
}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
//
//   Valid PasswordTypes: basic, long, maximum, medium, name, phrase, pin, short
func (m *MasterPW) MasterPassword() (string, error) {
	return MasterPassword(m.MasterPasswordSeed, m.PasswordType, m.Fullname, m.Password, m.Site, m.Counter)
}

	}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
//
//   Valid PasswordTypes: basic, long, maximum, medium, name, phrase, pin, short
//
//   NOTE: mpwseed == "", will use the default Master Password Seed, do not change unless you have specific requirements
func MasterPassword(mpwseed, passwordType, user, password, site string, counter uint32) (string, error) {
	if mpwseed == "" {
		mpwseed = MasterPasswordSeed
	}

	templates := passwordTypeTemplates[passwordType]
	if templates == nil {
		return "", fmt.Errorf("cannot find password template %s", passwordType)
	}

	if err := common.ValidateSiteCounter(counter); err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(mpwseed)
	binary.Write(&buffer, binary.BigEndian, uint32(len(user)))
	buffer.WriteString(user)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(password), salt, 32768, 8, 2, 64)
	if err != nil {
		return "", fmt.Errorf("failed to derive password: %s", err)
	}

	buffer.Truncate(len(mpwseed))
	binary.Write(&buffer, binary.BigEndian, uint32(len(site)))
	buffer.WriteString(site)
	binary.Write(&buffer, binary.BigEndian, counter)

	var hmacv = hmac.New(sha256.New, key)
	hmacv.Write(buffer.Bytes())
	var seed = hmacv.Sum(nil)
	var temp = templates[int(seed[0])%len(templates)]

	buffer.Truncate(0)
	for i, element := range temp {
		pass_chars := template_characters[element]
		pass_char := pass_chars[int(seed[i+1])%len(pass_chars)]
		buffer.WriteByte(pass_char)
	}

	return buffer.String(), nil
}
