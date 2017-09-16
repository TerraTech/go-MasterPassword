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
	"fmt"
	"os"
	"text/template"
)

const mpconfig = `-----------------
configFile         : {{ddd .ConfigFile}}
masterPasswordSeed : {{ddd .MasterPasswordSeed}}
fullName           : {{ddd .Fullname}}
password           : {{ddd .Password}}
passwordType       : {{ddd .PasswordType}}
siteName           : {{ddd .Site}}
siteCounter        : {{itoa .Counter | ddd}}
-----------------
`

var funcMap = template.FuncMap{
	"ddd":  ddd,
	"itoa": itoa,
}

// Dump will dump formatted output of the user configuration file.
func (c *MPConfig) Dump() error {
	t := template.Must(template.New("mpconfig").Funcs(funcMap).Parse(mpconfig))
	err := t.Execute(os.Stderr, c)

	return err
}

// ddd = dotdotdot for "" elements
func ddd(v string) string {
	if v != "" {
		return v
	}

	return "..."
}

func itoa(v uint32) string {
	if v == 0 {
		return ""
	}

	return fmt.Sprintf("%v", v)
}
