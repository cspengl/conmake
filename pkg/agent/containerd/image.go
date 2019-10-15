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
  "log"
  "github.com/containerd/containerd/images"
  "github.com/containerd/containerd"
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

func (a *ContainerdAgent) commitImage(container containerd.Container, name string) (error){

  //getting task
  task, err := container.Task(a.ctx, nil)

  if err != nil {
    return err
  }

  //getting Checkpoint
  image, err := task.Checkpoint(
    a.ctx,
    containerd.WithCheckpointName(name),
  )

  if err != nil {
    return err
  }

  //construct metaimage
  metaimage := images.Image{
    Name: name,
    Target: image.Target(),
  }

  //add image to image service
  a.client.ImageService().Create(a.ctx, metaimage)

  log.Println("image committed")

  return err
}
