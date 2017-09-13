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

package common

// DefaultMasterPasswordSeed defaults to the universal seed as defined by:
// http://masterpasswordapp.com/algorithm.html
const (
	// config
	DefaultConfigFilename = "gompw.toml"
	DefaultCounter        = 1
	DefaultPasswordType   = "long"

	// crypto
	DefaultMasterPasswordSeed = "com.lyndir.masterpassword"
	DefaultPasswordPurpose    = "auth"
)
