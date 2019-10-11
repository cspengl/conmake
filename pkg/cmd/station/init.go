package station

import (
  "github.com/cspengl/conmake/pkg/conmaker"
  "github.com/cspengl/conmake/pkg/cmd/flags"

  "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
  Use:    "init <station>",
  Short:  "Init station from Conmakefile",
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string){
     initStation(args[0])
  },
}


func initStation(stationName string) {

  cm, err := conmaker.InitConmaker(flags.ProjectPath, flags.ConmakefilePath)

  if err != nil {
    panic(err)
  }

  _ , err = cm.InitStation(stationName)

  if err != nil {
    panic(err)
  }

}
