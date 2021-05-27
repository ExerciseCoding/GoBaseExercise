package main

import (
	"context"
	"fmt"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
func main(){
	//1.建立连接
	//设置客户端选项
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	duration := 10 * time.Second
	clientOptions.ConnectTimeout = &duration
	client ,err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil{
		fmt.Println(err)
		return
	}


	//2.选择数据库my_db,选择表my_collection

	collection := client.Database("my_db").Collection("my_collection")

	fmt.Println(collection)




}
