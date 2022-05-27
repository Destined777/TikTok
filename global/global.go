package global

import (
	"TikTok/config"
	"gorm.io/gorm"
)

var (
	Config      *config.Config
	DB          *gorm.DB
)
