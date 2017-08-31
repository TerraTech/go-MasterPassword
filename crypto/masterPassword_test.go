// Copyright 2017 FutureQuest, Inc.

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
