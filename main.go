package main

import (
	"TikTok/config"
	"TikTok/global"
	"TikTok/model"
	"TikTok/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func init() {
	initConfig()
	initDB()
}

// 读取环境配置文件
func initConfig() {
	configFile := "config.json"
	global.Config = config.ReadSettingsFromFile(configFile)
}

func initDB() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second)
	go func(ctx context.Context) {
		conf := global.Config
		settings := conf.DbSettings
		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8",
			settings.Username, settings.Password, settings.Hostname, settings.Dbname)
		var err1 error
		var localDb *gorm.DB
		localDb, err1 = gorm.Open(mysql.Open(connStr), &gorm.Config{})
		if err1 != nil {
			panic("Database connect error," + err1.Error())
		}
		sqlDB, err := localDb.DB()
		if err != nil {
			panic("Database error")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10000)
		sqlDB.SetConnMaxLifetime(time.Minute)
		global.DB = localDb
		err = global.DB.AutoMigrate(&model.LogUser{}, &model.Follow{}, &model.Video{})
		if err != nil {
			return
		}
		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println("context timeout exceeded")
			panic("Failed to initialize database connection")
		case context.Canceled:
			fmt.Println("context cancelled by force. whole process is complete")
		}
	}
}