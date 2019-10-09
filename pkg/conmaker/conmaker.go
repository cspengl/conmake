package conmaker

import(
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

func InitConmaker() (*Conmaker, error) {

  //Read file
  f, err := ioutil.ReadFile("Conmakefile.yaml")

  //Parse file and construct models
  c, err := models.NewConmakefile(f)

  //Construct agent
  a, err := agent.NewDockerAgent("local", "1.40")


  //Project path as cwd
  p, err := os.Getwd()

  return &Conmaker{
    agent: a,
    conmakefile: c,
    projectpath: p,
  }, err
}


func (c *Conmaker) Perform(step string) error{

  err := c.InitStation(c.conmakefile.Steps[step].Workstation)

  if err != nil {
    return err
  }

  pConfig := agent.PerformConfig{
    ProjectName: c.conmakefile.Project,
    ProjectPath: c.projectpath,
    StepName: step,
    Step: c.conmakefile.Steps[step],
  }

  return c.agent.PerformStep(pConfig)
}

func (c *Conmaker) InitStation(station string) error{

  config := agent.StationConfig{
    ProjectName: c.conmakefile.Project,
    StationName: station,
    Workstation: c.conmakefile.Workstations[station],
  }

  return c.agent.InitStation(&config)
}
