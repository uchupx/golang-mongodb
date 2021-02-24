package main

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoConn *mgo.Session

func createConnection() (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs: []string{
			"cluster0.mv5s7.gcp.mongodb.net"},
		Username: "test",
		Password: "chapzz33",
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	return mgo.DialWithInfo(&dialInfo)
}

type MyEntity struct {
	Data []byte `json:"data" bson:"data"`
}

func main() {
	var err error
	mongoConn, err = createConnection()
	if err != nil {
		panic(err)
	}
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", get)
	// router.POST("/somePost", posting)
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

func get(c *gin.Context) {
	session := mongoConn.Copy()
	defer session.Close()

	entity := MyEntity{}
	err := session.DB("test").C("data").Find(bson.M{}).One(&entity)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, entity)
}
