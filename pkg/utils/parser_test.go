/*
Copyright 2020 cspengl

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cspengl/conmake/pkg/utils"
)

const (
	filePath    = "/../../testdata/Conmakefile.yaml"
	invalidPath = "/../../testdata/NotExisting.yaml"
)

var cwd, _ = os.Getwd()

func TestConmakeFileFromFile(t *testing.T) {
	_, err := utils.ConmakefileFromFile(cwd + filePath)
	if err != nil {
		t.Fail()
	}
}

func TestConmakeFileFromInvalidFile(t *testing.T) {
	_, err := utils.ConmakefileFromFile(cwd + invalidPath)

	t.Log(err)
	if err == nil {
		t.Error("Failed")
	}
}

func TestConmakefileFromByte(t *testing.T) {
	f, err := ioutil.ReadFile(cwd + filePath)

	if err != nil {
		t.Fail()
	} else {
		_, err := utils.ConmakefileFromByte(f)
		if err != nil {
			t.Error(err)
		}
	}
}
