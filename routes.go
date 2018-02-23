package main

import (
  "github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
  route := gin.Default()

  api := route.Group("/api") {
    api.GET("", func(c *gin.Context) {
      c.JSON(200, gin.H{"code": 200, "message": "api"})
    })
  }

  return route
}
