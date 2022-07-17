package cache



import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

const (


	TENMIN = 60*10*time.Second

)

func SetUp() error{
	ip :=viper.GetString("database.ip")
	port :=6379
	addr:= fmt.Sprintf("%s:%d",ip,port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_,err := rdb.Ping(ctx).Result()
	return err
}


func SetAndTime(key,value string,duration time.Duration) error {
	err := rdb.Set(ctx, key, value, TENMIN).Err()
	return err
}

func Get(key string)  (string,error) {
	val, err := rdb.Get(ctx, key).Result()
	return val,err
}

func IsExists(key string) bool  {
	// 判断key是否存在
	_, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return true
	} else {
		return  false
	}
}

//判断值是否过期:永不过期返回-1,已过期返回-2

func IsExpire(key string) (bool,error)  {
	result, err := rdb.TTL(ctx, key).Result()
	if int64(result)== -2{
		return true,err
	}else {
		return false,err
	}
}

func SetExpire(key string) bool {
	expire := rdb.Expire(ctx, key, 0)
	//fmt.Println(expire.Err())
	//fmt.Println(expire.Name())
	//fmt.Println(expire.Result())
	//fmt.Println(expire.String())
	//fmt.Println(expire.Args())
	//fmt.Println(expire.FullName())
	//fmt.Println(expire.Val())
	return expire.Val()
}






