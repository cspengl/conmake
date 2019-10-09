package models

type Workstation struct {
  Base string           `yaml:"base"`
  Autoinit bool         `yaml:"autoinit"`
  Preparation []string  `yaml:"preparation"`
}
