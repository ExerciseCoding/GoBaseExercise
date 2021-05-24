package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)
/**
使用Op方式取代
kv.Put
kv.Get
kv.Delete
 */
func main(){
	var(
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		putOp clientv3.Op
		opResp clientv3.OpResponse
		getOp clientv3.Op
		err error
	)

	//客户端配置
	config = clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	if client,err = clientv3.New(config); err != nil{
		fmt.Println(err)
		fmt.Println(client)
	}

	kv = clientv3.NewKV(client)

	//Op: operator
	putOp = clientv3.OpPut("/cron/jobs/job8","job8")

	//执行OP
	if opResp,err = kv.Do(context.TODO(),putOp);err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision:",opResp.Put().Header.Revision)

	//创建Op: getOp
	getOp = clientv3.OpGet("/cron/jobs/job8")
	//执行OP
	if opResp,err = kv.Do(context.TODO(),getOp); err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("数据Revision:",opResp.Get().Header.Revision)
	fmt.Println("数据value:",string(opResp.Get().Kvs[0].Value))


}