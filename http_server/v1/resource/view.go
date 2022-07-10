package resource

import (
  "sync"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/resource-coordinator/common/error"
  "github.com/fl-flow/resource-coordinator/resource/pool"
  "github.com/fl-flow/resource-coordinator/http_server/http/mixin"
  "github.com/fl-flow/resource-coordinator/http_server/http/response"
)


func ResourceRegisterView(context *gin.Context) {
  var f ResourceRegisterSerializer
	if ok := mixin.CheckJSON(context, &f); !ok {
		return
	}
  resourcepool.ResourceNodeMapRwMutex.Lock()
  defer resourcepool.ResourceNodeMapRwMutex.Unlock()
  _, ok := resourcepool.ResourceNodeMap[f.ResourceName]
  if !ok {
    resourcepool.ResourceNodeMap[f.ResourceName] = &resourcepool.ResourceType{
      NodeMap: make(map[string](*resourcepool.NodeResourceType)),
      ResouceRwMutex: new(sync.RWMutex),
    }
  }
  mixin.CommonResponse(context, resourcepool.ResourceNodeMap[f.ResourceName], nil)
}


func ResourceDetailView(context *gin.Context) {
  resourceName := context.Param("resource_name")
  if resourceName == "" {
    response.R(context, 100, "query resource_name is required", "query resource_name is required")
    return
  }
  d, ok1 := resourcepool.ResourceNodeMap[resourceName]
  if !ok1 {
    mixin.CommonResponse(context, "error", &error.Error{
      Code: 11000010,
      Hits: resourceName,
    })
    return
  }
  mixin.CommonResponse(context, d, nil)
}
