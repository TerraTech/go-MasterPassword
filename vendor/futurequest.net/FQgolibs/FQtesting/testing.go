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

package FQtesting

import (
	"fmt"
	"log"
	"os"
	"testing"

	"futurequest.net/FQgolibs/FQdebug"

	"github.com/stretchr/testify/assert"
)

const d_test = "/tmp/GoTesting"

type FQtesting struct {
	D_test        string
	createTestDir bool
}

func NewFQtesting(createTestDir bool) *FQtesting {
	return &FQtesting{
		D_test:        d_test,
		createTestDir: createTestDir,
	}
}

func ANE(t *testing.T, err error) bool {
	if !assert.NoError(t, err) {
		t.Fatal(err)
		return false
	}

	return true
}

// D is wrapper around FQdebug.D
func D(args ...interface{}) {
	FQdebug.D(args...)
}

func FATAL(args ...interface{}) {
	fmt.Println("\tFATAL:", args)
	os.Exit(1)
}

// CGT is Create Go Test directory
func (ft FQtesting) CGT() {
	if !ft.createTestDir {
		return
	}
	os.RemoveAll(ft.D_test)
	err := os.Mkdir(ft.D_test, 0770)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// RGT is Remove Go Test directory
func (ft FQtesting) RGT() {
	if ft.createTestDir {
		os.RemoveAll(ft.D_test)
	}
}

func (ft FQtesting) Begin() {
	ft.CGT()
}

func (ft FQtesting) End() {
	ft.RGT()
}
