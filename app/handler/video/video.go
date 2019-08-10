package video

import (
	"Go_Video/app/global/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadVideo 上傳影片
// @Summary 上傳影片
// @Description 上傳 mp4 影片
// @Tags Video
// @Produce  multipart/form-data
// @Param video formData file true "影片檔"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/upload_video [POST]
func UploadVideo(c *gin.Context) {
	c.JSON(http.StatusOK, helper.Success("123"))
}
