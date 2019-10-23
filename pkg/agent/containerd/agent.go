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

//Package containerd contains the implementation of an agent connecting
//to the containerd daemon on the machine
package containerd

import(
  "fmt"

  "github.com/cspengl/conmake/pkg/agent"

  "github.com/containerd/containerd"
  "github.com/containerd/containerd/namespaces"

  "github.com/opencontainers/runtime-spec/specs-go"

  "golang.org/x/net/context"
)

const(
  conmakeNamespace = "conmake"
  imageRef   = "docker.io/library"
  unixSocket = "/run/containerd/containerd.sock"
)

//ContainerdAgent models the agent for containerd
type ContainerdAgent struct {
    ctx         context.Context
    client      *containerd.Client
}

func NewContainerdAgent() (*ContainerdAgent, error) {

  //Creating context
  ctx := context.Background()

  cli, err := containerd.New(
    unixSocket,
    containerd.WithDefaultNamespace(conmakeNamespace),
  )

  if err != nil {
    return nil, err
  }

  //Create namespace if it does not already exits
  err = cli.NamespaceService().Create(ctx, conmakeNamespace, nil)


  //Create context of namespace
  ctx = namespaces.WithNamespace(ctx, conmakeNamespace)

  return &ContainerdAgent{
    ctx:      ctx,
    client:   cli,
  }, err
}

func (a *ContainerdAgent) PerformStep(c *agent.StationConfig) error {
  return nil
}


func (a *ContainerdAgent) InitStation(c *agent.StationConfig, existing bool) (string, error) {

  //spinning up station
  container, err := a.spinupStation(c)

  if err != nil {
    panic(err)
    return "", err
  }

  //committing new image from created container
  if err = a.commitImage(container, "testimage"); err != nil {
    panic(err)
    return "", err
  }

  //spinnung down station
  if err = a.stopContainer(container.ID()); err != nil {
    return "", err
  }

  //deleting container
  if err = a.deleteContainer(container.ID()); err != nil{
    return "", err
  }

  //print list
  a.StationList(c.ProjectName)

  return container.ID(), err

}

func (a *ContainerdAgent) DeleteStation(c *agent.StationConfig) error {

  err := a.deleteContainer(agent.ConstructStationContainerName(c))
  return err
}

func (a *ContainerdAgent) StationExists(c *agent.StationConfig) (bool, error) {
  return false, nil
}

func (a *ContainerdAgent) StationList(projectname string) (error) {

  images, err := a.client.ListImages(
    a.ctx,
  )

  if err != nil {
    return err
  }

  for _, img := range images{
    fmt.Printf("%v\n", img.Name())
  }

  return err
}

func (a *ContainerdAgent) Info() {

}

func (a *ContainerdAgent) spinupStation(c *agent.StationConfig) (containerd.Container, error) {
  present, err := a.imageExists(c.Image)

  if err != nil {
    return nil, err
  }

  if !present {
    err = a.downloadImage(c.Image)
  }

  if err != nil {
    return nil, err
  }

  cwd := agent.Workspace

  mount := specs.Mount{
    Destination: cwd,
    Type: "none",
    Source: c.Workspace,
    Options: []string{"rbind", "rw"},
  }

  script := []string{
    "sh",
    "-c",
    agent.GenShellScript(c.Script),
  }

  cont := container{
    ID: agent.ConstructStationContainerName(c),
    Image_ID: c.Image,
    Mounts: []specs.Mount{mount},
    Cmd: script,
    WorkdingDir: agent.Workspace,
  }

  fmt.Printf("%v\n", cont)

  container, err := a.createContainer(cont)

  if err != nil {
    panic(err)
    return nil, err
  }

  err = a.runContainer(container.ID())

  if err != nil {
    panic(err)
    return nil, err
  }

  return container, err
}
