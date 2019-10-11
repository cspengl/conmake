package station

import(
  "github.com/spf13/cobra"
)

var StationCmd = &cobra.Command{
  Use:    "station",
  Short:  "Parent command for managing workstations",
}

func init(){
  StationCmd.AddCommand(listCmd)
  StationCmd.AddCommand(initCmd)
  StationCmd.AddCommand(deleteCmd)
  StationCmd.AddCommand(cleanCmd)
}
