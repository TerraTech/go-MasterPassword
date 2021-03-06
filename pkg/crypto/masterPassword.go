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
	"github.com/TerraTech/go-MasterPassword/pkg/debug"
	"golang.org/x/crypto/scrypt"
)

// MpwSeries denotes the mpw cli client version compatibility.
const MpwSeries = "2.6"

// DefaultMasterPasswordSeed is the default seed and allows it to be compatible with
// http://masterpasswordapp.com/algorithm.html
const DefaultMasterPasswordSeed = common.DefaultMasterPasswordSeed

// Debugging helpers
var (
	Dbg  = debug.NewDebug().Dbg
	DbgO = debug.NewDebug().DbgO
)

// MasterPW contains all relevant items for MasterPassword to act upon.
type MasterPW struct {
	Config             *config.MPConfig
	masterPasswordSeed string
	passwordPurpose    PasswordPurpose
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
	var err error

	// Fixup MasterPasswordSeed if ""
	if mpw.masterPasswordSeed == "" && mpw.Config.MasterPasswordSeed == "" {
		mpw.Config.MasterPasswordSeed = DefaultMasterPasswordSeed
	}

	// merge (and validate) Config ==> MasterPW
	if err = mpw.MergeConfig(); err != nil {
		return "", err
	}

	// munge the master password seed depending on password purpose
	mpseed := mpw.masterPasswordSeed + mpw.purpose()

	// DUMP mpw
	if os.Getenv("MP_DUMP") != "" {
		fmt.Fprintf(os.Stderr, "\n== DUMP =======\n")
		FQdebug.D(mpw)
		fmt.Fprintf(os.Stderr, "===============\n\n")
	}
	// FIXME: convert to template
	//   Pro: cleans up the code and removes the Dbg() interstitials
	//   Con: if something panics, might not have reached the template call
	Dbg("-- mpw_masterKey (algorithm: 3)")
	Dbg("fullName: %s", mpw.fullname)
	Dbg("password: %s", mpw.password)
	Dbg("masterPassword.id: %s", mpwIDBuf([]byte(mpw.password)))
	Dbg("keyScope: %s", mpw.masterPasswordSeed)
	Dbg("masterKeySalt: keyScope=%s | #fullName=%08X | fullName=%s", mpw.masterPasswordSeed, len(mpw.fullname), mpw.fullname)

	templates := passwordTypeTemplates[mpw.passwordType]

	var buffer bytes.Buffer
	buffer.WriteString(mpw.masterPasswordSeed)
	if err = binary.Write(&buffer, binary.BigEndian, uint32(len(mpw.fullname))); err != nil {
		return "", err
	}
	buffer.WriteString(mpw.fullname)

	salt := buffer.Bytes()
	Dbg("  => masterKeySalt.id: %s", mpwIDBuf(salt))

	key, err := scrypt.Key([]byte(mpw.password), salt, 32768, 8, 2, 64)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %s", err)
	}
	Dbg("masterKey: scrypt( masterPassword, masterKeySalt, N=32768, r=8, p=2, keyLen=64")
	Dbg("  => masterKey.id: %s", mpwIDBuf(key))

	Dbg("-- mpw_siteKey (algorithm: 3)")
	Dbg("siteName: %s", mpw.site)
	Dbg("siteCounter: %d", mpw.counter)
	// FIXME: stringer doesn't appear to be working right
	Dbg("keyPurpose: %d (%s)", mpw.passwordPurpose, mpw.passwordPurpose.String())
	Dbg("keyContext: (null)") // not implemented
	Dbg("keyScope: %s", mpseed)
	Dbg("siteSalt: keyScope=%s | #siteName=%08X | siteName=%s | siteCounter=%08d | #keyContext=(null) | keyContext=(null)",
		mpseed, len(mpw.site), mpw.site, mpw.counter)

	// Danger Will Robinson, passwordPurpose comes into effect here, so caution with the Truncate()
	buffer.Truncate(len(mpw.masterPasswordSeed))
	buffer.WriteString(mpw.purpose()) // add the passwordPurpose suffix
	if err = binary.Write(&buffer, binary.BigEndian, uint32(len(mpw.site))); err != nil {
		return "", err
	}
	buffer.WriteString(mpw.site)
	if err = binary.Write(&buffer, binary.BigEndian, mpw.counter); err != nil {
		return "", err
	}
	Dbg("  => siteSalt.id: %s", mpwIDBuf(buffer.Bytes()))

	Dbg("siteKey: hmac-sha256( masterKey.id=%s, siteSalt )", mpwIDBuf(key))
	var hmacv = hmac.New(sha256.New, key)
	if _, err = hmacv.Write(buffer.Bytes()); err != nil {
		return "", err
	}
	var seed = hmacv.Sum(nil)
	Dbg("  => siteKey.id: %s", mpwIDBuf(seed))

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
func MasterPassword(mpwseed, passwordType, passwordPurpose, fullname, password, site string, counter uint32) (string, error) {
	mpw := &MasterPW{
		Config:             &config.MPConfig{},
		masterPasswordSeed: mpwseed,
		passwordType:       passwordType,
		fullname:           fullname,
		password:           password,
		site:               site,
		counter:            counter,
	}
	// needs to be set via method for validation
	if err := mpw.SetPasswordPurpose(passwordPurpose); err != nil {
		return "", err
	}

	return mpw.MasterPassword()
}
