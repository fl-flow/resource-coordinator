package v1

import (
	"github.com/gin-gonic/gin"

  "github.com/fl-flow/resource-coordinator/http_server/v1/node"
	"github.com/fl-flow/resource-coordinator/http_server/v1/resource"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.GET("/version/", Version)
	node.RegisterRouter(Router.Group("node"))
	resource.RegisterRouter(Router.Group("resource"))
}
