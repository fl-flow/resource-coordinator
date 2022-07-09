package resource

import (
  "sync"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/resource-coordinator/common/error"
  "github.com/fl-flow/resource-coordinator/resource/pool"
  "github.com/fl-flow/resource-coordinator/http_server/http/mixin"
)


func ResourceRegisterView(context *gin.Context) {
  var f ResourceRegisterSerializer
	if ok := mixin.CheckJSON(context, &f); !ok {
		return
	}
  // n, e := getNode(f.Node)
  n, e := resourcepool.NodeResourceMap.GetNode(f.Node)
  if e != nil {
    mixin.CommonResponse(context, "error", e)
    return
  }
  if f.Init > f.Max || f.Init < f.Min || f.Max <= f.Min {
    mixin.CommonResponse(context, "error", &error.Error{
      Code: 11000011,
      Hits: "",
    })
    return
  }
  n.NodeRwMutex.Lock()
  defer n.NodeRwMutex.Unlock()
  if _, rok := n.ResourceMap[f.ResourceName]; !rok {
    n.ResourceMap[f.ResourceName] = &resourcepool.ResourceType {
      Max: f.Max,
      Min: f.Min,
      Allocated: f.Init,
      ResouceRwMutex: new(sync.RWMutex),
    }
  }
  mixin.CommonResponse(context, n.ResourceMap[f.ResourceName], nil)
  return
}


func ResourceChangeView(context *gin.Context) {
  var f ResourceChangeSerializer
  if ok := mixin.CheckJSON(context, &f); !ok {
    return
  }
  resource, e := ResourceChangeController(f)
  mixin.CommonResponse(context, resource, e)
}
