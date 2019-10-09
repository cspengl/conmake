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
  endpoint    string
  apiversion  string
  client      *client.Client
}

const(
  localEndpoint = "local"
  unixSocket    = "unix:///var/run/docker.sock"
  imagePrefix   = "docker.io"
  conmakeTag    = "conmake"
  workspace     = "/workspace/"
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

func (a *DockerAgent) InitStation(config *StationConfig)  error {

  // Initialize context
  ctx := context.Background()

  // Define name of container for preparing the station
  containerName := "temp-"+config.StationName

  // Constructing image name based on project and stationname
  imageName := constructStationName(config)

  //Constructing preparation command based on config.Workstation.Preparation script
  script := config.Workstation.Preparation

  if len(cmd) == 0{
    script =  []string{"pwd"}
  }else {
    script = []string{genScript(config.Workstation.Preparation)}
  }

  //Check if there already is a workstation image
  present, err := a.imageExists(imageName)

  if err != nil{
    return err
  }

  //Found workstation -> Set config.Workstation.Base to imageName
  if present == true{
    fmt.Println("Found existing station")
    config.Workstation.Base = imageName
  }

  //Starting initialization
  fmt.Printf("Initialize station %s as image %s from base %v\n", config.StationName, imageName, config.Workstation.Base)

  id, err := a.spinUpStation(config, script, containerName)

  if err != nil {
    return err
  }

  //Remove old image
  _, err = a.client.ImageRemove(
    ctx,
    imageName,
    types.ImageRemoveOptions{},
  )

  if err == nil{
      fmt.Println("\tRemoved old station")
  }

  //Commit new image
  _, err = a.client.ContainerCommit(
    ctx,
    containerName,
    types.ContainerCommitOptions{
      Reference: imagePrefix+"/"+imageName,
    },
  )

  if err != nil {
    return err
  }

  fmt.Println("\tCommitted new station")

  //Stop and Remove container
  err = a.client.ContainerRemove(
    ctx,
    id,
    types.ContainerRemoveOptions{
      Force: true})

  if err != nil {
    return err
  }

  fmt.Println("\tRemoved station")

  fmt.Println("Station initialized")

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

func (a *DockerAgent) StationInitialized(config StationConfig) (bool, error) {

  tag := (config.ProjectName + "/" + config.StationName + ":conmake")

  return a.imageExists(tag)
}

func (a *DockerAgent) spinUpStation(config *StationConfig, cmd []string, containerName string) (string, error){

  ctx := context.Background()

  //Check if given base image exists
  present, err := a.imageExists(config.Workstation.Base)

  if err != nil{
    return "", err
  }

  //If base image does not exists it tries to download the base
  if present == false{
    a.downloadImage(config.Workstation.Base)
  }

  //Creating container
  resp, err := a.client.ContainerCreate(
    ctx,
    &container.Config{
      Image: config.Workstation.Base,
      User: "1000:1000",
      Shell: cmd,
      Tty: true},
    nil,
    nil,
    containerName)

  if err != nil {
    return "", err
  }

  fmt.Println("\tStation created")

  err = a.client.ContainerStart(
    ctx,
    resp.ID,
    types.ContainerStartOptions{})

  if err != nil{
    return "", err
  }

  fmt.Println("\tStation prepared")

  return resp.ID, err
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

func (a *DockerAgent) downloadImage(image string) error {

  ctx := context.Background()

  reader, err := a.client.ImagePull(
    ctx,
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

func genScript(script []string) string {
  res := ""
  for _, cmd := range script{
    res = res + cmd + " && "
  }

  return res[:len(res)-4]
}

func constructStationName(config *StationConfig) (string) {
  return config.ProjectName+"/"+config.StationName+":"+conmakeTag
}
