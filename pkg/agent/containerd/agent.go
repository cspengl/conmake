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
  "github.com/containerd/containerd/oci"
  "github.com/containerd/containerd/namespaces"
  //"github.com/containerd/containerd/cio"

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

  present, err := a.imageExists(c.Image)

  if err != nil {
    return "", err
  }

  if !present {
    err = a.downloadImage(c.Image)
  }

  if err != nil {
    return "", err
  }

  container, err := a.createContainer(
    agent.ConstructStationContainerName(c),
    c.Image,
  )

  if err != nil {
    return "", err
  }

  return container.ID(), err
}

func (a *ContainerdAgent) DeleteStation(c *agent.StationConfig) error {
  return a.deleteContainer(agent.ConstructStationContainerName(c))
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

func (a *ContainerdAgent) downloadImage(image string) error {

  _, err := a.client.Pull(
    a.ctx,
    imageRef+"/"+image,
    containerd.WithPullUnpack,
  )

  return err
}

func (a *ContainerdAgent) deleteImage(image string) error {

  err := a.client.ImageService().Delete(
    a.ctx,
    imageRef+"/"+image,
  )

  return err
}

func (a *ContainerdAgent) imageExists(image string) (bool, error){

  images, err := a.client.ListImages(a.ctx)

  if err != nil{
    return false, err
  }

  for _, img := range images {
    if img.Name() == image {
      return true, err
    }
  }

  return false, err

}

func (a *ContainerdAgent) createContainer(id, image string) (containerd.Container, error) {

    img, err := a.client.GetImage(a.ctx, imageRef+"/"+image)

    if err != nil {
      return nil, err
    }

    return a.client.NewContainer(
      a.ctx,
      id,
      containerd.WithNewSpec(oci.WithImageConfig(img)),
      // containerd.WithNewSnapshot(id+"-rootfs", img),
    )
}

func (a *ContainerdAgent) deleteContainer(id string) (error) {
  return a.client.ContainerService().Delete(a.ctx, id)
}
