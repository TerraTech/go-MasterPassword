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

package debug

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"futurequest.net/FQgolibs/FQdebug"
)

var globalDebug *Debug

var (
	D        = FQdebug.D
	MP_DEBUG string
)

type Debug struct {
	enabled bool
	files   []string
}

func init() {
	globalDebug = NewDebug()
	mpDebug := os.Getenv("MP_DEBUG")
	if mpDebug != "" {
		globalDebug.files = strings.Split(mpDebug, ",")
		globalDebug.enabled = len(globalDebug.files) > 0
	}
}

func NewDebug() *Debug {
	if globalDebug != nil {
		return globalDebug
	}
	globalDebug = &Debug{
		files: make([]string, 0, 1),
	}

	return globalDebug
}

func (d *Debug) Dbg(format string, a ...interface{}) {
	if !d.enabled {
		return
	}

	_, f, _, ok := runtime.Caller(1)
	if !ok {
		return
	}

	if !d.wantDebug(f) {
		return
	}

	logIt(format, a...)
}

func (d *Debug) DbgO(format string, a ...interface{}) {
	logIt(format, a...)
}

func (d *Debug) SetFilename(f string) {
	// Stores the filename normalized
	d.files = append(d.files, normalize(f))
	d.enabled = true
}

// wantDebug will do a suffix based match for given filepath
//
//  e.g. /^.*${fp}$/
func (d *Debug) wantDebug(cf string) bool {
	//cf == caller filename
	cf = normalize(cf)
	for _, fn := range d.files {
		if fn == "all" || strings.HasSuffix(cf, fn) {
			return true
		}
	}

	return false
}

func logIt(format string, a ...interface{}) {
	log.Printf("[DEBUG] "+format, a...)
}

func normalize(f string) string {
	// 1) make sure it has a leading slash
	// 2) cleaned up
	// 3) normalized for forward slash (Unix style)
	f = filepath.Clean("/" + f)
	f = filepath.ToSlash(f)

	return f
}