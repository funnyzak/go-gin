package model

import (
	"go-gin/pkg/mygin"
	"go-gin/pkg/utils"

	"gorm.io/gorm"
)

type Attachment struct {
	Common
	Num         string `json:"num,omitempty" gorm:"column:num"`
	Name        string `json:"name,omitempty" gorm:"column:name"`
	SavePath    string `json:"-" gorm:"column:save_path"`
	MD5         string `json:"md5,omitempty" gorm:"column:md5"`
	Description string `json:"description,omitempty" gorm:"column:description"`
	Url         string `json:"url,omitempty" gorm:"-"`
	Size        int64  `json:"size,omitempty" gorm:"column:size"`
	MiMe        string `json:"mime,omitempty" gorm:"column:mime"`
	CreatedUser uint64 `json:"created_user,omitempty" gorm:"column:created_user"`
	CoverUrl    string `json:"cover_url,omitempty" gorm:"column:cover_url"`
	Width       int    `json:"width,omitempty" gorm:"column:width"`
	Height      int    `json:"height,omitempty" gorm:"column:height"`
	ExtConfig   string `json:"ext_config,omitempty" gorm:"column:ext_config"`
}

func (a *Attachment) Create(db *gorm.DB) (err error) {
	err = db.Model(&Attachment{}).Create(&a).Error
	return err
}

func (a *Attachment) CreateByAttachment(db *gorm.DB, uploadedFile mygin.AttachmentUploadedFile) (err error) {
	a.Num = utils.GenHexStr(32)
	a.Name = uploadedFile.OriginalName
	a.SavePath = uploadedFile.SavePath
	a.MD5 = uploadedFile.MD5
	a.Url = uploadedFile.Url
	a.Size = uploadedFile.Size
	a.MiMe = uploadedFile.MiMe
	a.Width = uploadedFile.Width
	a.Height = uploadedFile.Height
	err = db.Model(&Attachment{}).Create(&a).Error
	return err
}

func (a *Attachment) GetByNum(db *gorm.DB, num string) (attachment Attachment, err error) {
	err = db.Model(&Attachment{}).Where("num = ?", num).First(&attachment).Error
	return attachment, err
}

func (a *Attachment) GetByMD5(db *gorm.DB, md5 string) (attachment Attachment, err error) {
	err = db.Model(&Attachment{}).Where("md5 = ?", md5).First(&attachment).Error
	return attachment, err
}
