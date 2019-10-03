package cmd

import(
  "fmt"

  "github.com/cspengl/conmake/pkg/utils"

  "github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
  Use:    "version",
  Short:  "Print the version number",
  Run:    func(cmd *cobra.Command, args []string){
      fmt.Printf("conmake [version: %v]\n", utils.Version)
  },
}
