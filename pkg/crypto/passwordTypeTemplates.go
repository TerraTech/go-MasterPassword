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
	"sort"
)

func init() {
	var ptt = passwordTypeTemplates

	// add shortcodes
	ptt["b"] = ptt["basic"]
	ptt["l"] = ptt["long"]
	ptt["x"] = ptt["maximum"]
	ptt["m"] = ptt["medium"]
	ptt["n"] = ptt["name"]
	ptt["p"] = ptt["phrase"]
	ptt["i"] = ptt["pin"]
	ptt["s"] = ptt["short"]
}

// MasterPasswordTypes is for listing the current supported password types.
//
//   Default: long
const MasterPasswordTypes = "basic, long, maximum, medium, name, phrase, pin, short"

var (
	passwordTypeTemplates = map[string][][]byte{
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

	templateCharacters = map[byte]string{
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
)

// GetPasswordTypes returns a sorted list of valid password types
func (m *MasterPW) GetPasswordTypes() []string {
	keys := make([]string, len(passwordTypeTemplates))
	i := 0
	for k := range passwordTypeTemplates {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}
