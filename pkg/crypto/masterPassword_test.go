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

var mpwseeds = []string{
	crypto.MasterPasswordSeed,
	"overrideDefaultMPWseed",
	"liveLifeBeyondAllYourTomorrrows",
	"danceLikeNoOneIsLooking",
}

type mpw struct {
	*crypto.MasterPW
}

type testVector struct {
	ms           string
	c            uint32
	pt, u, pw, s string
	expect       string
}

func newMpw(tv testVector) *mpw {
	return &mpw{
		&crypto.MasterPW{
			MasterPasswordSeed: tv.ms,
			Counter:            tv.c,
			PasswordType:       tv.pt,
			Fullname:           tv.u,
			Password:           tv.pw,
			Site:               tv.s,
		},
	}
}

func TestMasterPassword(t *testing.T) {
	expectations := []testVector{
		{mpwseeds[0], 1, "long", "user", "password", "example.com", "ZedaFaxcZaso9*"},
		{mpwseeds[0], 2, "long", "user", "password", "example.com", "Fovi2@JifpTupx"},
		{mpwseeds[0], 1, "maximum", "user", "password", "example.com", "pf4zS1LjCg&LjhsZ7T2~"},
		{mpwseeds[0], 1, "medium", "user", "password", "example.com", "ZedJuz8$"},
		{mpwseeds[0], 1, "basic", "user", "password", "example.com", "pIS54PLs"},
		{mpwseeds[0], 1, "short", "user", "password", "example.com", "Zed5"},
		{mpwseeds[0], 1, "pin", "user", "password", "example.com", "6685"},
		{mpwseeds[0], 1, "name", "user", "password", "example.com", "zedjuzoco"},
		{mpwseeds[0], 1, "phrase", "user", "password", "example.com", "ze juzxo sax taxocre"},
	}

	expectations_bad := []testVector{
		{mpwseeds[0], 1, "invalidType", "user", "password", "example.com", "1111"},
	}

	for _, tv := range expectations {
		pw, err := crypto.MasterPassword(tv.ms, tv.pt, tv.u, tv.pw, tv.s, tv.c)
		assert.NoError(t, err)
		assert.Equal(t, tv.expect, pw)
	}

	for _, tv := range expectations_bad {
		_, err := crypto.MasterPassword(tv.ms, tv.pt, tv.u, tv.pw, tv.s, tv.c)
		assert.Error(t, err)
	}

	// Test method call
	for _, tv := range expectations {
		mpw := newMpw(tv)
		pw, err := mpw.MasterPassword()
		assert.NoError(t, err)
		assert.Equal(t, tv.expect, pw)
	}
}
