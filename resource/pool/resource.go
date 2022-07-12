package resourcepool

import (
  "sync"

  "github.com/fl-flow/resource-coordinator/common/error"
)


type NodeResourceType struct {
  Max               uint                                `json:"max"`  // TODO: []uint,  {height1, height2 ... low1, low2}
  Min               uint                                `json:"min"`  // TODO: []uint,  {height1, height2 ... low1, low2}
  Allocated         uint                                `json:"allocated"`
  Stream            *(map[string]uint)                   `json:"stream"`
  NodeRwMutex       *sync.RWMutex                       `json:"-"`
}


type ResourceType struct {
  NodeMap         map[string](*NodeResourceType)  `json:"node_map"`// resouceName -> NodeResourceType
  ResouceRwMutex  *sync.RWMutex                   `json:"-"`
}


type NodeResourceMapType map[string](*ResourceType)


var ResourceNodeMap NodeResourceMapType // node -> ResourceType
var ResourceNodeMapRwMutex *sync.RWMutex


func init() {
  ResourceNodeMap = make(NodeResourceMapType)
  ResourceNodeMapRwMutex = new(sync.RWMutex)
}


func (m NodeResourceMapType) GetResource(resource string) (*ResourceType, *error.Error) {
  n, ok := ResourceNodeMap[resource]
  if !ok {
    return nil, &error.Error{
      Code: 11000010,
      Hits: resource,
    }
  }
  return n, nil
}


func (r ResourceType) GetNodeResource(node string) (*NodeResourceType, *error.Error) {
  resource, rok := r.NodeMap[node]
  if !rok {
    return nil, &error.Error{
      Code: 11000020,
      Hits: node,
    }
  }
  return resource, nil
}
