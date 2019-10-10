package models

type Workstation struct {
  Base string           `yaml:"base"`
  Autoinit bool         `yaml:"autoinit"`
  Script []string  `yaml:"preparation"`
}
