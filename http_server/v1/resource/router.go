package resource

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.POST("", ResourceRegisterView)
	Router.PATCH("", ResourceChangeView)
}
