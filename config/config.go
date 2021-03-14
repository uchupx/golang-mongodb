package config

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Database struct {
		Driver string `json:"driver"`
		Mongo  struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Host     string `json:"host"`
			Database string `json:"database"`
		} `json:"mongo"`
		Mysql struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Host     string `json:"host"`
			Database string `json:"database"`
		} `json:"mysql"`
	} `json:"database"`
}

const driverMongo = "mongo"
const driverMysql = "mysql"

func ReadingConf() (*Config, error) {
	var conf Config

	viper.SetConfigName("config")                 // name of config file (without extension)
	viper.SetConfigType("json")                   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/golang-mongodb/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.golang-mongodb/") // call multiple times to add many search paths
	viper.AddConfigPath(".")                      // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	conf.Database.Driver = viper.GetString("database.driver")
	if conf.Database.Driver == driverMongo {
		conf.Database.Mongo.Username = viper.GetString("database.mongo.username")
		conf.Database.Mongo.Password = viper.GetString("database.mongo.password")
		conf.Database.Mongo.Host = viper.GetString("database.mongo.host")
		conf.Database.Mongo.Database = viper.GetString("database.mongo.database")
	} else if conf.Database.Driver == driverMysql {
		conf.Database.Mysql.Username = viper.GetString("database.mysql.username")
		conf.Database.Mysql.Password = viper.GetString("database.mysql.password")
		conf.Database.Mysql.Host = viper.GetString("database.mysql.host")
		conf.Database.Mysql.Database = viper.GetString("database.mysql.database")
	} else {
		panic(fmt.Errorf("Fatal error no database driver selected"))
	}

	return &conf, nil
}

func ConnectionMongo(conf *Config) (*mongo.Database, error) {
	ctx := context.TODO()
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", conf.Database.Mongo.Username, conf.Database.Mongo.Password, conf.Database.Mongo.Host)
	fmt.Println("connecting to", conf.Database.Mongo.Host)

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	database := client.Database(conf.Database.Mongo.Database)

	return database, nil
}
