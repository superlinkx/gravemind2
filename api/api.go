package api

import "github.com/gin-gonic/gin"

type action struct {
	Params      string `json:"parameters"`
	Description string `json:"description"`
}

func setAllowHeader(allow string, c *gin.Context) {
	c.Writer.Header().Set("Allow", allow)
}

// GravemindOptions gives OPTIONS response for entire API, including current version
func GravemindOptions(c *gin.Context) {
	setAllowHeader("OPTIONS", c)

	paths := map[string]string{
		"v1": "/gravemind/v1",
	}

	c.JSON(200, gin.H{
		"message": "Welcome to the gravemind API. Current Version is at /gravemind/v1",
		"paths":   paths,
	})
}

// V1Options gives OPTIONS response for API V1
func V1Options(c *gin.Context) {
	setAllowHeader("OPTIONS", c)

	paths := map[string]string{
		"dashboard": "/gravemind/v1/dashboard",
	}

	c.JSON(200, gin.H{
		"message": "Welcome to V1 of the gravemind API. This is the most current version.",
		"paths":   paths,
	})
}
