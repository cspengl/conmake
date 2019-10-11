package cmd

import (
  "fmt"
  "os"

  "github.com/cspengl/conmake/pkg/cmd/station"
  "github.com/cspengl/conmake/pkg/cmd/do"

  "github.com/spf13/cobra"
)

var ConmakeCmd = &cobra.Command {
  Use:    "conmake",
  Short:  "Build tool running inside docker container",
  Long:   `conmake is a command line tool similar to
           make or cmake running the steps inside a docker container`,
  Run: func(cmd *cobra.Command, args []string){
    cmd.Help()
  },
}

func init(){
  ConmakeCmd.AddCommand(versionCmd)
  ConmakeCmd.AddCommand(station.StationCmd)
  ConmakeCmd.AddCommand(do.DoCmd)
}

func Execute() {
  if err := ConmakeCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
