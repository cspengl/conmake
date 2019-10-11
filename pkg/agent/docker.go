package agent

import(
  "fmt"
  "io"
  "bytes"

  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types/container"
  "github.com/docker/docker/api/types/mount"
  "github.com/docker/docker/api/types"

  "golang.org/x/net/context"

  "github.com/tj/go-spin"
)

type DockerAgent struct{
  ctx         context.Context
  endpoint    string
  apiversion  string
  client      *client.Client
}

const(
  localEndpoint = "local"
  unixSocket    = "unix:///var/run/docker.sock"
  imagePrefix   = "docker.io"
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
    ctx: context.Background(),
    endpoint: endpoint,
    apiversion: apiversion,
    client: cli,
  }, err
}

func (a *DockerAgent) Info() {
    fmt.Printf("Host: %v, Apiversion: %v\n", a.endpoint, a.apiversion)
}

func (a *DockerAgent) spinUpStation(config *StationConfig) (string, error) {

  //Check if image exists
  imageExists, err := a.imageExists(config.Image)

  if err != nil {
    return "", err
  }

  //If image not present agent tries to download it
  if !imageExists {
    err = a.downloadImage(config.Image)
  }

  //Command
  cmd := []string{
    "sh",
    "-c",
    genShellScript(config.Script),
  }

  //Creating container from image
  resp, err := a.client.ContainerCreate(
    a.ctx,
    &container.Config{
      User: "1000:1000",
      Image: config.Image,
      Tty: true,
      Cmd: cmd,
      WorkingDir: Workspace,
    },
    &container.HostConfig{
      Mounts: []mount.Mount{
        {
          Type: mount.TypeBind,
          Source: config.Workspace,
          Target: Workspace,
        },
      },
    },
    nil,
    config.StationName,
  )

  if err != nil {
    return "", err
  }

  //Start created container
   err = a.client.ContainerStart(
    a.ctx,
    resp.ID,
    types.ContainerStartOptions{},
  )

  return resp.ID, err
}

func (a *DockerAgent) InitStation(config *StationConfig, existing bool) (string, error) {

  //Creating image name
  imageName := ConstructStationImageName(config)

  //Spinup station
  stationContainerID, err := a.spinUpStation(config)

  if err != nil {
    return "", err
  }

  //Delete old station
  if existing {
    err = a.DeleteStation(config)
  }

  if err != nil {
    return "", err
  }

  //Committing station as new image
  _, err = a.client.ContainerCommit(
    a.ctx,
    config.StationName,
    types.ContainerCommitOptions{
      Reference: imagePrefix+"/"+imageName,
    },
  )

  if err != nil {
    return "", err
  }

  //Remove running station container
  err = a.client.ContainerRemove(
    a.ctx,
    stationContainerID,
    types.ContainerRemoveOptions{
      Force: true,
    },
  )

  return imageName, err
}

func (a *DockerAgent) DeleteStation(config *StationConfig) error {
  _, err := a.client.ImageRemove(
    a.ctx,
    config.Image,
    types.ImageRemoveOptions{
      Force: true,
      PruneChildren: true,
    },
  )

  return err
}

func (a *DockerAgent) PerformStep(config *StationConfig) error {

  //Spin up station
  stationContainerID, err := a.spinUpStation(config)

  if err != nil {
    return err
  }

  //Attaching to container logs
  out, err := a.client.ContainerLogs(
    a.ctx,
    stationContainerID,
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

  //Remove running station container after shell script runned
  err = a.client.ContainerRemove(
    a.ctx,
    stationContainerID,
    types.ContainerRemoveOptions{
      Force: true,
    },
  )

  return err
}

func (a *DockerAgent) StationExists(config *StationConfig) (bool, error) {
  return a.imageExists(ConstructStationImageName(config))
}


func (a *DockerAgent) imageExists(imageTag string)(bool, error){

  images, err := a.client.ImageList(
    a.ctx,
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

func (a *DockerAgent) downloadImage(image string) error {

  reader, err := a.client.ImagePull(
    a.ctx,
    imagePrefix+"/library/" + image,
    types.ImagePullOptions{},
  )

  buf := make([]byte, 32*2048)
  spinner := spin.New()
  for{
    _, er := reader.Read(buf)
    if er != nil {
      if er == io.EOF{
        break
      }
    }

    fmt.Printf("\rDownloading image %s ", spinner.Next())
  }

  fmt.Println("")
  defer reader.Close()

  return err
}
