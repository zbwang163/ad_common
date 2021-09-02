package db

import "github.com/go-redis/redis"

var (
	Redis       map[string]*redis.Client
	redisConfig = map[string]int{
		"ad.info.account_server": 1,
		"ad.info.content_server": 2,
	}
)

func InitRedis(psm string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",               // no password set
		DB:       redisConfig[psm], // use default DB
	})
	Redis[psm] = client
}
