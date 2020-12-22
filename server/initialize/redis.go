package initialize

import (
	"github.com/go-redis/redis"
)

func Redis()  {
	redisConfig := global.BB_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		Password: redisConfig.Password,
		DB: redisConfig.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.BB_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.BB_LOG.Info("redis connect ping response:", zap.String("pong",pong))
		global.BB_REDIS = client
	}
}