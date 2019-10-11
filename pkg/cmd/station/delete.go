package station

import (
  "github.com/cspengl/conmake/pkg/conmaker"
  "github.com/cspengl/conmake/pkg/cmd/flags"

  "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
  Use:    "delete <station>",
  Short:  "Deletes existing station",
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string){
      deleteStation(args[0])
  },
}

func deleteStation(stationName string){
  cm, err := conmaker.InitConmaker(flags.ProjectPath, flags.ConmakefilePath)

  if err != nil {
    panic(err)
  }

  err = cm.DeleteStation(stationName)

  if err != nil {
    panic(err)
  }
}
