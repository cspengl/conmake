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

package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/cspengl/conmake/api/v1"
)

func ConmakefileFromFile(path string) (*v1.Conmakefile, error) {
	//Read file
	f, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Conmakefile not found")
		return nil, err
	}

	return ConmakefileFromByte(f);
}

//FromPath parses a Conmakfile from bytes into a Conmakefile struct
func ConmakefileFromByte(data []byte) (*v1.Conmakefile, error) {
	c := v1.Conmakefile{}
	err := yaml.Unmarshal(data, &c)
	return &c, err
}