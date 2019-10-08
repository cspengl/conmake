package models

type Workstation struct {
  Base string           `yaml:"base"`
  Preparation []string  `yaml:"preparation"`
}
