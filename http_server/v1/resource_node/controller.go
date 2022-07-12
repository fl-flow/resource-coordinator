package resourcenode

import (
  "fmt"

  "github.com/fl-flow/resource-coordinator/common/error"
  "github.com/fl-flow/resource-coordinator/resource/pool"
)


func ResourceNodeUpController(f ResourceNodeUpSerializer) (*resourcepool.NodeResourceType, *error.Error) {
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

  newAllocatedValue := resource.Allocated + f.Value
  if newAllocatedValue > resource.Max {
    return nil, &error.Error{
      Code: 11000030,
      Hits: fmt.Sprintf("where is only %v free resource", resource.Max - resource.Allocated),
    }
  }
  _, ok := (*((*resource).Stream))[f.Uid]
  if ok {
    return nil, &error.Error{
      Code: 11000030,
      Hits: fmt.Sprintf("uid %v is already exsited", f.Uid),
    }
  }
  (*((*resource).Stream))[f.Uid] = f.Value
  (*resource).Allocated = newAllocatedValue
  return resource, nil
}


func ResourceNodeDownController(f ResourceNodeDownSerializer) (*resourcepool.NodeResourceType, *error.Error) {
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

  _, ok := (*(*resource).Stream)[f.Uid]
  if !ok {
    return nil, &error.Error{
      Code: 11000030,
      Hits: fmt.Sprintf("uid %v is not exsited", f.Uid),
    }
  }
  newAllocatedValue := resource.Allocated - (*((*resource).Stream))[f.Uid]
  if newAllocatedValue < resource.Min {
    return nil, &error.Error{
      Code: 11000030,
      Hits: fmt.Sprintf("where is only %v resource allocated", resource.Allocated - resource.Min),
    }
  }
  delete(*((*resource).Stream), f.Uid)
  (*resource).Allocated = newAllocatedValue
  return resource, nil
}
