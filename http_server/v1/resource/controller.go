package resource

import (
  "fmt"

  "github.com/fl-flow/resource-coordinator/common/error"
  "github.com/fl-flow/resource-coordinator/resource/pool"
)


func ResourceChangeController(f ResourceChangeSerializer) (*resourcepool.ResourceType, *error.Error) {
  n, e := resourcepool.NodeResourceMap.GetNode(f.Node)
  if e != nil {
    return nil, e
  }
  resource , e := n.GetResource(f.ResourceName)
  if e != nil {
    return nil, e
  }
  resource.ResouceRwMutex.Lock()
  defer resource.ResouceRwMutex.Unlock()
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
  fmt.Println(resource, "resourceresource", resourcepool.NodeResourceMap)
  (*resource).Allocated = newAllocatedValue
  fmt.Println(resource, "resourceresource", resourcepool.NodeResourceMap)
  return resource, nil
}
