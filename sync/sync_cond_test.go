package sync
import (
"sync"
"testing"
)

func TestCond(t *testing.T){
	var (
		mailbox uint8
		lock sync.RWMutex
	)
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())
}


