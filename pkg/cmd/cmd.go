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
