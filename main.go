package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/uchupx/golang-mongodb/config"
	"github.com/uchupx/golang-mongodb/middleware"
	"github.com/uchupx/golang-mongodb/transport"
)

func main() {
	conf, err := config.ReadingConf()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router := http.NewServeMux()

	trans := transport.TransportHandler{}

	userHandler := trans.NewUserRequest(conf)
	// historyHandler := trans.NewUserRequest(conf)

	router.HandleFunc("/user", userHandler.Users)
	// http.HandleFunc("/history", userHandler.Insert)
	// http.HandleFunc("/history", userHandler.Insert)

	var logger log.Logger

	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// logFile, err := os.Open("./logging.log")
	// if err != nil {
	// 	panic(err)
	// }

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	// Create an instance of our LoggingMiddleware with our configured logger
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	loggedRouter := loggingMiddleware(router)

	if err := http.ListenAndServe(":3000", loggedRouter); err != nil {
		logger.Log("status", "fatal", "err", err)
		os.Exit(1)
	}
}
