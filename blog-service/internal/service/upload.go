package service

import (
	"GoprogrammingTourBook/blog-service/global"
	"GoprogrammingTourBook/blog-service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not upported")
	}
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.ChexkMaxSize(fileType, file) {
		return nil, errors.New("exceeded maxinum file limit")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permisions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
