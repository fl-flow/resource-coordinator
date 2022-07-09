package node

import (
  "sync"
  "github.com/gin-gonic/gin"

  "github.com/fl-flow/resource-coordinator/resource/pool"
  "github.com/fl-flow/resource-coordinator/http_server/http/mixin"
)


func NodeRegisterView(context *gin.Context) {
  var f NodeRegisterSerializer
	if ok := mixin.CheckJSON(context, &f); !ok {
		return
	}
  resourcepool.NodeResourceMapRwMutex.Lock()
  defer resourcepool.NodeResourceMapRwMutex.Unlock()
  _, ok := resourcepool.NodeResourceMap[f.Node]
  if !ok {
    resourcepool.NodeResourceMap[f.Node] = &resourcepool.NodeType{
      ResourceMap: make(map[string](*resourcepool.ResourceType)),
      NodeRwMutex: new(sync.RWMutex),
    }
  }
  mixin.CommonResponse(context, resourcepool.NodeResourceMap[f.Node], nil)
  return
}
