package station

import (
  "github.com/cspengl/conmake/pkg/conmaker"

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
  cm, err := conmaker.InitConmaker()

  if err != nil {
    panic(err)
  }

  err = cm.DeleteStation(stationName)

  if err != nil {
    panic(err)
  }
}
