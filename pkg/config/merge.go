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
	"reflect"
)

// fields2Merge is a validation map of all MPConfig fields that will be merged
type fields2MergeT map[string]struct{}

var (
	// testFields2Merge is used for test overrides
	testFields2Merge    fields2MergeT
	defaultFields2Merge = func() fields2MergeT {
		// use a func() to ensure it stays pristine on invocation during tests
		return fields2MergeT{
			"MasterPasswordSeed": struct{}{},
			"PasswordType":       struct{}{},
			"Fullname":           struct{}{},
			"Password":           struct{}{},
			"PasswordPurpose":    struct{}{},
			"Site":               struct{}{},
			"Counter":            struct{}{},
		}
	}
)

// Merge will merge c ==> mpc for any nil entries
func (mpc *MPConfig) Merge(c *MPConfig) {
	// safetynet to make sure a member change to MPConfig is caught
	fields2merge := defaultFields2Merge()
	if testFields2Merge != nil {
		// test override
		fields2merge = testFields2Merge
	}

	// These are whitelisted fields used for internal debugging
	whitelisted := map[string]bool{
		"ConfigFile": true,
		"dump":       true,
	}

	v := reflect.ValueOf(mpc).Elem()
	vNF := v.NumField()
	// need to compensate for the added debugging fields
	if vNF != len(fields2merge)+len(whitelisted) {
		println(vNF, len(fields2merge))
		panic("config.Merge mismatch with MPConfig, please check the merge actions and/or MPConfig members")
	}

	for i := 0; i < vNF; i++ {
		field := v.Type().Field(i)
		if _, exists := fields2merge[field.Name]; !exists {
			if whitelisted[field.Name] {
				continue
			}
			panic("config.Merge() would not transfer all MPConfig members due to add/del/renames: " + field.Name)
		}
	}

	if mpc.MasterPasswordSeed == "" {
		mpc.MasterPasswordSeed = c.MasterPasswordSeed
	}
	if mpc.PasswordType == "" {
		mpc.PasswordType = c.PasswordType
	}
	if mpc.Fullname == "" {
		mpc.Fullname = c.Fullname
	}
	if mpc.Password == "" {
		mpc.Password = c.Password
	}
	if mpc.PasswordPurpose == "" {
		mpc.PasswordPurpose = c.PasswordPurpose
	}
	if mpc.Site == "" {
		mpc.Site = c.Site
	}
	if mpc.Counter == 0 {
		mpc.Counter = c.Counter
	}
}
