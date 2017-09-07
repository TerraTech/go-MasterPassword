//==============================================================================
// This file is part of FQgolibs
// Copyright (c) 2017, FutureQuest, Inc.
//   https://www.FutureQuest.net
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//==============================================================================

package FQversion_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"futurequest.net/FQgolibs/FQtesting"

	"github.com/stretchr/testify/assert"
)

import "futurequest.net/FQgolibs/FQversion"

//#########################
const createTestDir = false

//#########################
var _ = assert.Equal
var D = FQtesting.D
var FATAL = FQtesting.FATAL
var ane = FQtesting.ANE
var ft = FQtesting.NewFQtesting(createTestDir)

func TestMain(m *testing.M) {
	ft.Begin()
	rc := m.Run()
	ft.End()
	os.Exit(rc)
}

//##########################

type tv struct {
	Name, Version, Build, String string
}

var vFQgolibs = tv{"FQgolibs", FQversion.VERSION, FQversion.BUILD, fmt.Sprintf("FQgolibs:\t%s\t%s", FQversion.VERSION, FQversion.BUILD)}

var testVersions = []tv{
	vFQgolibs,
	{"bar", "2.2.2", "20020202@12:12:12", "bar:\t2.2.2\t20020202@12:12:12"},
	{"foo", "1.1.1", "20010101@11:11:11", "foo:\t1.1.1\t20010101@11:11:11"},
	{"baz", "3.3.3", "20030303@13:13:13", "baz:\t3.3.3\t20030303@13:13:13"},
}

var testVersionsSorted = []tv{
	vFQgolibs,
	{"bar", "2.2.2", "20020202@12:12:12", "bar:\t2.2.2\t20020202@12:12:12"},
	{"baz", "3.3.3", "20030303@13:13:13", "baz:\t3.3.3\t20030303@13:13:13"},
	{"foo", "1.1.1", "20010101@11:11:11", "foo:\t1.1.1\t20010101@11:11:11"},
}

var tGVsorted []string

var tGVAsorted = []string{
	"FQzap:....4.4.4..20040404@14:14:14",
	fmt.Sprintf("FQgolibs:.%s.%s", vFQgolibs.Version, vFQgolibs.Build),
	"bar:......2.2.2..20020202@12:12:12",
	"baz:......3.3.3..20030303@13:13:13",
	"foo:......1.1.1..20010101@11:11:11",
}

var tSVsortedString string
var tSVAsortedString string

// tcv == testCallerVersion
var tcv = tv{"FQzap", "4.4.4", "20040404@14:14:14", "FQzap:\t4.4.4\t20040404@14:14:14"}

func init() {
	for _, tv := range testVersions[1:] { // index from [1:] to skip FQgolibs
		FQversion.Register(tv.Name, tv.Version, tv.Build)
	}
	tGVsorted = append(tGVsorted, tcv.String)
	for i, _ := range testVersionsSorted {
		tGVsorted = append(tGVsorted, testVersionsSorted[i].String)
	}
	tSVsortedString = strings.Join(tGVsorted, "\n")
	tSVAsortedString = strings.Join(tGVAsorted, "\n")
}

func TestCatalog(t *testing.T) {
	c := FQversion.Catalog()
	for i, _ := range testVersionsSorted { // need to range due to registeredVersions and tv impedence mismatch
		assert.EqualValues(t, testVersionsSorted[i], c[i])
	}
}

func TestShowCatalog(t *testing.T) {
	var e = vFQgolibs.String
	e += "\nbar:\t2.2.2\t20020202@12:12:12\nbaz:\t3.3.3\t20030303@13:13:13\nfoo:\t1.1.1\t20010101@11:11:11\n"
	sc := FQversion.ShowCatalog()
	//t.Log("\n" + sc)
	assert.Equal(t, e, sc)
}

func TestGetVersions(t *testing.T) {
	gv := FQversion.GetVersions(tcv.Name, tcv.Version, tcv.Build)
	assert.Equal(t, tGVsorted, gv)
}

func TestGetVersionsAligned(t *testing.T) {
	gva := FQversion.GetVersionsAligned(tcv.Name, tcv.Version, tcv.Build)
	assert.Equal(t, tGVAsorted, gva)
}

func TestShowVersions(t *testing.T) {
	sv := FQversion.ShowVersions(tcv.Name, tcv.Version, tcv.Build)
	assert.Equal(t, tSVsortedString, sv)
}

func TestShowVersionsAligned(t *testing.T) {
	sva := FQversion.ShowVersionsAligned(tcv.Name, tcv.Version, tcv.Build)
	assert.Equal(t, tSVAsortedString, sva)
}
