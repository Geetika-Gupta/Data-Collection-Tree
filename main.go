package main

import (
	"github.com/project/user"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	user.InitializeRoutes(router)
	router.Run()
}
