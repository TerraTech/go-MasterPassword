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

package FQversion

import (
	"bytes"
	"fmt"
	"sort"
)

var (
	catalog = []registeredVersion{}
)

type registeredVersion struct {
	Name    string
	Version string
	Build   string
	String  string
}

func Catalog() []registeredVersion {
	if !sort.IsSorted(byName(catalog)) {
		sort.Sort(byName(catalog))
	}

	return catalog
}

func ShowCatalog() string {
	var buf bytes.Buffer
	for rv := range _nvb() {
		buf.WriteString(rv + "\n")
	}

	return buf.String()
}

func ShowCatalogAligned() string {
	var buf bytes.Buffer
	tw := newTabWriter(&buf)

	for rv := range _nvb() {
		fmt.Fprintln(tw, rv)
	}
	tw.Flush()

	return buf.String()
}

func _nvb() <-chan string {
	ch := make(chan string, 10)
	go func() {
		for _, rv := range Catalog() {
			ch <- nvb(rv.Name, rv.Version, rv.Build)
		}
		close(ch)
	}()
	return ch
}

type byName []registeredVersion

func (n byName) Len() int           { return len(n) }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byName) Less(i, j int) bool { return n[i].Name < n[j].Name }
