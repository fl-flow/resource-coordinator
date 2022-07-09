package node

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.POST("", NodeRegisterView)
}
