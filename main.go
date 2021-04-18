package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/uchupx/golang-mongodb/config"
	"github.com/uchupx/golang-mongodb/transport"
)

func main() {
	conf, err := config.ReadingConf()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	trans := transport.TransportHandler{}

	userHandler := trans.NewUserRequest(conf)
	// historyHandler := trans.NewUserRequest(conf)

	http.HandleFunc("/user", userHandler.Users)
	// http.HandleFunc("/history", userHandler.Insert)
	// http.HandleFunc("/history", userHandler.Insert)

	http.ListenAndServe(":3000", nil)
}
