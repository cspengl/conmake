package conmaker

import(
  "io/ioutil"

  "github.com/cspengl/conmake/pkg/agent"
  "github.com/cspengl/conmake/pkg/models"
)

type Conmaker struct{
    agent agent.Agent
    conmakefile *models.Conmakefile
}

func NewConmaker(a agent.Agent, c *models.Conmakefile) *Conmaker {
  return &Conmaker{
    agent: a,
    conmakefile: c,
  }
}

func InitConmaker() (*Conmaker, error) {

  //Read file
  f, err := ioutil.ReadFile("Conmakefile.yaml")

  //Parse file and construct models
  c, err := models.NewConmakefile(f)

  //Construct agent
  a, err := agent.NewDockerAgent("local", "1.40")

  return &Conmaker{
    agent: a,
    conmakefile: c,
  }, err
}


func (c *Conmaker) Perform(step string) error{

  workstation = c.conmakefile[step].Workstation

  // err := c.agent.InitStation(
  //   c.conmakefile.Project,
  //   workstation,
  //   c.conmakefile.Workstations[workstation],
  // )
  //
  // if err != nil{
  //   return err
  // }

  err := c.agent.PerformStep(
    c.conmakefile.Project,
    step,
    c.conmakefile.Steps[step],
  )

  return err
}

func (c *Conmaker) InitStation(station string) error{
  return c.agent.InitStation(
    c.conmakefile.Project,
    station,
    c.conmakefile.Workstations[station],
  )
}
