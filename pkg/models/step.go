package models

import(
  "strings"
)

type Step struct {
  Workstation string    'yaml:"workstation"'
  Script []string       'yaml:"script"'
  Artifacts []Artifact  'yaml:"artifacts"'
}

func  ScriptToCmd(script []string) []string{
  cmds := []string{}
  for _, cmd := range script{
     cmds = append(cmds, strings.Fields(cmd)...)
     cmds = append(cmds, " && ")
  }
  return cmds[:len(cmds) - 1]
}
