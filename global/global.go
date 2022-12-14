package g

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"main/modal/config"
)

var (
	Mdb *gorm.DB
	//Rdb
	Logger *zap.Logger
	Config *config.Config
)
