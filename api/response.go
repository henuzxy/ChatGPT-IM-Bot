package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "SUCCESS",
	})
}
func ResponseFailure(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 101,
		"msg":  "server error",
	})
}
