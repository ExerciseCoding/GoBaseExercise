package distributed_lock
import(
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"time"
)
func incr(){
	client := redis.NewClient(&redis.option{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	var lockKey = "counter_lock"
	var counterKey = "counter"

	resp := client.SetNX(lockKey,1,time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess{
		fmt.Println("lock failed")
		return
	}

	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err != nil {
		cntValue++
		resp := client.Set(counterKey,cntValue,0)
		_, err := resp.Result()
		if err != nil {
			println("set value error")
		}

	}

	println("current counter ")

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success")
	} else{
		println("unlock failed")
	}
}