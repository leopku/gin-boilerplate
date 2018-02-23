package main

import (
  "time"

  jwt "github.com/dgrijalva/jwt-go"
  jwtMiddleware "github.com/gin-gonic/contrib/jwt"
  "github.com/gin-gonic/gin"
)

var (
  route gin.Engine
)

func init() {
  route = gin.Default()

  public := route.Group("/api")

  public.GET("/", func(c *gin.Context) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    token.Claims = jwt.MapClaims{
      "userId": "1234",
      "exp": time.Now().Add(time.Hour * 1).Unix(),
    }

    tokenStr, err := token.SignedString([]byte("my secret"))
    if err != nil {
      c.JSON(500, gin.H{"code": 500, "message": "could not generate token"})
    }
    c.JSON(200, gin.H{"data": gin.H{"token": tokenStr}})
  })
}
