package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	HostRun := ":7600"

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	go TimeoutCheck()
	router.POST("/", SetPresence)

	router.Run(HostRun)
}
