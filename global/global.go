package g

import (
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"main/modal/config"
)

var (
	Mdb    *gorm.DB
	Rdb    *redis.Client
	Logger *zap.Logger
	Config *config.Config
)
