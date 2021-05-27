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
//jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"` //JobName赋值为job10

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
	//3.按照jobName字段过滤，想找出jobName=job10
	cond := &FindByJobName{"job10"}

	//5.查询
	var limit int64 = 2
	var skip int64 = 0
	var findoptions *options.FindOptions = &options.FindOptions{
		Limit:               &limit,
		Skip:                &skip,
	}
	cursor,err := collection.Find(context.TODO(),cond,findoptions)
	defer cursor.Close(context.TODO())
	if err != nil{
		fmt.Println(err)
		return
	}

	//遍历结果集
	for cursor.Next(context.TODO()){
		//定义一个日志对象
		record := &LogRecord{}

		//反序列化bson到对象
		if err := cursor.Decode(record); err != nil{
			fmt.Println(err)
			return
		}
		//把日志打印出来
		fmt.Println(*record)
	}

}
