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
		getResp *clientv3.GetResponse
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

	//读取/cron/jobs/为前缀的所有key
	if getResp,err = kv.Get(context.TODO(),"/cron/jobs/",clientv3.WithPrefix()); err != nil{
		fmt.Println(err)
	}else{
		//fmt.Println(getResp.Kvs)
		fmt.Print()
	}

	for kvs,value := range getResp.Kvs{
		fmt.Println(kvs,string(value.Value))
	}
}