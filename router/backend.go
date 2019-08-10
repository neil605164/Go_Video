package router

import (
	"Go_Video/app/handler/video"

	"github.com/gin-gonic/gin"
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {
	backend := r.Group("/backend")
	{
		backend.POST("/upload_video", video.UploadVideo)
	}
}
