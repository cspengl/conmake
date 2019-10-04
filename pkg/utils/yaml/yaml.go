package yaml

import(
  "io/ioutil"

  "gopkg.in/yaml.v2"
)

const(
  filePermissions = 0644
)

//Generic Value
type YAML interface{
  Data() interface{}
  Get(interface{}) interface{}
  Set(interface{}, interface{})
}

//Value for YAML list
type List struct {
    data []interface{}
}

func NewList(data []interface{}) *List {
  return &List{data: data}
}

func (yl *List) Data() interface{}{
  return yl.data
}

func (yl *List) Set(index, data interface{}) {
    yl.data[index.(int)] = data
}

func (yl *List) Get(index interface{}) interface{}{
    return yl.data[index.(int)]
}

func (yl *List) Len() int {
  return len(yl.data)
}

func (yl *List) Append(obj interface{}){
  yl.data = append(yl.data, obj)
}

//Value for mapping
type Map struct{
  data  map[string]interface{}
}

func NewMap(data map[string]interface{}) *Map {
  return &Map{data: data}
}

func (ym *Map) Data() interface{}{
  return ym.data
}

func (ym *Map) Set(key, data interface{}){
  ym.data[key.(string)] = data
}

func (ym *Map) Get(key interface{}) interface{}{
  return ym.data[key.(string)]
}

func Load(data []byte) YAML {
  raw := make(map[interface{}]interface{})
  yaml.Unmarshal(data, &raw)
  return generate(raw).(YAML)
}

func FromFile(path string) YAML {
  file, _ := ioutil.ReadFile(path)
  return Load(file)
}

func Dump(yamlObj YAML) ([]byte, error) {
  return yaml.Marshal(Raw(yamlObj))
}

func ToFile(yamlObj YAML, path string){
  raw, _ := Dump(yamlObj)
  ioutil.WriteFile(path, raw, filePermissions)
}

func Raw(yamlObj interface{}) interface{} {
  switch yamlObj.(type){
  default:
    return yamlObj
  case *List:
    ymlList := yamlObj.(*List)
    rawList := make([]interface{}, ymlList.Len())
    for i , y := range ymlList.Data().([]interface{}){
      rawList[i] = Raw(y)
    }
    return rawList
  case *Map:
    ymlMap := yamlObj.(*Map)
    rawMap := make(map[interface{}]interface{})
    for key, y := range ymlMap.Data().(map[string]interface{}){
       rawMap[key] = Raw(y)
    }
    return rawMap
  }
}

func generate(raw interface{}) interface{}{
  switch raw.(type){
  default:
    //Generate scalar
    return raw
  case []interface{}:
    //Generate list
    rawList := raw.([]interface{})
    ymlList := make([]interface{}, len(raw.([]interface{})))
    for i , element := range rawList {
      ymlList[i] = generate(element)
    }
    return &List{data: ymlList}
  case map[interface{}]interface{}:
    //Generate map
    rawMap := raw.(map[interface{}]interface{})
    ymlMap := make(map[string]interface{})
    for key , element := range rawMap{
      ymlMap[key.(string)] = generate(element)
    }
    return &Map{data: ymlMap}
  }
}
