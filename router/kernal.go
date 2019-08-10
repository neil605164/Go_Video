package router

import (
	"Go_Video/app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	r.Use(middleware.WriteLog)

	// 載入 router 設定
	LoadBackendRouter(r)

	// 載入測試用API
	if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" {
		LoadTestRouter(r)
	}
}
