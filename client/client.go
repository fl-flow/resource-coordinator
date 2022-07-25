package client

import (
  "fmt"
)


type Client struct {
  Schema  string
  IP      string
  Port    int
}


type ResourceType struct {
  Client        Client
  ResourceName  string
}


type ResourceNodeType struct {
  Resource    ResourceType
  Node        string
  Max         uint64
  Min         uint64
  Init        uint64
}


func (c Client) NewResource (resourceName string) (*ResourceType, bool) {
  url := fmt.Sprintf("%v://%v:%v/api/v1/resource/", c.Schema, c.IP, c.Port)
  _, e := fetch("POST", url, []byte(fmt.Sprintf(`{"resource_name":"%v"}`, resourceName)))
  if e != nil {
    return nil, false
  }
  return &ResourceType{
    Client: c,
    ResourceName: resourceName,
  }, true
}


func (r ResourceType) NewResourceNode (
  node string,
  max uint64,
  min uint64,
  init uint64,
) (*ResourceNodeType, bool) {
  url := fmt.Sprintf("%v://%v:%v/api/v1/resource-node/", r.Client.Schema, r.Client.IP, r.Client.Port)
  _, e := fetch(
    "POST",
    url,
    []byte(
      fmt.Sprintf(
        `{"resource_name":"%v","node":"%v","max":%v,"min":%v,"init":%v}`,
        r.ResourceName,
        node,
        max,
        min,
        init,
      ),
    ),
  )
  if e != nil {
    return nil, false
  }
  return &ResourceNodeType{
    Resource: r,
    Node: node,
    Max: max,
    Min: min,
    Init: init,
  }, true
}



func (n ResourceNodeType) ResourceNodeUp(diff uint64, uid string) bool {
  url := fmt.Sprintf(
    "%v://%v:%v/api/v1/resource-node/up/",
    n.Resource.Client.Schema,
    n.Resource.Client.IP,
    n.Resource.Client.Port,
  )
  _, e := fetch(
    "PATCH",
    url,
    []byte(
      fmt.Sprintf(
        `{"resource_name":"%v","node":"%v","value":%v,"uid":"%v"}`,
        n.Resource.ResourceName,
        n.Node,
        diff,
        uid,
      ),
    ),
  )
  if e != nil {
    return false
  }
  return true
}


func (n ResourceNodeType) ResourceNodeDown(uid string) bool {
  url := fmt.Sprintf(
    "%v://%v:%v/api/v1/resource-node/down/",
    n.Resource.Client.Schema,
    n.Resource.Client.IP,
    n.Resource.Client.Port,
  )
  _, e := fetch(
    "PATCH",
    url,
    []byte(
      fmt.Sprintf(
        `{"resource_name":"%v","node":"%v","uid":"%v"}`,
        n.Resource.ResourceName,
        n.Node,
        uid,
      ),
    ),
  )
  if e != nil {
    return false
  }
  return true
}


func (r ResourceType) Alloc (value uint) (string, string, bool) {
  url := fmt.Sprintf(
    "%v://%v:%v/api/v1/resource/alloc/",
    r.Client.Schema,
    r.Client.IP,
    r.Client.Port,
  )
  ret, e := fetch(
    "POST",
    url,
    []byte(fmt.Sprintf(`{"resource_name":"%v", "value": %v}`, r.ResourceName, value)),
  )
  if e != nil {
    return "", "", false
  }
  data := ret.Data.(map[string]interface{})
  return data["node"].(string), data["uid"].(string), true
}
