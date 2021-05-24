package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
使用etcd实现乐观锁
1.op操作
2.txn事务: if else then
//lease实现锁自动过期，节点宕机以后，释放锁
 */



func main(){
	var(
		config clientv3.Config
		client *clientv3.Client
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		keepResp *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		ctx context.Context
		cancelFunc context.CancelFunc
		kv clientv3.KV
		txn clientv3.Txn
		txnResp *clientv3.TxnResponse
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

	//1.上锁（创建租约 自动续租 拿着租约去抢占一个key）
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
	//准备一个用于取消自动续租的context
	ctx,cancelFunc = context.WithCancel(context.TODO())

	//确保函数退出后，自动续租会停止
	defer cancelFunc()
	defer lease.Revoke(context.TODO(),leaseId)
	//自动续租
	//ctx,_ := context.WithTimeout(context.TODO(),5*time.Second)
	////续租了5秒，停止了续租，10秒的生命期 = 15秒的生命期
	if keepRespChan,err  = lease.KeepAlive(ctx,leaseId); err != nil{
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
	//if 不存在key，then设置它，else 抢锁失败
	kv = clientv3.NewKV(client)

	//创建事务
	txn = kv.Txn(context.TODO())
	//定义事务

	//如果key不存在
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"),"=",0)).Then(
		clientv3.OpPut("/cron/lock/job9","xxx",clientv3.WithLease(leaseId))).Else(clientv3.OpGet("/cron/lock/job9"))//否则抢锁失败

	//提交事务
	if txnResp,err = txn.Commit(); err != nil{
		fmt.Println(err)
		return
	}

	//判断是否抢到了锁
	if !txnResp.Succeeded{
		fmt.Println("锁被占用:",string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}
	//2.处理业务
	fmt.Println("处理任务")
	time.Sleep(5 *time.Second)
	//在锁内，很安全

	//3.释放锁(取消自动续租,释放租约)
	//defer 会把租约释放掉，关联的kv就被删除了


}