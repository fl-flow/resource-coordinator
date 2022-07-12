package resourcenode

import (
	"github.com/gin-gonic/gin"
)


func RegisterRouter(Router *gin.RouterGroup)  {
  Router.POST("", ResourceNodeRegisterView)
	Router.PATCH("/up/", ResourceNodeUpView)
	Router.PATCH("/down/", ResourceNodeDownView)
}
