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

	"github.com/TerraTech/go-MasterPassword/crypto"
	"github.com/stretchr/testify/assert"
)

func TestDerivePassword(t *testing.T) {
	type testVector struct {
		c            uint32
		pt, u, pw, s string
		expect       string
	}
	expectations := []testVector{
		{1, "long", "user", "password", "example.com", "ZedaFaxcZaso9*"},
		{2, "long", "user", "password", "example.com", "Fovi2@JifpTupx"},
		{1, "maximum", "user", "password", "example.com", "pf4zS1LjCg&LjhsZ7T2~"},
		{1, "medium", "user", "password", "example.com", "ZedJuz8$"},
		{1, "basic", "user", "password", "example.com", "pIS54PLs"},
		{1, "short", "user", "password", "example.com", "Zed5"},
		{1, "pin", "user", "password", "example.com", "6685"},
		{1, "name", "user", "password", "example.com", "zedjuzoco"},
		{1, "phrase", "user", "password", "example.com", "ze juzxo sax taxocre"},
	}

	expectations_bad := []testVector{
		{1, "invalidType", "user", "password", "example.com", "1111"},
	}

	for _, m := range expectations {
		pw, err := crypto.MasterPassword(m.c, m.pt, m.u, m.pw, m.s)
		assert.NoError(t, err)
		assert.Equal(t, m.expect, pw)
	}

	for _, m := range expectations_bad {
		_, err := crypto.MasterPassword(m.c, m.pt, m.u, m.pw, m.s)
		assert.Error(t, err)
	}
}
