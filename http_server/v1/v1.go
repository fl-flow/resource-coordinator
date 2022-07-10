package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/fl-flow/resource-coordinator/http_server/v1/resource"
	"github.com/fl-flow/resource-coordinator/http_server/v1/resource_node"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.GET("/version/", Version)
	resource.RegisterRouter(Router.Group("resource"))
	resourcenode.RegisterRouter(Router.Group("resource-node"))
}
