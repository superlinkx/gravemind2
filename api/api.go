package api

import "github.com/gin-gonic/gin"

// GravemindOptions gives OPTIONS response for entire API, including current version
func GravemindOptions(c *gin.Context) {}

// V1Options gives OPTIONS response for API V1
func V1Options(c *gin.Context) {}
