package db

import (
	"anonymous.net.cn/public-system-backend/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
	"time"
)

var MongoClient *mongo.Client

// Init 初始化客户端连接池
func Init() {
	host := config.Config.DB.Mongo.Host
	port := config.Config.DB.Mongo.Port
	db := config.Config.DB.Mongo.DB
	username := config.Config.DB.Mongo.Username
	password := config.Config.DB.Mongo.Password
	var uri string
	if username == "" {
		uri = "mongodb://" + host + ":" + strconv.Itoa(port) + "/" + db
	} else {
		uri = "mongodb://" + username + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + db
	}
	fmt.Println(uri)
	clientOptions := options.Client().ApplyURI(uri).SetMaxPoolSize(3000).SetMinPoolSize(4).SetMaxConnIdleTime(30 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln("数据库连接失败: ", err.Error())
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalln("数据库连接失败: ", err.Error())
	}
	MongoClient = client
}
