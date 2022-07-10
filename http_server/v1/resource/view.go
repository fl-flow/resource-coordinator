package resource

import (
  "sync"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/resource-coordinator/resource/pool"
  "github.com/fl-flow/resource-coordinator/http_server/http/mixin"
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
  return
}
