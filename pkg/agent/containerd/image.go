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
  "github.com/containerd/containerd"
  "github.com/containerd/containerd/diff/apply"
  "github.com/containerd/containerd/rootfs"
  "github.com/containerd/containerd/images"

  "github.com/opencontainers/image-spec/identity"
)

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

func (a *ContainerdAgent) removeImage(name string) (error) {
  return a.client.ImageService().Delete(a.ctx, name)
}

func (a *ContainerdAgent) commitImage(container containerd.Container, name string) (error){

  //getting image
  image, err := container.Image(a.ctx)

  if err != nil {
    return err
  }

  //getting diffIDs
  diffIDs, err := image.RootFS(a.ctx)

  //gettings parent
  parent := identity.ChainID(diffIDs).String()

  //getting containermodel
  containermodel, err := a.client.ContainerService().Get(a.ctx, container.ID())

  if err != nil {
    return err
  }

  //getting snappshotter
  snapshotter := a.client.SnapshotService(containermodel.Snapshotter)

  //create active snapshot
  mounts, err := snapshotter.Prepare(a.ctx, name, parent)

  if err != nil {
    return err
  }

  //creating diff to base image
  diff, err := rootfs.CreateDiff(a.ctx, containermodel.SnapshotKey, snapshotter, a.client.DiffService())

  if err != nil {
    return err
  }

  //getting applier
  applier := apply.NewFileSystemApplier(a.client.ContentStore())

  //apply diff
  _, err = applier.Apply(a.ctx, diff, mounts)

  if err != nil {
    return err
  }

  //construct image
  imageModel := images.Image{
    Name: name,
    Target: image.Target(),
  }

  //add image to contentstore
  _, err = a.client.ImageService().Create(a.ctx, imageModel)

  return err
}
