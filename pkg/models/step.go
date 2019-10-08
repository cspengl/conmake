package models

type Step struct {
  Workstation string    `yaml:"workstation"`
  Script []string       `yaml:"script"`
}
