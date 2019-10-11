package station

import (
  "github.com/cspengl/conmake/pkg/conmaker"
  "github.com/cspengl/conmake/pkg/cmd/flags"

  "github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
  Use:    "clean <station>",
  Short:  "Creates a new station from base image",
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string){
      cleanStation(args[0])
  },
}

func cleanStation(stationName string){
  cm, err := conmaker.InitConmaker(flags.ProjectPath, flags.ConmakefilePath)

  if err != nil {
    panic(err)
  }

  err = cm.CleanStation(stationName)

  if err != nil {
    panic(err)
  }
}
