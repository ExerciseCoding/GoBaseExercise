package main

import (
	"context"
	"fmt"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime int64 `bson:"endTime"`
}
//一条日志
type LogRecord struct {
	JobName string `bson:"jobName"`//任务名
	Command string `bson:"command"`//shell命令
	Err string `bson:"err"`//脚本错误
	Context string `bson:"context"`//脚本输出
	TimePoint TimePoint `bson:"timePoint"` //执行时间点
}

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

	collection := client.Database("cron").Collection("log")

	fmt.Println(collection)

	//3.插入记录(bson)
	record := &LogRecord{
		JobName: "job10",
		Command: "echo hello",
		Err:     "",
		Context: "hello",
		TimePoint:TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix()+10,
		},
	}

	result,err := collection.InsertOne(context.TODO(),record)
	if err != nil {
		fmt.Println(err)
		return
	}
	//_id: 默认生成一个全局唯一的ID,类型: ObjectID: 12字节的二进制
	fmt.Println(result.InsertedID)
	

}
