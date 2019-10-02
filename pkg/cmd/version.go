package cmd

import(
  "fmt"

  "github.com/spf13/cobra"
)

var Version string //will be set by -ldflags
/*go build|run|install [...] -ldflags "-X github.com/cspengl/conmake/pkg/cmd.Version=<version>" [.. .]*/


var versionCmd = &cobra.Command{
  Use:    "version",
  Short:  "Print the version number",
  Run:    func(cmd *cobra.Command, args []string){
      fmt.Printf("conmake [version: %v]\n", Version)
  },
}
