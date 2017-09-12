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
	"os"

	"futurequest.net/FQgolibs/FQdebug"
	"github.com/TerraTech/go-MasterPassword/pkg/common"
	"github.com/TerraTech/go-MasterPassword/pkg/config"
	"golang.org/x/crypto/scrypt"
)

// MpwSeries denotes the mpw cli client version compatibility.
const MpwSeries = "2.6"

// MasterPasswordSeed is the default seed and allows it to be compatible with
// http://masterpasswordapp.com/algorithm.html
const MasterPasswordSeed = common.MasterPasswordSeed

// MasterPW contains all relevant items for MasterPassword to act upon.
type MasterPW struct {
	Config             *config.MPConfig
	masterPasswordSeed string
	passwordType       string
	fullname           string
	password           string
	site               string
	counter            uint32
}

// NewMasterPassword returns a new empty MasterPW struct
func NewMasterPassword() *MasterPW {
	return &MasterPW{
		Config: config.NewMPConfig(),
	}
}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
//
//   Valid PasswordTypes: basic, long, maximum, medium, name, phrase, pin, short
func (mpw *MasterPW) MasterPassword() (string, error) {
	// Fixup MasterPasswordSeed if ""
	if mpw.masterPasswordSeed == "" && mpw.Config.MasterPasswordSeed == "" {
		mpw.Config.MasterPasswordSeed = MasterPasswordSeed
	}

	// merge (and validate) Config ==> MasterPW
	if err := mpw.MergeConfig(); err != nil {
		return "", err
	}

	// DUMP mpw
	if os.Getenv("MP_DUMP") != "" {
		fmt.Fprintf(os.Stderr, "\n== DUMP =======\n")
		FQdebug.D(mpw)
		fmt.Fprintf(os.Stderr, "===============\n\n")
	}

	templates := passwordTypeTemplates[mpw.passwordType]
	if templates == nil {
		return "", fmt.Errorf("cannot find password template %s", mpw.passwordType)
	}

	if err := ValidateCounter(mpw.counter); err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(mpw.masterPasswordSeed)
	binary.Write(&buffer, binary.BigEndian, uint32(len(mpw.fullname)))
	buffer.WriteString(mpw.fullname)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(mpw.password), salt, 32768, 8, 2, 64)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %s", err)
	}

	buffer.Truncate(len(mpw.masterPasswordSeed))
	binary.Write(&buffer, binary.BigEndian, uint32(len(mpw.site)))
	buffer.WriteString(mpw.site)
	binary.Write(&buffer, binary.BigEndian, mpw.counter)

	var hmacv = hmac.New(sha256.New, key)
	hmacv.Write(buffer.Bytes())
	var seed = hmacv.Sum(nil)
	var temp = templates[int(seed[0])%len(templates)]

	buffer.Truncate(0)
	for i, element := range temp {
		passChars := templateCharacters[element]
		passChar := passChars[int(seed[i+1])%len(passChars)]
		buffer.WriteByte(passChar)
	}

	return buffer.String(), nil
}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
//
//   Valid PasswordTypes: basic, long, maximum, medium, name, phrase, pin, short
func MasterPassword(mpwseed, passwordType, fullname, password, site string, counter uint32) (string, error) {
	mpw := &MasterPW{
		Config:             &config.MPConfig{},
		masterPasswordSeed: mpwseed,
		passwordType:       passwordType,
		fullname:           fullname,
		password:           password,
		site:               site,
		counter:            counter,
	}

	return mpw.MasterPassword()
}
