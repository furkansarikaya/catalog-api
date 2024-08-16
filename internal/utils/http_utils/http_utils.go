package http_utils

import (
	"github.com/furkansarikaya/catalog-api/internal/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

func RespondJson(c *gin.Context, statusCode int, body interface{}) {
	c.IndentedJSON(statusCode, body)
}

func RespondError(c *gin.Context, err rest_errors.RestErr) {
	RespondJson(c, err.Status(), err)
}
