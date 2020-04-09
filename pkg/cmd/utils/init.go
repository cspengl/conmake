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

//Package cmd is the parent package for all commands
//Contains the root command and the version command
package utils

import (
	"os"

	"github.com/cspengl/conmake/pkg/conmaker"

	"github.com/cspengl/conmake/pkg/agent"
	"github.com/cspengl/conmake/pkg/agent/docker"

	"github.com/cspengl/conmake/pkg/cmd/flags"

	"github.com/cspengl/conmake/pkg/utils"
)

func ConmakerFromCmd() (*conmaker.Conmaker, error) {
	
	//Reading Conmakefile from path
	cmConmakefile, err := utils.ConmakefileFromFile(flags.ConmakefilePath)

	if err != nil {
		return nil, err
	}

	//Creating agent
	var cmAgent agent.Agent
	switch flags.Agent {
	case agent.SignDocker:
		cmAgent, err = docker.NewDockerAgentFromEnv()
		break
	default:
		cmAgent, err = docker.NewDockerAgentFromEnv()
	}

	if err != nil {
		return nil, err
	}

	//Reading project path
	cmProjectpath := flags.ProjectPath
	if cmProjectpath == "./" {
		cmProjectpath, err = os.Getwd()
	}

	//Return conmaker
	return conmaker.NewConmaker(
		cmAgent,
		cmConmakefile,
		cmProjectpath), nil
}



