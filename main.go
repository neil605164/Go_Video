package main

import (
	"Go_Video/app/global"
	"Go_Video/app/global/helper"
	"Go_Video/app/model"
	_ "Go_Video/docs"
	"Go_Video/router"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var c *gin.Context

// @title 上傳影片
// @version 1.0
// @description 上傳影片練習用
// @termsOfService https://google.com
// @contact.name Neil_Hsieh
// @contact.url https://google.com
// @contact.email neil605164@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9999
// @BasePath /
func main() {
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.ErrorHandle(global.WarnLog, fmt.Sprintf("[❌ Fatal❌ ]: %v", err), "")
			fmt.Println("[❌ Fatal❌ ]:", err)
		}
	}()

	// 開發時，console視窗不顯示有顏色的log
	gin.DisableConsoleColor()
	r := gin.Default()

	// 載入環境設定(所有動作須在該func後執行)
	global.Start()

	// 檢查 DB 機器服務
	model.DBPing()

	// 檢查 Redis 機器服務
	// repository.RedisPing()

	// 背景
	// go task.Schedule()

	// 載入router設定
	router.RouteProvider(r)
	r.Run(":8080")
}
