package v1

import (
  "net/http"
	"github.com/gin-gonic/gin"
)


func Version(c *gin.Context) {
  c.String(http.StatusOK, "1.0.0")
}
