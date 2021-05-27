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

type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}
//{"timePoint.startTime":{"$lt":timestamp}}
type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
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

	//3.删除开始时间早于当前时间的所有日志($lt是less then)
	//delete({"timePoint.startTime":{"$lt":当前时间}})

	delcond := &DeleteCond{beforeCond:TimeBeforeCond{Before:time.Now().Unix()}}
	//执行删除
	delResult,err := collection.DeleteMany(context.TODO(),delcond)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数",delResult.DeletedCount)







}
