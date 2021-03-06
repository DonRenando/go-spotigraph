// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"go.zenithar.org/spotigraph/build/mage/docker"
	"go.zenithar.org/spotigraph/build/mage/golang"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// Calculate file paths
var toolsBinDir = normalizePath(path.Join(curDir, "tools", "bin"))

func init() {
	time.Local = time.UTC

	// Add local bin in PATH
	err := os.Setenv("PATH", fmt.Sprintf("%s:%s", toolsBinDir, os.Getenv("PATH")))
	if err != nil {
		panic(err)
	}
}

type Code mg.Namespace

func (Code) Lint() {
	mg.Deps(Code.Format)

	color.Red("## Lint source")
	mg.Deps(golang.Lint("."))
}

func (Code) Format() {
	color.Red("## Formatting all sources")
	mg.SerialDeps(golang.Format, golang.Import)
}

func (Code) Licenser() error {
	mg.SerialDeps(golang.Format, golang.Import)

	color.Red("## Add license banner")
	return sh.RunV("go-licenser")
}

// -----------------------------------------------------------------------------

type API mg.Namespace

func (API) Generate() error {
	color.Blue("### Regenerate API")
	return sh.RunV("prototool", "all", "api/proto")
}

// -----------------------------------------------------------------------------

type Docker mg.Namespace

func (Docker) Spotigraph() error {
	return docker.Build(&docker.Command{
		Bin:         "spotigraph",
		Name:        "Spotigraph",
		Description: "Spotify agile model data microservice",
		URL:         "https://github.com/Zenitha/go-spotigraph/tree/master/cmd/spotigraph",
	})()
}

// -----------------------------------------------------------------------------

type Debug mg.Namespace

// Dockerfile is used to generate a Dockerfile from template in order to validate
// it with hadolint.
func (Debug) Dockerfile() error {
	return docker.Generate(&docker.Command{
		Bin:         "spotigraph",
		Name:        "Spotigraph",
		Description: "Spotify agile model data microservice",
		URL:         "https://github.com/Zenitha/go-spotigraph/tree/master/cmd/spotigraph",
	})()
}

// normalizePath turns a path into an absolute path and removes symlinks
func normalizePath(name string) string {
	absPath := mustStr(filepath.Abs(name))
	return absPath
}

func mustStr(r string, err error) string {
	if err != nil {
		panic(err)
	}
	return r
}
