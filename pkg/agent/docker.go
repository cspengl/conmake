package agent

import(
  "fmt"

  "github.com/docker/docker/client"
)

type DockerAgent struct{
  endpoint    string
  apiversion  string
  client      *client.Client
}

const unixSocket = "unix:///var/run/docker.sock"

func NewDockerAgent(endpoint, apiversion string) DockerAgent, err {

  if endpoint == "local"{
    endpoint := unixSocket
  }
  cli, err := client.NewClient(
    endpoint,
    apiversion,
    nil,
    nil)

  return DockerAgent{
    endpoint: endpoint,
    apiversion: apiversion,
    client: cli,
  }, err
}

func (*a DockerAgent) Info() {
    fmt.Printf("Host: %v, Apiversion: %v", a.endpoint, a.apiversion)
}
