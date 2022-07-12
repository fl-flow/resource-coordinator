package resourcenode

import (
  "fmt"
  "sync"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/resource-coordinator/common/error"
  "github.com/fl-flow/resource-coordinator/resource/pool"
  "github.com/fl-flow/resource-coordinator/http_server/http/mixin"
)


func ResourceNodeRegisterView(context *gin.Context) {
  var f ResourceNodeRegisterSerializer
	if ok := mixin.CheckJSON(context, &f); !ok {
		return
	}
  if f.Init > f.Max || f.Init < f.Min || f.Max <= f.Min {
    mixin.CommonResponse(context, "error", &error.Error{
      Code: 11000011,
      Hits: "",
    })
    return
  }
  r, e := resourcepool.ResourceNodeMap.GetResource(f.ResourceName)
  if e != nil {
    mixin.CommonResponse(context, "error", e)
    return
  }
  r.ResouceRwMutex.Lock()
  defer r.ResouceRwMutex.Unlock()
  if _, rok := r.NodeMap[f.Node]; !rok {
    r.NodeMap[f.Node] = &resourcepool.NodeResourceType {
      Max: f.Max,
      Min: f.Min,
      Allocated: f.Init,
      NodeRwMutex: new(sync.RWMutex),
      Stream: &(map[string]uint{}),
    }
  }
  mixin.CommonResponse(context, r.NodeMap[f.Node], nil)
  return
}


func ResourceNodeUpView(context *gin.Context) {
  var f ResourceNodeUpSerializer
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  resource, e := ResourceNodeUpController(f)
  fmt.Println(resource.Allocated, resource.Stream)
  mixin.CommonResponse(context, resource, e)
}


func ResourceNodeDownView(context *gin.Context) {
  var f ResourceNodeDownSerializer
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  resource, e := ResourceNodeDownController(f)
  fmt.Println(resource.Allocated, resource.Stream)
  mixin.CommonResponse(context, resource, e)
}
