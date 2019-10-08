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

func NewConmakefile(data []byte) (*Conmakefile, error) {
  c := Conmakefile{}
  err := yaml.Unmarshal(data, &c)
  return &c, err
}
