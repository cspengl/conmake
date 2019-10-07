package station

import(
  "github.com/spf13/cobra"
)

var StationCmd = &cobra.Command{
  Use:    "station",
  Short:  "Parent command for managing stations for the given environment",
}

func init(){
  StationCmd.AddCommand(listCmd)
  StationCmd.AddCommand(initCmd)
}
