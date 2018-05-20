package components

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/timidsmile/pspace/setting"
)

var (
	RedisClient *redis.Pool
	clusterName string
	host        string
	port        int
)

const redisTemplate = "%s:%d"

func getRedisAddr() string {
	return fmt.Sprintf(redisTemplate, host, port)

}

func InitRedis(cfg *setting.Config) error {
	clusterName = cfg.Redis.Cluster
	host = cfg.Redis.Host
	port = cfg.Redis.Port
	return nil
}

func Set(key string, value string) error {
	redisClient, err := redis.Dial("tcp", getRedisAddr())
	defer redisClient.Close()

	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return err
	}

	_, err = redisClient.Do("SET", key, value)

	return err
}

func Get(key string) (string, error) {
	redisClient, err := redis.Dial("tcp", getRedisAddr())
	defer redisClient.Close()

	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println(err)
		fmt.Println("redis connect error")
		return "", err
	}

	value, err := redis.String(redisClient.Do("GET", key))

	return value, err
}

func SetEx(key string, value string, ttl int) error {
	conn, err := redis.Dial("tcp", getRedisAddr())
	defer conn.Close()

	// 若连接出错，则打印错误信息，返回
	if err != nil {
		fmt.Println("redis connect error")
		return err
	}

	_, err = conn.Do("set", key, value, "ex", ttl)
	fmt.Println(fmt.Sprintf("set redis key = %s , value = %s , ttl = %d", key, value, ttl))
	return err
}
