package resourcenode

import (
  "fmt"

  "github.com/fl-flow/resource-coordinator/common/error"
  "github.com/fl-flow/resource-coordinator/resource/pool"
)


func ResourceNodeChangeController(f ResourceNodeChangeSerializer) (*resourcepool.NodeResourceType, *error.Error) {
  r, e := resourcepool.ResourceNodeMap.GetResource(f.ResourceName)
  if e != nil {
    return nil, e
  }
  resource , e := r.GetNodeResource(f.Node)
  if e != nil {
    return nil, e
  }
  resource.NodeRwMutex.Lock()
  defer resource.NodeRwMutex.Unlock()
  var newAllocatedValue uint
  fmt.Println(resource.Allocated, "resource.Allocatedresource.Allocated")
  if f.Type {
    newAllocatedValue = resource.Allocated + f.Value
    if newAllocatedValue > resource.Max {
      return nil, &error.Error{
        Code: 11000030,
        Hits: fmt.Sprintf("where is only %v free resource", resource.Max - resource.Allocated),
      }
    }
  } else {
    newAllocatedValue = resource.Allocated - f.Value
    if newAllocatedValue < resource.Min {
      return nil, &error.Error{
        Code: 11000030,
        Hits: fmt.Sprintf("where is only %v resource allocated", resource.Allocated - resource.Min),
      }
    }
  }
  fmt.Println(resource, "resourceresource", resourcepool.ResourceNodeMap)
  (*resource).Allocated = newAllocatedValue
  fmt.Println(resource, "resourceresource", resourcepool.ResourceNodeMap)
  return resource, nil
}
