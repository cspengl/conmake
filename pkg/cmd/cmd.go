package cmd

import (
  "fmt"
  "os"
  "io/ioutil"

  "github.com/cspengl/conmake/pkg/cmd/station"
  "github.com/cspengl/conmake/pkg/conmakefile"

  "github.com/spf13/cobra"
)

var conmakeCmd = &cobra.Command {
  Use:    "conmake",
  Short:  "Build tool running inside docker container",
  Long:   `conmake is a command line tool similar to
           make or cmake running the steps inside a docker container`,
  Run: func(cmd *cobra.Command, args[] string){
    file, _ := ioutil.ReadFile("examples/Conmakefile.yaml")
    c := conmakefile.Parse(file)

    fmt.Printf("Version: %v\nProject: %v\nAgent: %v\n", c.Version, c.Project, c.Agent)
  },
}

func init(){
  conmakeCmd.AddCommand(versionCmd)
  conmakeCmd.AddCommand(station.StationCmd)
}

func Execute() {
  if err := conmakeCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
