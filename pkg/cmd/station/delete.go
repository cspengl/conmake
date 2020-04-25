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

package station

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/cspengl/conmake/pkg/cmd/utils"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <station>",
	Short: "Deletes existing station",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteStation(args[0])
	},
}

func deleteStation(station string) {
	cm, output, err := utils.ConmakerFromCmd()

	if err != nil {
		panic("Failed to create conmaker")
	}

	go func() {
		if err = cm.DeleteStation(station); err != nil {
			panic("Failed to delete station")
		}
	}()

	//Copy output to console
	io.Copy(os.Stdout, output)

	//Closing reader
	output.Close()
}
