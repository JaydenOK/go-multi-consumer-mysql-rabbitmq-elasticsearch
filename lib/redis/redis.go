package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	redisClient *redis.Client
)

func InitRedisClient() {
	host := viper.GetString("redis.host")
	password := viper.GetString("redis.password")
	port := viper.GetString("redis.port")
	redisClient = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               host + ":" + port,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           password,
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		fmt.Println("redis连接异常:", err.Error())
		panic(err)
	}

}

func GetRedisClient() *redis.Client {
	return redisClient
}
