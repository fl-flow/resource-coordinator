package resource

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.POST("", ResourceRegisterView)
	Router.GET("/:resource_name/", ResourceDetailView)
	Router.POST("/alloc/", ResourceAlloc)
	Router.GET("/alloc/ws/", ResourceAllocWs)
}
