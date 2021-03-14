package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/golang-mongodb/config"
	"github.com/uchupx/golang-mongodb/transport"
)

func main() {
	conf, err := config.ReadingConf()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// mongoConn
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// conn, err := config.ConnectionMongo(conf)
	// if err != nil {
	// 	panic(err)
	// }
	trans := transport.TransportHandler{}

	userHandler := trans.NewUserRequest(conf)

	router.GET("/user", userHandler.FindAll)
	router.POST("/user", userHandler.Insert)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run()
	router.Run(":3000")
}
