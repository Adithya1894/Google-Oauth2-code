package main

import (
	"./logincontroller"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(logincontroller.RandToken(64)))

	router.Use(sessions.Sessions("goquestsession", store))

	router.LoadHTMLGlob("views/*")

	router.GET("/login", logincontroller.LoginHandler)
	router.GET("/auth", logincontroller.AuthHandler)

	
}
