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

	"github.com/TerraTech/go-MasterPassword/pkg/common"
	"github.com/TerraTech/go-MasterPassword/pkg/config"
	"github.com/TerraTech/go-MasterPassword/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

var mpwseeds = []string{
	common.DefaultMasterPasswordSeed,
	"overrideDefaultMPWseed",
	"liveLifeBeyondAllYourTomorrrows",
	"danceLikeNoOneIsLooking",
}

type testVector struct {
	ms     string
	c      uint32
	pt     string
	pp     string
	expect string
}

// d == default
var d = struct {
	u, pw, s string
}{
	"user", "password", "example.com",
}

func newMpConfig(tv testVector) *config.MPConfig {
	return &config.MPConfig{
		MasterPasswordSeed: tv.ms,
		Counter:            tv.c,
		PasswordType:       tv.pt,
		PasswordPurpose:    tv.pp,
		Fullname:           d.u,
		Password:           d.pw,
		Site:               d.s,
	}
}

func newMpw(tv testVector) (*crypto.MasterPW, error) {
	mpw := &crypto.MasterPW{Config: &config.MPConfig{}}
	c := newMpConfig(tv)

	if err := mpw.MergeConfigEX(c); err != nil {
		return nil, err
	}

	return mpw, nil
}

func TestMasterPassword(t *testing.T) {
	expectations := []testVector{
		{mpwseeds[0], 1, "long", "auth", "ZedaFaxcZaso9*"},
		{mpwseeds[0], 2, "long", "auth", "Fovi2@JifpTupx"},
		{mpwseeds[0], 1, "maximum", "auth", "pf4zS1LjCg&LjhsZ7T2~"},
		{mpwseeds[0], 1, "medium", "auth", "ZedJuz8$"},
		{mpwseeds[0], 1, "basic", "auth", "pIS54PLs"},
		{mpwseeds[0], 1, "short", "auth", "Zed5"},
		{mpwseeds[0], 1, "pin", "auth", "6685"},
		{mpwseeds[0], 1, "name", "auth", "zedjuzoco"},
		{mpwseeds[0], 1, "phrase", "auth", "ze juzxo sax taxocre"},
	}

	expectations_bad := []testVector{
		{mpwseeds[0], 1, "invalidType", "auth", "1111"},
	}

	for _, tv := range expectations {
		pw, err := crypto.MasterPassword(tv.ms, tv.pt, tv.pp, d.u, d.pw, d.s, tv.c)
		assert.NoError(t, err)
		assert.Equal(t, tv.expect, pw)
	}

	for _, tv := range expectations_bad {
		_, err := crypto.MasterPassword(tv.ms, tv.pt, tv.pp, d.u, d.pw, d.s, tv.c)
		assert.Error(t, err)
	}

	// Test method call
	for _, tv := range expectations {
		mpw, err := newMpw(tv)
		assert.NoError(t, err)
		pw, err := mpw.MasterPassword()
		assert.NoError(t, err)
		assert.Equal(t, tv.expect, pw)
	}
}

func TestMasterPasswordSeeds(t *testing.T) {
	expectations := [][]testVector{
		{
			{mpwseeds[1], 1, "long", "auth", "NukiConqYocu1*"},
			{mpwseeds[1], 2, "long", "auth", "MiwkVuruDile0_"},
			{mpwseeds[1], 1, "maximum", "auth", "CR(m#EbdFijOx8u!bX1$"},
			{mpwseeds[1], 1, "medium", "auth", "NukKun1:"},
			{mpwseeds[1], 1, "basic", "auth", "CbL24Pbd"},
			{mpwseeds[1], 1, "short", "auth", "Nuk2"},
			{mpwseeds[1], 1, "pin", "auth", "5902"},
			{mpwseeds[1], 1, "name", "auth", "nukkunequ"},
			{mpwseeds[1], 1, "phrase", "auth", "nu kunno rom tolivna"},
		},
		{
			{mpwseeds[2], 1, "long", "auth", "GibuKaqoNeld5/"},
			{mpwseeds[2], 2, "long", "auth", "QuncPute3/Wuzk"},
			{mpwseeds[2], 1, "maximum", "auth", "a7?OMCHdbHoa1Q4&mc2)"},
			{mpwseeds[2], 1, "medium", "auth", "Gib9;Luq"},
			{mpwseeds[2], 1, "basic", "auth", "aiq91zOd"},
			{mpwseeds[2], 1, "short", "auth", "Gib9"},
			{mpwseeds[2], 1, "pin", "auth", "9779"},
			{mpwseeds[2], 1, "name", "auth", "gibmeluqe"},
			{mpwseeds[2], 1, "phrase", "auth", "gi melqo bod kahuwqa"},
		},
		{
			{mpwseeds[3], 1, "long", "auth", "Mobl2-BicuKasp"},
			{mpwseeds[3], 2, "long", "auth", "JeyzXawx5~Heye"},
			{mpwseeds[3], 1, "maximum", "auth", "z3)1NfH6^3B(octEFYFU"},
			{mpwseeds[3], 1, "medium", "auth", "Mob7(Rer"},
			{mpwseeds[3], 1, "basic", "auth", "zuP7mfR2"},
			{mpwseeds[3], 1, "short", "auth", "Mob7"},
			{mpwseeds[3], 1, "pin", "auth", "1317"},
			{mpwseeds[3], 1, "name", "auth", "moblirere"},
			{mpwseeds[3], 1, "phrase", "auth", "mobl rer nuksiri wuc"},
		},
	}

	for seedn, tvs := range expectations {
		for _, tv := range tvs {
			mpw, err := newMpw(tv)
			assert.NoError(t, err)
			err = mpw.SetMasterPasswordSeed(mpwseeds[seedn+1])
			assert.NoError(t, err)
			pw, err := mpw.MasterPassword()
			assert.NoError(t, err)
			assert.Equal(t, tv.expect, pw)
		}
	}
}

// TestMasterPasswordPurpose tests reasonable permutaions of MasterPasswordSeed and PasswordPurpose.
//
// Due to scrypt, these tests take a long time, therefore the tests will use the following matrix.
//   [seed#:purpose]
//   0:auth
//   1:ident
//   2:rec
func TestMasterPasswordPurposeGood(t *testing.T) {
	expectations := [][]testVector{
		{
			{mpwseeds[0], 1, "long", "auth", "ZedaFaxcZaso9*"},
			{mpwseeds[0], 2, "long", "auth", "Fovi2@JifpTupx"},
			{mpwseeds[0], 1, "maximum", "auth", "pf4zS1LjCg&LjhsZ7T2~"},
			{mpwseeds[0], 1, "medium", "auth", "ZedJuz8$"},
			{mpwseeds[0], 1, "basic", "auth", "pIS54PLs"},
			{mpwseeds[0], 1, "short", "auth", "Zed5"},
			{mpwseeds[0], 1, "pin", "auth", "6685"},
			{mpwseeds[0], 1, "name", "auth", "zedjuzoco"},
			{mpwseeds[0], 1, "phrase", "auth", "ze juzxo sax taxocre"},
		},
		{
			{mpwseeds[1], 1, "long", "ident", "PomfGidl9(Yamo"},
			{mpwseeds[1], 1, "maximum", "ident", "F8*W(egRYrmddQI^U(bn"},
			{mpwseeds[1], 1, "medium", "ident", "Pom1/Qil"},
			{mpwseeds[1], 1, "basic", "ident", "Ft61eGO9"},
			{mpwseeds[1], 1, "short", "ident", "Pom1"},
			{mpwseeds[1], 1, "pin", "ident", "1861"},
			{mpwseeds[1], 1, "name", "ident", "pomfuqilu"},
			{mpwseeds[1], 1, "phrase", "ident", "pom gidluzabi zuna"},
		},
		{
			{mpwseeds[2], 1, "long", "rec", "Viwa2_VaruJusk"},
			{mpwseeds[2], 1, "maximum", "rec", "lJNzh4i%^rz0amLkPe0_"},
			{mpwseeds[2], 1, "medium", "rec", "ViwJip9~"},
			{mpwseeds[2], 1, "basic", "rec", "lSo5rrW0"},
			{mpwseeds[2], 1, "short", "rec", "Viw5"},
			{mpwseeds[2], 1, "pin", "rec", "7245"},
			{mpwseeds[2], 1, "name", "rec", "viwjipubu"},
			{mpwseeds[2], 1, "phrase", "rec", "viwj pub daysixu jad"},
		},
	}

	for _, tvs := range expectations {
		for _, tv := range tvs {
			mpw, err := newMpw(tv)
			assert.NoError(t, err)

			err = mpw.SetMasterPasswordSeed(tv.ms)
			assert.NoError(t, err)

			pw, err := mpw.MasterPassword()
			assert.NoError(t, err)
			assert.Equal(t, tv.expect, pw)
		}
	}
}

func TestMasterPasswordPurposeBad(t *testing.T) {
	expectations := []struct {
		tv        testVector
		sppE, mpE error
	}{
		{testVector{mpwseeds[0], 1, "long", "noexist", "noneshallpass!"}, crypto.ErrPasswordPurposeInvalid, crypto.ErrPasswordPurposeInvalid},
		{testVector{mpwseeds[0], 69, "long", "noexist", "'tisbutascratch!"}, crypto.ErrPasswordPurposeInvalid, crypto.ErrPasswordPurposeInvalid},
		{testVector{mpwseeds[0], 69, "long", "ident", "ascratch?yourarmsoff"}, crypto.ErrPasswordPurposeInvalid, crypto.ErrPasswordPurposeCounterOutOfRange},
		{testVector{mpwseeds[0], 69, "long", "rec", "it'sjustafleshwound"}, crypto.ErrPasswordPurposeInvalid, crypto.ErrPasswordPurposeCounterOutOfRange},
	}

	for _, e := range expectations {
		mpw := &crypto.MasterPW{Config: newMpConfig(e.tv)}

		// test SetPasswordPurpose() validation path
		err := mpw.SetPasswordPurpose(e.tv.ms)
		if assert.Error(t, err) {
			assert.Equal(t, err, e.sppE)
		}

		// test MasterPassword() validation path via MergeConfig()
		_, err = mpw.MasterPassword()
		if assert.Error(t, err) {
			assert.Equal(t, err, e.mpE)
		}
	}
}
