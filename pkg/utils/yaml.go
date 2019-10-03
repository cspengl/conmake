package utils

import(
  "io/ioutil"

  "gopkg.in/yaml.v2"
)

//Generic Value
type YAML interface{
  Data() interface{}
  Get(interface{}) YAML
}

//Value for YAML object
type YAMLScalar struct {
    data interface{}
}

func (ys YAMLScalar) Data() interface{}{
  return ys.data
}

//Should not be used but is needed to implement Scalar as native YAML value
func (ys YAMLScalar) Get(not_used interface{}) YAML{
  return ys
}

//Value for YAML list
type YAMLList struct {
    data []YAML
}

func (yl YAMLList) Data() interface{}{
  return yl.data
}

func (yl YAMLList) Get(abstractIndex interface{}) YAML{
    return yl.data[abstractIndex.(int)]
}

//Value for mapping
type YAMLMap struct{
  data  map[string]YAML
}

func (ym YAMLMap) Data() interface{}{
  return ym.data
}

func (ym YAMLMap) Get(key interface{}) YAML{
  return ym.data[key.(string)]
}

func Parse(data []byte) YAML {
  raw := make(map[interface{}]interface{})
  yaml.Unmarshal(data, &raw)
  return generate(raw)
}

func FromFile(path string) YAML {
  file, _ := ioutil.ReadFile(path)
  return Parse(file)
}

func generate(raw interface{}) YAML{
  switch raw.(type){
  default:
    //Generate scalar
    return YAMLScalar{
      data: raw,
    }
  case []interface{}:
    //Generate list
    list := []YAML{}
    for _, element := range raw.([]interface{}) {
      list = append(list, generate(element))
    }
    return YAMLList{
      data: list,
    }
  case map[interface{}]interface{}:
    //Generate map
    yamlMap := make(map[string]YAML)
    for key , element := range raw.(map[interface{}]interface{}){
      yamlMap[key.(string)] = generate(element)
    }
    return YAMLMap{
      data: yamlMap,
    }
  }
}
