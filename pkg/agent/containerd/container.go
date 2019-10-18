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
  "log"
  "errors"
  "syscall"

  "github.com/containerd/containerd"
  "github.com/containerd/containerd/oci"
  "github.com/containerd/containerd/cio"

  "github.com/opencontainers/runtime-spec/specs-go"
)

type container struct {
  ID          string
  Image_ID    string
  Mounts      []specs.Mount
  WorkdingDir string
  Cmd         []string
}

type exec struct {
  Script      []string
  Cwd         string
}

func (a *ContainerdAgent) createContainer(container container) (containerd.Container, error) {

    img, err := a.client.GetImage(a.ctx, imageRef+"/"+container.Image_ID)

    if err != nil {
      return nil, err
    }

    return a.client.NewContainer(
      a.ctx,
      container.ID,
      containerd.WithImage(img),
      containerd.WithNewSnapshot(container.ID+"-rootfs", img),
      containerd.WithNewSpec(
        oci.WithImageConfig(img),
        oci.WithImageConfigArgs(img, container.Cmd),
        oci.WithMounts(container.Mounts),
        oci.WithProcessCwd(container.WorkdingDir),
      ),
    )
}

func (a *ContainerdAgent) getContainer(id string) (containerd.Container, error) {
  containers, err := a.client.Containers(a.ctx)

  if err != nil {
    return nil, err
  }

  for _, c := range containers {
    if c.ID() == id{
      return c, err
    }
  }

  return nil, errors.New("Container not found")
}

func (a *ContainerdAgent) runContainer(id string) (error){

  //getting container c
  c, err := a.getContainer(id)

  if err != nil  {
    return err
  }

  //create new task
  task, err := c.NewTask(
    a.ctx,
    cio.NewCreator(cio.WithStdio),
  )

  if err != nil {
    return err
  }

  //starting task
  err = task.Start(a.ctx)

  return err

}

func (a* ContainerdAgent) execScript(id, exec_id string, e exec) (int, error){

  //getting container
  c, err := a.getContainer(id)

  if err != nil {
    return 1, err
  }

  //getting task
  task, err := c.Task(a.ctx, nil)

  if err != nil {
    return 1, err
  }

  //model process
  processmodel := &specs.Process{
    Terminal: false,
    Args: e.Script,
    Cwd: e.Cwd,
  }

  //execute process
  process, err := task.Exec(
    a.ctx,
    exec_id,
    processmodel,
    cio.NewCreator(),
  )

  if err != nil {
    return 1, err
  }

  process.Start(a.ctx)

  statusC, err := process.Wait(a.ctx)

  if err != nil {
    return 1, err
  }

  status := <- statusC

  code, exitedAt, err := status.Result()

  log.Printf("Exec %v executed with result %v at %v", exec_id, code, exitedAt)

  return int(code), err
}

func (a *ContainerdAgent) stopContainer(id string) (error) {
  container, err := a.getContainer(id)

  if err != nil {
    return err
  }

  task, err := container.Task(a.ctx, nil)

  if err != nil {
    return err
  }

  err = task.CloseIO(a.ctx)

  if err != nil {
    return err
  }

  status, err := task.Status(a.ctx)

  if status.Status == containerd.Running{
    exitStatus, err := task.Wait(a.ctx)

    if err != nil {
      return err
    }

    err = task.Kill(a.ctx, syscall.SIGTERM)

    if err != nil {
      return err
    }

    status := <-exitStatus

    code, _, err := status.Result()

    fmt.Printf("ExitStatus: %v\n", code)
  }

  task.Delete(a.ctx)

  return err
}

func (a *ContainerdAgent) deleteContainer(id string) (error) {

  container, err := a.getContainer(id)

  if err != nil {
    return err
  }

  return container.Delete(a.ctx, containerd.WithSnapshotCleanup)
}
