package client

import (
  "fmt"
  "log"
  "sync"
  "encoding/json"
  "golang.org/x/net/websocket"
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


func (r ResourceType) Allocating (
  value uint,
  node *string,
  uid *string,
  argWait *sync.WaitGroup,
  wait *sync.WaitGroup,
) {
  // TODO: schema
  ws, err := websocket.Dial(
    fmt.Sprintf(
      "ws://%v:%v/api/v1/resource/alloc/ws/",
      r.Client.IP,
      r.Client.Port,
    ),
    "",
    fmt.Sprintf(
      "%v://%v:%v",
      r.Client.Schema,
      r.Client.IP,
      r.Client.Port,
    ),
  )
  message := []byte(fmt.Sprintf(`{"resource_name":"%v", "value": %v}`, r.ResourceName, value))
  _, err = ws.Write(message)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Send: %s\n", message)
  var msg = make([]byte, 512)
  m, err := ws.Read(msg)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Receive: %s\n", msg[:m])
  // if msg[:m] != []byte("true") {
  //   log.Fatal("lalalalala")
  // }
  var ret map[string]string
  if err := json.Unmarshal(msg[:m], &ret); err != nil {
    log.Fatal("lalalalala111")
  }
  *node = ret["node"]
  *uid = ret["uid"]
  argWait.Done()
  wait.Wait()
  ws.Close()
}


func (n ResourceType) ResourceNodeDown(node string, uid string) bool {
  url := fmt.Sprintf(
    "%v://%v:%v/api/v1/resource-node/down/",
    n.Client.Schema,
    n.Client.IP,
    n.Client.Port,
  )
  _, e := fetch(
    "PATCH",
    url,
    []byte(
      fmt.Sprintf(
        `{"resource_name":"%v","node":"%v","uid":"%v"}`,
        n.ResourceName,
        node,
        uid,
      ),
    ),
  )
  if e != nil {
    return false
  }
  return true
}
