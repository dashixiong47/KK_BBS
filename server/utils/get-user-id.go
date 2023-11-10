package utils

import "github.com/gin-gonic/gin"

func UserIDInt(c *gin.Context) int {
	value, exists := c.Get("id")
	if !exists {
		return 0
	}
	return int(value.(float64))
}
func UserIDUint(c *gin.Context) uint {
	value, exists := c.Get("id")
	if !exists {
		return 0
	}
	return uint(value.(float64))
}
