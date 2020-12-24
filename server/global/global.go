package global

import (
	"beau-blog/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/spf13/viper"
)

var (
	BB_DB     *gorm.DB
	BB_REDIS  *redis.Client
	BB_CONFIG config.Server
	BB_VP     *viper.Viper
	BB_LOG    *zap.Logger
)
