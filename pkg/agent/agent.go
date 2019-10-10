package agent

const(
  ConmakeTag = "conmake"
  Workspace = "/workspace/"
)

type StationConfig struct {
  ProjectName string
  StationName string
  Image       string
  Script      []string
  Workspace   string
}

type Agent interface {
  PerformStep(*StationConfig) error
  InitStation(*StationConfig, bool) (string, error)
  DeleteStation(*StationConfig) error
  StationExists(*StationConfig) (bool, error)
  Info()
}

func ConstructStationImageName(config *StationConfig) string{
  return ConstructStationImageNameFromRaw(config.ProjectName, config.StationName)
}

func ConstructStationImageNameFromRaw(projectname, stationname string) string {
  return projectname+"/"+stationname+":"+ConmakeTag
}

func ConstructStationContainerName(config *StationConfig) string{
  return config.ProjectName+"-"+config.StationName
}
