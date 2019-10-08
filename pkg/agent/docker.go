package agent

import(
  "fmt"
  "bytes"

  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types/container"
  "github.com/docker/docker/api/types/mount"
  "github.com/docker/docker/api/types"
  "golang.org/x/net/context"
)

type DockerAgent struct{
  endpoint    string
  apiversion  string
  client      *client.Client
}

const(
  unixSocket  = "unix:///var/run/docker.sock"
  imagePrefix = "docker.io"
  workspace = "/workspace/"
)


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

func (a *DockerAgent) InitStation(config StationInitConfig)  error {

  ctx := context.Background()
  containerName := "temp-"+config.StationName
  imageName := imagePrefix+"/"+config.ProjectName+"/"+config.StationName+":conmake"

  fmt.Printf("Initialize station %s as image %s from base %v\n", config.StationName, imageName, config.Workstation.Base)

  cmd := config.Workstation.Preparation

  if len(cmd) == 0{
    cmd =  []string{"pwd"}
  }else {
    cmd = []string{
      "sh",
      "-c",
      genScript(config.Workstation.Preparation),
    }
  }

  idri, err := a.client.ImageRemove(
    ctx,
    imageName,
    types.ImageRemoveOptions{},
  )

  if err == nil {
    fmt.Printf("Found existing station and deleted image [%v]\n", idri[0].Untagged)
  }

  if present, _ := a.imageExists(config.Workstation.Base); present == false{
    fmt.Println("Try pulling image")
    r, _ := a.client.ImagePull(
      ctx,
      "docker.io/library/" + config.Workstation.Base,
      types.ImagePullOptions{},
    )

    defer r.Close()
  }

  resp, err := a.client.ContainerCreate(
    ctx,
    &container.Config{
      Image: config.Workstation.Base,
      User: "1000:1000",
      Cmd: cmd,
      Tty: true},
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

  _, err = a.client.ContainerCommit(
    ctx,
    resp.ID,
    types.ContainerCommitOptions{
      Reference: imageName,
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

  return err
}

func (a *DockerAgent) PerformStep(config PerformConfig) error {

  ctx := context.Background()
  containerName := config.ProjectName + "-" + config.Step.Workstation
  imageName := imagePrefix+"/"+config.ProjectName+"/"+config.Step.Workstation+":conmake"

  cmd := []string{
    "sh",
    "-c",
    genScript(append([]string{"cd " + workspace}, config.Step.Script...)),
  }

  resp, err := a.client.ContainerCreate(
    ctx,
    &container.Config{
      Image: imageName,
      User: "1000:1000",
      Cmd: cmd,
      Tty: true,
      OpenStdin: true,
    },
    &container.HostConfig{
      AutoRemove: true,
      Mounts: []mount.Mount{
        {
          Type: mount.TypeBind,
          Source: config.ProjectPath,
          Target: workspace,
        },
      },
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

func (a *DockerAgent) StationInitialized(config StationInitConfig) (bool, error) {

  tag := (config.ProjectName + "/" + config.StationName + ":conmake")

  return a.imageExists(tag)
}

func (a *DockerAgent) imageExists(imageTag string)(bool, error){

  ctx := context.Background()

  images, err := a.client.ImageList(
    ctx,
    types.ImageListOptions{
      All: true,
    },
  )

  if err != nil {
    return false, err
  }

  for _, img := range images{
    for _, tag := range img.RepoTags{
      if tag == imageTag{
        return true, err
      }
    }
  }

  return false, err
}

func genScript(script []string) string {
  res := ""
  for _, cmd := range script{
    res = res + cmd + " && "
  }

  return res[:len(res)-4]
}
