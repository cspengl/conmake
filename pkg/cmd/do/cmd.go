package do

import(
  "github.com/cspengl/conmake/pkg/conmaker"

  "github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
  Use:    "do stepname",
  Short:  "Performs stepname from Conmakefile",
  Args:   cobra.ExactArgs(1),
  Run:    func(cmd *cobra.Command, args []string){
      run(args)
  },
}

func run(args []string){
  cm, err := conmaker.InitConmaker()
  if err != nil {
    panic(err)
  }

  err = cm.Perform(args[0])

  if err != nil {
    panic(err)
  }
}
