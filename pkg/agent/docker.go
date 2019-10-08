package agent

import(
  "fmt"
  "bytes"

  "github.com/cspengl/conmake/pkg/models"

  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types/container"
  "github.com/docker/docker/api/types"
  "golang.org/x/net/context"
)

type DockerAgent struct{
  endpoint    string
  apiversion  string
  client      *client.Client
}

const unixSocket = "unix:///var/run/docker.sock"

func NewDockerAgent(endpoint, apiversion string) (*DockerAgent, error) {

  if endpoint == "local"{
    endpoint = unixSocket
  }
  cli, err := client.NewClient(
    endpoint,
    apiversion,
    nil,
    nil)

  return &DockerAgent{
    endpoint: endpoint,
    apiversion: apiversion,
    client: cli,
  }, err
}

func (a *DockerAgent) Info() {
    fmt.Printf("Host: %v, Apiversion: %v\n", a.endpoint, a.apiversion)
}

func (a *DockerAgent) InitStation(project, stationname string, s models.Workstation)  error {

  ctx := context.Background()
  containerName := "temp-"+stationname
  imageName := project+"/"+stationname+":conmake"

  fmt.Printf("Initialize station %s as image %s from base %v\n", stationname, imageName, s.Base)

  cmd := s.Preparation

  if len(cmd) == 0{
    cmd =  []string{"/bin/bash"}
  }else {
    cmd = models.ScriptToCmd(s.Preparation)
  }

  resp, err := a.client.ContainerCreate(
    ctx,
    &container.Config{
      Image: s.Base,
      Cmd: cmd,
      Tty: true,
      OpenStdin: true},
    nil,
    nil,
    containerName)

  if err != nil {
    return err
  }

  fmt.Println("Container created")

   err = a.client.ContainerStart(
    ctx,
    resp.ID,
    types.ContainerStartOptions{})

  if err != nil{
    return err
  }

  fmt.Println("Container started")

   image, err := a.client.ContainerCommit(
    ctx,
    resp.ID,
    types.ContainerCommitOptions{
      Pause: true,
    })

  if err != nil {
    return err
  }

  fmt.Println("Container committed")

  err = a.client.ContainerRemove(
    ctx,
    resp.ID,
    types.ContainerRemoveOptions{
      Force: true})

  if err != nil {
    return err
  }

  fmt.Println("Container removed")

  err = a.client.ImageTag(
    ctx,
    image.ID,
    "testapp",
  )

  fmt.Println("Image tagged")

  return err
}

func (a *DockerAgent) PerformStep(project, stepname string, s models.Step) error {

  ctx := context.Background()
  containerName := project + "-" + s.Workstation

  resp, err := a.client.ContainerCreate(
    ctx,
    &container.Config{
      Image: s.Workstation,
      Cmd: models.ScriptToCmd(s.Script),
      Tty: true,
      OpenStdin: true,
    },
    &container.HostConfig{
      AutoRemove: true,
    },
    nil,
    containerName,
  )

  if err := a.client.ContainerStart(
    ctx,
    resp.ID,
    types.ContainerStartOptions{}); err != nil {
    return err
  }

  out, err := a.client.ContainerLogs(
    ctx,
    resp.ID,
    types.ContainerLogsOptions{
      ShowStderr: true,
      ShowStdout: true,
      Timestamps: true,
      Follow: true,
      Tail: "40",
    },
  )

  buf := new(bytes.Buffer)
  buf.ReadFrom(out)
  fmt.Printf(buf.String())

  return err
}
