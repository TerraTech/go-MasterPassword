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

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeBad(t *testing.T) {
	c := &MPConfig{}

	testFields2Merge = defaultFields2Merge()
	defer func() { testFields2Merge = nil }() // reset at end

	// field count mismatch
	//   perturb testFields2Merge to simulate a MPConfig struct member change
	// -1
	testFields2Merge["plusOne"] = struct{}{}
	assert.Panics(t, func() { c.Merge(c) })

	// +1
	delete(testFields2Merge, "plusOne")
	delete(testFields2Merge, "MasterPasswordSeed")
	assert.Panics(t, func() { c.Merge(c) })

	// field name rename
	testFields2Merge["FieldRenamed"] = struct{}{}
	assert.Panics(t, func() { c.Merge(c) })
}
