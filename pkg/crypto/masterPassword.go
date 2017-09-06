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
	"sort"

	"github.com/TerraTech/go-MasterPassword/pkg/common"
	"golang.org/x/crypto/scrypt"
)

// MpwSeries denotes the mpw cli client version compatibility.
//   This is mostly for tracking the password type templates.
const MpwSeries = "2.6"

// MasterPasswordTypes is for listing the current supported password types.
const MasterPasswordTypes = "basic, long, maximum, medium, name, phrase, pin, short"

const MasterPasswordSeed = "com.lyndir.masterpassword"

// MasterPW contains all relevant items for MasterPassword to act upon.
type MasterPW struct {
	PasswordType       string `toml:"passwordType,omitempty"`
	Fullname           string `toml:"fullname,omitempty"`
	Password           string `toml:"password,omitempty"`
	Site               string `toml:"site,omitempty"`
	Counter            uint32 `toml:"counter,omitempty"` // Counter >= 1
}

var password_type_templates = map[string][][]byte{
	"basic": {[]byte("aaanaaan"), []byte("aannaaan"), []byte("aaannaaa")},
	"long": {[]byte("CvcvnoCvcvCvcv"), []byte("CvcvCvcvnoCvcv"), []byte("CvcvCvcvCvcvno"), []byte("CvccnoCvcvCvcv"), []byte("CvccCvcvnoCvcv"),
		[]byte("CvccCvcvCvcvno"), []byte("CvcvnoCvccCvcv"), []byte("CvcvCvccnoCvcv"), []byte("CvcvCvccCvcvno"), []byte("CvcvnoCvcvCvcc"),
		[]byte("CvcvCvcvnoCvcc"), []byte("CvcvCvcvCvccno"), []byte("CvccnoCvccCvcv"), []byte("CvccCvccnoCvcv"), []byte("CvccCvccCvcvno"),
		[]byte("CvcvnoCvccCvcc"), []byte("CvcvCvccnoCvcc"), []byte("CvcvCvccCvccno"), []byte("CvccnoCvcvCvcc"), []byte("CvccCvcvnoCvcc"),
		[]byte("CvccCvcvCvccno")},
	"maximum": {[]byte("anoxxxxxxxxxxxxxxxxx"), []byte("axxxxxxxxxxxxxxxxxno")},
	"medium":  {[]byte("CvcnoCvc"), []byte("CvcCvcno")},
	"name":    {[]byte("cvccvcvcv")},
	"phrase":  {[]byte("cvcc cvc cvccvcv cvc"), []byte("cvc cvccvcvcv cvcv"), []byte("cv cvccv cvc cvcvccv")},
	"pin":     {[]byte("nnnn")},
	"short":   {[]byte("Cvcn")},
}

var template_characters = map[byte]string{
	'V': "AEIOU",
	'C': "BCDFGHJKLMNPQRSTVWXYZ",
	'v': "aeiou",
	'c': "bcdfghjklmnpqrstvwxyz",
	'A': "AEIOUBCDFGHJKLMNPQRSTVWXYZ",
	'a': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz",
	'n': "0123456789",
	'o': "@&%?,=[]_:-+*$#!'^~;()/.",
	'x': "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()",
	' ': " ",
}

// NewMasterPassword returns a new empty MasterPW struct with counter==1 and pwtype=="long"
func NewMasterPassword() *MasterPW {
	return &MasterPW{
		Counter:      1,
		PasswordType: "long",
	}
}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
// Valid password_types: maximum, long, medium, short, basic, pin
func (m *MasterPW) MasterPassword() (string, error) {
	return MasterPassword(m.Counter, m.PasswordType, m.Fullname, m.Password, m.Site)
}

// GetPasswordTypes returns a sorted list of valid password types
func (m *MasterPW) GetPasswordTypes() []string {
	keys := make([]string, len(password_type_templates))
	i := 0
	for k, _ := range password_type_templates {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

func (m *MasterPW) IsValidPasswordType(password_type string) bool {
	_, exists := password_type_templates[password_type]
	return exists
}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
func MasterPassword(counter uint32, password_type, user, password, site string) (string, error) {
	templates := password_type_templates[password_type]
	if templates == nil {
		return "", fmt.Errorf("cannot find password template %s", password_type)
	}

	if err := common.ValidateSiteCounter(counter); err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(MasterPasswordSeed)
	binary.Write(&buffer, binary.BigEndian, uint32(len(user)))
	buffer.WriteString(user)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(password), salt, 32768, 8, 2, 64)
	if err != nil {
		return "", fmt.Errorf("failed to derive password: %s", err)
	}

	buffer.Truncate(len(MasterPasswordSeed))
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
