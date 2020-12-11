package global

import (
	"go.uber.org/zap"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"beau-blog/config"
)

var (
	BB_DB     *gorm.DB
	BB_REDIS  *redis.Client
	BB_CONFIG config.Server
	BB_LOG    *zap.Logger
)
