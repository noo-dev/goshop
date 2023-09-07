package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Result: data,
	})
}
