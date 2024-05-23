package util

import "github.com/gin-gonic/gin"

func GetUrlQueryParams[T any](c *gin.Context) T {
	var payload T
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		return payload
	}
	return payload
}
