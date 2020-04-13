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

//Package station contains all commands under the subcommand 'station'
package station

import (
	"github.com/spf13/cobra"
)

//StationCmd represents parent command for all sub commands of station
var StationCmd = &cobra.Command{
	Use:   "station",
	Short: "Parent command for managing workstations",
}

func init() {
	StationCmd.AddCommand(listCmd)
	StationCmd.AddCommand(initCmd)
	StationCmd.AddCommand(deleteCmd)
}
