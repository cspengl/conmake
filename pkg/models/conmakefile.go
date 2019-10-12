/*
Copyright 2019 cspengl

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

package models

import(
  "gopkg.in/yaml.v2"
)

type Conmakefile struct {
  Version string                        `yaml:"version"`
  Project string                        `yaml:"project"`
  Steps map[string]Step                 `yaml:"steps`
  Workstations map[string]Workstation   `yaml:"workstations`
}

type Step struct {
  Workstation string    `yaml:"workstation"`
  Script []string       `yaml:"script"`
}

type Workstation struct {
  Base string           `yaml:"base"`
  Autoinit bool         `yaml:"autoinit"`
  Script []string       `yaml:"preparation"`
}

func NewConmakefile(data []byte) (*Conmakefile, error) {
  c := Conmakefile{}
  err := yaml.Unmarshal(data, &c)
  return &c, err
}
