package redis

import (
	"dqh-test/libs/config"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

var redisDB = map[string]*redis.Client{}

func GetRedisConnect(key string) *redis.Client {
	_, ok := redisDB[key]
	if !ok {
		redisConf, err := config.GetConf("redis", key)
		if err != nil {
			log.Fatal(err)
		}
		var option = &redis.Options{
			Addr: redisConf["host"].(string) + ":" + strconv.Itoa(redisConf["port"].(int)),
			DB:   redisConf["select"].(int),
		}

		connection := redis.NewClient(option)
		redisDB[key] = connection
	}
	connection := redisDB[key]

	_, ping_err := connection.Ping().Result()
	if ping_err != nil {
		log.Fatal(ping_err)
	}
	return connection
}
