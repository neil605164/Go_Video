package video

import (
	"Go_Video/app/global/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadVideo 上傳影片
// @Summary 上傳影片
// @Description 上傳 mp4 影片
// @Tags Video
// @Accept mpfd
// @Produce application/json
// @Param file formData file true "影片檔"
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /backend/upload_video [POST]
func UploadVideo(c *gin.Context) {

	// 取得影片資訊
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// fmt.Println(file)
	// fmt.Println(file.Filename)
	// fmt.Println(file.Size)

	// filename := filepath.Base(file.Filename)
	filename := fmt.Sprintf("./upload/%s", file.Filename)
	fmt.Println(filename)

	// if err := c.SaveUploadedFile(file, filename); err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 	return
	// }

	c.JSON(http.StatusOK, helper.Success("123"))
}
