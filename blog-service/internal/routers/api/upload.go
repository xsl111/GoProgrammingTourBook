package api

import (
	"GoprogrammingTourBook/blog-service/global"
	"GoprogrammingTourBook/blog-service/internal/service"
	"GoprogrammingTourBook/blog-service/pkg/app"
	"GoprogrammingTourBook/blog-service/pkg/convert"
	"GoprogrammingTourBook/blog-service/pkg/errcode"
	"GoprogrammingTourBook/blog-service/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errResp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
