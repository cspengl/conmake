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
	"github.com/cspengl/conmake/pkg/cmd/flags"
	"github.com/cspengl/conmake/pkg/conmaker"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <station>",
	Short: "Init station from Conmakefile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		initStation(args[0])
	},
}

func initStation(stationName string) {

	cm, err := conmaker.InitConmaker(flags.ProjectPath, flags.ConmakefilePath)

	if err != nil {
		panic(err)
	}

	_, err = cm.InitStation(stationName)

	if err != nil {
		panic(err)
	}

}
