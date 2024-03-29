package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)
func main(){
	var(
		config clientv3.Config
		client *clientv3.Client
		putResp *clientv3.PutResponse
		err error
		kv clientv3.KV
	)

	//客户端配置
	config = clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	if client,err = clientv3.New(config); err != nil{
		fmt.Println(err)
	}

	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	//
	if putResp, err = kv.Put(context.TODO(),"/cron/jobs/job2","hello"); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(putResp.Header.Revision)
	}
}