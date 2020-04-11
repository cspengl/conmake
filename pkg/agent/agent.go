/*
Copyright 2020 cspengl

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

//Package agent contains the generic agent definition
package agent

import (
	"io"

	ocispec "github.com/opencontainers/runtime-spec/specs-go"
)

//StationConfig models the configuration of a station to be spinned up and used
type StationConfig struct {
	// ContainerID is the id of the station container
	ContainerID string
	// ImageID is the id of the image for the station to use 
	// (either base or existing station image)
	ImageID 	string
	// Mounts are OCI Mounts for the station container (workspace mount)
	Mounts		[]ocispec.Mount
	// Process is the OCI process to execute on the staton container
	Process		ocispec.Process
	// User is the OCI user which executes the specified process
	User		ocispec.User
	// Output is a ReadCloser for processing the output of the process (standard output)
	Output		io.ReadCloser
}

// Agent defines a generic interface for working with
// station containers on a OCI runtime
type Agent interface {
	// ImagePresent returns if an image specific by imageID exists.
	ImagePresent(imageID string) (bool, error)
	// DownloadImage downloads an image specified by imageID and returns
	// a io.ReadCloser for processing the download (e.g. progress)
	DownloadImage(imageID string) (io.ReadCloser, string)
	// CreateStationContainer creates a container based on a StationConfig.
	CreateStationContainer(config StationConfig) (error)
	// RunStationContainer runs a created station container specified by a containerID
	RunStationContainer(containerID string) (error)
	// DestroyStationContainer destroys a station container specified by a containerID
	DestroyStationContainer(containerID string) (error)
	// SaveStationContainer saves a stopped station container specified by 
	// a containerID as a new image with the given imageID 
	SaveStationContainer(containerID, imageID string) (error)
}