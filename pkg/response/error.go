package response

import (
	"github.com/gin-gonic/gin"
	"goshop/pkg/config"
)

func Error(c *gin.Context, status int, err error, msg string) {
	cfg := config.GetConfig()
	errorResponse := map[string]interface{}{
		"message": msg,
	}

	if cfg.Environment != config.PRODUCTION_ENV {
		errorResponse["debug"] = err.Error()
	}

	c.JSON(status, Response{Error: errorResponse})
}
