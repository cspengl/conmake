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

package cmd

import (
  "fmt"
  "os"

  "github.com/cspengl/conmake/pkg/cmd/flags"
  "github.com/cspengl/conmake/pkg/cmd/station"
  "github.com/cspengl/conmake/pkg/cmd/do"

  "github.com/spf13/cobra"
)

var ConmakeCmd = &cobra.Command {
  Use:    "conmake",
  Short:  "Build tool running steps in containers",
  Long:   `
conmake is a command line tool similar to
make running the steps inside a container`,
  Run: func(cmd *cobra.Command, args []string){
    cmd.Help()
  },
}

func init(){
  //Adding commands
  ConmakeCmd.AddCommand(versionCmd)
  ConmakeCmd.AddCommand(station.StationCmd)
  ConmakeCmd.AddCommand(do.DoCmd)

  //Adding flags
  ConmakeCmd.PersistentFlags().StringVarP(&flags.ProjectPath, "path", "p", "./", "Absolute path to the project")
  ConmakeCmd.PersistentFlags().StringVarP(&flags.ConmakefilePath, "conmakefile", "f", "./Conmakefile.yaml", "Path to the Conmakefile to use")
}

func Execute() {
  if err := ConmakeCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
