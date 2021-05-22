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
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		kv clientv3.KV
		putResp *clientv3.PutResponse
		leaseId clientv3.LeaseID
		getResp *clientv3.GetResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
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
	}

	//申请一个lease(租约)
	lease = clientv3.NewLease(client)
	fmt.Println("---")
	//申请一个10s的租约
	if leaseGrantResp,err  = lease.Grant(context.Background(),10); err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("---")
	//拿到租约的ID
	leaseId = clientv3.LeaseID(leaseGrantResp.ID)

	//自动续租
	//ctx,_ := context.WithTimeout(context.TODO(),5*time.Second)
	////续租了5秒，停止了续租，10秒的生命期 = 15秒的生命期
	if keepRespChan,err  = lease.KeepAlive(context.TODO(),leaseId); err != nil{
		fmt.Println(err)
	}
	//
	go func() {
		for  {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效")
					goto END
				}else{ //每秒续租一次，会收到一次应答
					fmt.Println("收到自动续租应答:",keepResp.ID)
				}
			}
		}
		END:
	}()
	//获取kv api子集

	kv = clientv3.NewKV(client)
	//Put一个kv,让它与租约关联起来，从而实现10秒后自动过期
	if putResp, err = kv.Put(context.TODO(),"/cron/lock/job1","",clientv3.WithLease(leaseId));err != nil{
		fmt.Println(err)
	}
	fmt.Println("写入成功:",putResp.Header.Revision)


	//定时探测key是否过期
	for {
		if getResp,err = kv.Get(context.TODO(),"/cron/lock/job1"); err != nil{
			fmt.Println(err)
			return
		}
		if len(getResp.Kvs) == 0{
			fmt.Println("kv 过期了")
			break
		}
		fmt.Println("还没过期",getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}