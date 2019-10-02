package station

import (
  "fmt"
  
  "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
  Use:    "list",
  Short:  "Lists existing stations",
  Run:    func(cmd *cobra.Command, args []string){
      list()
  },
}

func list(){
  fmt.Println("Has to be implemented")
}
