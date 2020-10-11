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

package docker

import (
	"fmt"
	"errors"
	"strings"
	"os"
	"archive/tar"
	"io"
	"path/filepath"

	"github.com/cspengl/conmake/pkg/agent"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"golang.org/x/net/context"
)

//DockerAgent models the docker agent and contains the docker client and
//configuration
type DockerAgent struct {
	ctx        context.Context
	endpoint   string
	apiversion string
	client     *client.Client
}

const (
	localEndpoint = "local"
	unixSocket    = "unix:///var/run/docker.sock"
	imagePrefix   = "docker.io"
	tmpDockerfile = "tmpDockerfile"
)

//NewDockerAgent creates new docker agent based on an endpoint and a API version
func NewDockerAgent(endpoint, apiversion string) (*DockerAgent, error) {

	if endpoint == "local" {
		endpoint = unixSocket
	}
	cli, err := client.NewClient(
		endpoint,
		apiversion,
		nil,
		nil)

	return &DockerAgent{
		ctx:        context.Background(),
		endpoint:   endpoint,
		apiversion: apiversion,
		client:     cli,
	}, err
}

//NewDockerAgentFromEnv returns a new docker agent based on the detected docker
//engine on the system
func NewDockerAgentFromEnv() (*DockerAgent, error) {
	cli, err := client.NewEnvClient()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	return &DockerAgent{
		ctx:        ctx,
		endpoint:   unixSocket,
		apiversion: "",
		client:     cli,
	}, err
}

// CreateStationContainer creates a new container
// on the docker runtime based on a station config
func (a *DockerAgent) CreateStationContainer(config agent.StationConfig) error {

	//Preparing container config
	containerConfig := &container.Config{
		AttachStdout: true,
		Image:        config.ImageID,
		Tty:          config.Process.Terminal,
		Cmd:          config.Process.Args,
		WorkingDir:   config.Process.Cwd,
	}

	//Preparing mounts
	var mounts = []mount.Mount{}
	for _, ociMount := range config.Mounts {
		mounts = append(mounts, mount.Mount{
			Type:   mount.Type(ociMount.Type),
			Source: ociMount.Source,
			Target: ociMount.Destination,
		})
	}

	//Preparing hostConfig
	hostConfig := &container.HostConfig{
		Mounts: mounts,
	}

	//Create the container
	_, err := a.client.ContainerCreate(
		a.ctx,
		containerConfig,
		hostConfig,
		nil,
		config.ContainerID,
	)

	return err
}

// RunStationContainer runs an existing station container
// by a given container id. If the quiet flag is false it will
// return a io.ReadCloser for reading the containers outputs.
func (a *DockerAgent) RunStationContainer(containerID string, quiet bool) (io.ReadCloser, error) {

	//Find docker id
	dockerID, err := a.findDockerID(containerID)

	if err != nil {
		return nil, err
	}

	//Start container by id
	err = a.client.ContainerStart(
		a.ctx,
		dockerID,
		types.ContainerStartOptions{},
	)

	if err != nil {
		return nil, err
	}

	if !quiet {
		return a.client.ContainerLogs(
			a.ctx,
			dockerID,
			types.ContainerLogsOptions{
				ShowStderr: true,
				ShowStdout: true,
				Follow:     true,
			},
		)
	}

	return nil, err
}

// DestroyStationContainer deletes an existing station container by a
// given id
func (a *DockerAgent) DestroyStationContainer(containerID string) error {

	//Find docker id
	dockerID, err := a.findDockerID(containerID)

	if err != nil {
		return err
	}

	//Remove container by id
	return a.client.ContainerRemove(
		a.ctx,
		dockerID,
		types.ContainerRemoveOptions{
			Force: true,
		},
	)
}

func (a *DockerAgent) BuildStation(imageID string, config agent.StationConfig) (error) {
	
	//Getting a dockerfile from the given config
	err, dockerfile := configToDockerfile(config)

	if err != nil {
		return err
	}

	//Write dockerfile to temporary file
	err = writeDockerfile(dockerfile)

	if err != nil {
		return err
	}

	//Get cwd
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	//taring cwd
	bCtx, err := tarBuildContext(cwd)

	if err != nil {
		return err
	}

	res, err := a.client.ImageBuild(
		a.ctx,
		bCtx,
		types.ImageBuildOptions{
			Tags: []string{imageID},
		},
	)

	if err != nil {
		return err
	}

	outputbuffer := make([]byte, 32*2048)
	for {
		_, readerError := res.Body.Read(outputbuffer)
		if readerError != nil {
			if readerError == io.EOF {break}
		}
	}

	defer os.Remove("Dockerfile")

	return err
}

// ImagePresent checks if an image is present in the underlying image store
// (local docker imagedb) and returns true if yes (false if not).
func (a *DockerAgent) ImagePresent(imageID string) (bool, error) {

	images, err := a.client.ImageList(
		a.ctx,
		types.ImageListOptions{
			All: true,
		},
	)

	if err != nil {
		return false, err
	}

	for _, img := range images {
		for _, tag := range img.RepoTags {
			if tag == imageID {
				return true, nil
			}
		}
	}

	return false, err
}

// DownloadImage downloads an image from the official Docker Registry
func (a *DockerAgent) DownloadImage(imageID string) (error) {
	progress, err := a.client.ImagePull(
		a.ctx,
		imagePrefix+"/library/"+imageID,
		types.ImagePullOptions{},
	)

	if err != nil {
		return err
	}

	progressBuffer := make([]byte, 32*2048)
	for {
		_, downloadErr := progress.Read(progressBuffer)
		if downloadErr != nil {
			if downloadErr == io.EOF {
				break
			}
		}
	}
	
	defer progress.Close()
	return err
}

// DeleteImage deletes an image specified by a imageID
func (a *DockerAgent) DeleteImage(imageID string) error {
	_, err := a.client.ImageRemove(
		a.ctx,
		imageID,
		types.ImageRemoveOptions{
			Force:         true,
			PruneChildren: true,
		},
	)

	return err
}

// Private functions

func (a *DockerAgent) findDockerID(containerName string) (string, error) {

	//Find containerid (container name)
	containers, err := a.client.ContainerList(
		a.ctx,
		types.ContainerListOptions{
			All: true,
		},
	)

	if err != nil {
		return "", err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == ("/" + containerName) {
				return container.ID, nil
			}
		}
	}

	return "", errors.New("Container not found")
}

func configToDockerfile(c agent.StationConfig) (error, string) {

	//FROM c.Base
	from := fmt.Sprintf("FROM %s\n", c.ImageID)
	
	//RUN c.Cmd
	cmd := fmt.Sprintf("RUN %s\n", strings.Join(c.Process.Args, " "))

	return nil, fmt.Sprintf("%s%s",from,cmd)
}

func writeDockerfile(dockerfile string) (error) {

	//Creat file
	df, err := os.Create("Dockerfile")

	if err != nil {
		return err
	}

	defer df.Close()

	df.WriteString(dockerfile)

	df.Sync()

	df.Close()

	return nil
}

func tarBuildContext(path string) (io.Reader, error) {

	pReader, pWriter := io.Pipe()

	tWriter := tar.NewWriter(pWriter)

	var err error
	go func() {	

		defer pWriter.Close()
		defer tWriter.Close()

		err = filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !fi.Mode().IsRegular() {
				return nil
			}

			//create tar header
			tarHeader, err := tar.FileInfoHeader(fi, fi.Name())
				
			if err != nil {
				return err
			}

			// update the name to correctly reflect the desired destination when untaring
			tarHeader.Name = strings.TrimPrefix(strings.Replace(file, path, "", -1), string(filepath.Separator))

			//write header
			if err := tWriter.WriteHeader(tarHeader); err != nil {
				return err
			}

			//open file
			f, err := os.Open(file)
			if err != nil {
				return err
			}

			//copy file data to tar writer
			if _, err := io.Copy(tWriter, f); err != nil {
				return err
			}

			f.Close()

			return nil
		})
	}()

	if err != nil {
		return nil, err
	}
	return pReader, nil
}
