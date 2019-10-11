package main

import(
  "log"
  "github.com/cspengl/conmake/pkg/cmd"

  "github.com/spf13/cobra/doc"
)

func main(){
  err := doc.GenMarkdownTree(cmd.ConmakeCmd, "docs/reference/cli/")
  if err != nil {
    log.Fatal(err)
  }
}
