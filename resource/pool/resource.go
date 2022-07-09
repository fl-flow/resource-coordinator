package resourcepool

import (
  "sync"

  "github.com/fl-flow/resource-coordinator/common/error"
)


type ResourceType struct {
  Max               uint          `json:"max"`  // TODO: []uint,  {height1, height2 ... low1, low2}
  Min               uint          `json:"min"`  // TODO: []uint,  {height1, height2 ... low1, low2}
  Allocated         uint          `json:"allocated"`
  ResouceRwMutex    *sync.RWMutex `json:"-"`
}


type NodeType struct {
  ResourceMap       map[string](*ResourceType )  `json:"resource_map"`// resouceName -> ResourceType
  NodeRwMutex       *sync.RWMutex             `json:"-"`
}


type NodeResourceMapType map[string](*NodeType)


var NodeResourceMap NodeResourceMapType // node -> NodeType
var NodeResourceMapRwMutex *sync.RWMutex


func init() {
  NodeResourceMap = make(NodeResourceMapType)
  NodeResourceMapRwMutex = new(sync.RWMutex)
}


func (m NodeResourceMapType) GetNode(node string) (*NodeType, *error.Error) {
  n, ok := NodeResourceMap[node]
  if !ok {
    return nil, &error.Error{
      Code: 11000010,
      Hits: node,
    }
  }
  return n, nil
}


func (n NodeType) GetResource(resouceName string) (*ResourceType, *error.Error) {
  resource, rok := n.ResourceMap[resouceName]
  if !rok {
    return nil, &error.Error{
      Code: 11000020,
      Hits: resouceName,
    }
  }
  return resource, nil
}
