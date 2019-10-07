package agent

import(
  "github.com/cspengl/conmake/pkg/models"
)

type Agent interface {
  PerformStep(string, string, models.Step) error
  InitStation(string, string, models.Workstation) error
  GetStations(interface{}) []interface{}
  Info()
}
