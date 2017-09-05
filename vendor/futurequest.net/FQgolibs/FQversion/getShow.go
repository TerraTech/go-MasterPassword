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
	"os"
	"strings"
	"time"
)

type Version registeredVersion

func NewVersion(name, version, build string) *Version {
	return &Version{
		Name:    name,
		Version: version,
		Build:   build,
	}
}

func (v *Version) GetVersions() []string        { return GetVersions(v.Name, v.Version, v.Build) }
func (v *Version) GetVersionsAligned() []string { return GetVersionsAligned(v.Name, v.Version, v.Build) }
func (v *Version) ProgVersion() string          { return ProgVersion(v.Name, v.Version, v.Build) }
func (v *Version) ShowVersions() string         { return ShowVersions(v.Name, v.Version, v.Build) }
func (v *Version) ShowVersionsAligned() string  { return ShowVersionsAligned(v.Name, v.Version, v.Build) }

// GetBUILD returns a standard formatted BUILD date string
// e.g. date '+%Y%m%d@%T'
func GetBUILD() string {
	return time.Now().Format("20060102@15:04:05")
}

// GetBUILDHOST returns the server this was compiled on
func GetBUILDHOST() string {
	host, err := os.Hostname()
	if err != nil {
		host = err.Error()
	}

	return host
}

// GetVersions return a list of Caller and Catalog versions
func GetVersions(name, version, build string) []string {
	v := make([]string, 0, len(catalog)+1)
	v = append(v, nvb(name, version, build))
	for _, cv := range Catalog() {
		v = append(v, cv.String)
	}

	return v
}

func GetVersionsAligned(name, version, build string) []string {
	var buf bytes.Buffer
	sv := GetVersions(name, version, build)
	va := make([]string, len(sv))
	tw := newTabWriter(&buf)

	for _, v := range sv {
		fmt.Fprintln(tw, v)
	}
	tw.Flush()

	for i, v := range bytes.Split(bytes.TrimSpace(buf.Bytes()), []byte("\n")) {
		va[i] = string(v)
	}

	return va
}

func ProgVersion(name, version, build string) string {
	return fmt.Sprintf("%s: %s (%s)", name, version, build)
}

func ShowVersions(name, version, build string) string {
	return strings.Join(GetVersions(name, version, build), "\n")
}

func ShowVersionsAligned(name, version, build string) string {
	return strings.Join(GetVersionsAligned(name, version, build), "\n")
}
