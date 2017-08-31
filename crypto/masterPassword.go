// Copyright 2017 FutureQuest, Inc.

/*
 *   Attribution: technosophos
 *     Derived from Materminds/sprig: crypto.go/derivePassword()
 */

package crypto

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const master_password_seed = "com.lyndir.masterpassword"

var password_type_templates = map[string][][]byte{
	"maximum": {[]byte("anoxxxxxxxxxxxxxxxxx"), []byte("axxxxxxxxxxxxxxxxxno")},
	"long": {[]byte("CvcvnoCvcvCvcv"), []byte("CvcvCvcvnoCvcv"), []byte("CvcvCvcvCvcvno"), []byte("CvccnoCvcvCvcv"), []byte("CvccCvcvnoCvcv"),
		[]byte("CvccCvcvCvcvno"), []byte("CvcvnoCvccCvcv"), []byte("CvcvCvccnoCvcv"), []byte("CvcvCvccCvcvno"), []byte("CvcvnoCvcvCvcc"),
		[]byte("CvcvCvcvnoCvcc"), []byte("CvcvCvcvCvccno"), []byte("CvccnoCvccCvcv"), []byte("CvccCvccnoCvcv"), []byte("CvccCvccCvcvno"),
		[]byte("CvcvnoCvccCvcc"), []byte("CvcvCvccnoCvcc"), []byte("CvcvCvccCvccno"), []byte("CvccnoCvcvCvcc"), []byte("CvccCvcvnoCvcc"),
		[]byte("CvccCvcvCvccno")},
	"medium": {[]byte("CvcnoCvc"), []byte("CvcCvcno")},
	"short":  {[]byte("Cvcn")},
	"basic":  {[]byte("aaanaaan"), []byte("aannaaan"), []byte("aaannaaa")},
	"pin":    {[]byte("nnnn")},
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
}

// MasterPassword returns a derived password according to: http://masterpasswordapp.com/algorithm.html
func MasterPassword(counter uint32, password_type, user, password, site string) (string, error) {
	templates := password_type_templates[password_type]
	if templates == nil {
		return "", fmt.Errorf("cannot find password template %s", password_type)
	}

	var buffer bytes.Buffer
	buffer.WriteString(master_password_seed)
	binary.Write(&buffer, binary.BigEndian, uint32(len(user)))
	buffer.WriteString(user)

	salt := buffer.Bytes()
	key, err := scrypt.Key([]byte(password), salt, 32768, 8, 2, 64)
	if err != nil {
		return "", fmt.Errorf("failed to derive password: %s", err)
	}

	buffer.Truncate(len(master_password_seed))
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
