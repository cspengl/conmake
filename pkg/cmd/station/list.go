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
	"github.com/spf13/cobra"

	"github.com/cspengl/conmake/pkg/cmd/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists existing stations",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	cm, err := utils.ConmakerFromCmd()

	if err != nil {
		panic(err)
	}

	err = cm.StationList()

	if err != nil {
		panic(err)
	}
}
