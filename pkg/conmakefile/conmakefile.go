package conmakefile

import(
  "gopkg.in/yaml.v2"
)

const(
  keyVersion  = "version"
  keyProject  = "project"
  keyAgent    = "agent"
  keySteps    = "steps"
  keyStations = "stations"
)

type Conmakefile struct {
  Version string
  Project string
  Agent map[interface{}]interface{}
  Steps []interface{}
}

func Parse(data []byte) Conmakefile {
  toplevel := make(map[string]interface{})
  yaml.Unmarshal(data, &toplevel)

  return Conmakefile {
    Version: toplevel[keyVersion].(string),
    Project: toplevel[keyProject].(string),
    Agent: toplevel[keyAgent].(map[interface{}]interface{}),
    Steps: toplevel[keySteps].([]interface{}),
  }
}
