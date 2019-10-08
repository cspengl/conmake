package station

import (
  "github.com/cspengl/conmake/pkg/conmaker"

  "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
  Use:    "init stationname",
  Short:  "Init station from Conmakfile",
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string){
     initStation(args[0])
  },
}


func initStation(stationName string) {

  cm, err := conmaker.InitConmaker()

  if err != nil {
    panic(err)
  }

  err = cm.InitStation(stationName)

  if err != nil {
    panic(err)
  }

}
