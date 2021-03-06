/*
Copyright 2019 cspengl

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Package do contains all commands under the sub command 'do'
package do

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/cspengl/conmake/pkg/cmd/utils"
)

//DoCmd represents the command performing a step from a Conmakfile
var DoCmd = &cobra.Command{
	Use:   "do <stepname> [args...]",
	Short: "Performs stepname from Conmakefile",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		run(args[0], args[1:])
	},
}

func run(step string, args []string) {
	cm, output, err := utils.ConmakerFromCmd()

	if err != nil {
		panic("Failed to create conmaker")
	}

	go func() {
		if err = cm.PerformStep(step, args...); err != nil {
			fmt.Println("Failed to perform step")
			os.Exit(1)
		}
	}()

	//Copy output to console
	io.Copy(os.Stdout, output)

	//Closing reader
	output.Close()

}
