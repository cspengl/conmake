package conmaker

import(
  "log"
  "os"
  "io/ioutil"

  "github.com/cspengl/conmake/pkg/agent"
  "github.com/cspengl/conmake/pkg/models"
)

type Conmaker struct{
    agent agent.Agent
    conmakefile *models.Conmakefile
    projectpath string
}

func NewConmaker(a agent.Agent, c *models.Conmakefile, p string) *Conmaker {
  return &Conmaker{
    agent: a,
    conmakefile: c,
    projectpath: p,
  }
}

func InitConmaker(projectpath, conmakefile string) (*Conmaker, error) {

  //Read file
  f, err := ioutil.ReadFile(conmakefile)

  if err != nil {
    log.Fatal("Conmakefile not found")
  }

  //Parse file and construct models
  c, err := models.NewConmakefile(f)

  //Construct agent
  a, err := agent.NewDockerAgent("local", "1.40")

  if projectpath == "./"{
    projectpath, err = os.Getwd()
  }

  return &Conmaker{
    agent: a,
    conmakefile: c,
    projectpath: projectpath,
  }, err
}


func (c *Conmaker) Perform(step string) error{

  imageName, err := c.InitStation(c.conmakefile.Steps[step].Workstation)

  if err != nil {
    return err
  }

  config := &agent.StationConfig{
    ProjectName: c.conmakefile.Project,
    StationName: step,
    Image: imageName,
    Script: c.conmakefile.Steps[step].Script,
    Workspace: c.projectpath,
  }

  return c.agent.PerformStep(config)
}

func (c *Conmaker) InitStation(station string) (string, error){

  config := &agent.StationConfig{
    ProjectName: c.conmakefile.Project,
    StationName: station,
    Image: agent.ConstructStationImageNameFromRaw(
      c.conmakefile.Project,
      station,
    ),
    Script: c.conmakefile.Workstations[station].Script,
    Workspace: c.projectpath,
  }

  stationExists, err := c.agent.StationExists(config)

  if err != nil{
    return "", err
  }

  if !stationExists {
    log.Println("Station not found, initializing from base...")
    config.Image = c.conmakefile.Workstations[station].Base
  }

  return c.agent.InitStation(config, stationExists)
}

func (c *Conmaker) DeleteStation(station string) error {
  config := &agent.StationConfig{
    ProjectName: c.conmakefile.Project,
    StationName: station,
    Image: agent.ConstructStationImageNameFromRaw(
      c.conmakefile.Project,
      station,
    ),
    Script: c.conmakefile.Workstations[station].Script,
    Workspace: c.projectpath,
  }

  return c.agent.DeleteStation(config)
}

func (c* Conmaker) CleanStation(station string) error {
  config := &agent.StationConfig{
    ProjectName: c.conmakefile.Project,
    StationName: station,
    Image: agent.ConstructStationImageNameFromRaw(
      c.conmakefile.Project,
      station,
    ),
    Script: c.conmakefile.Workstations[station].Script,
    Workspace: c.projectpath,
  }

  err := c.agent.DeleteStation(config)

  if err != nil {
    return err
  }

  config.Image = c.conmakefile.Workstations[station].Base

  _, err = c.InitStation(station)

  return err
}
