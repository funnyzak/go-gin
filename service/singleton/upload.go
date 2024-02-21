package singleton

import (
	"go-gin/pkg/mygin"
)

var (
	AttachmentUpload *mygin.AttachmentUpload
)

func LoadUpload() {
	AttachmentUpload = &mygin.AttachmentUpload{
		BaseURL:          Conf.Site.BaseURL + Conf.Upload.VirtualPath,
		MaxSize:          Conf.Upload.MaxSize,
		AllowTypes:       Conf.Upload.AllowTypes,
		FormName:         "file",
		StoreDir:         Conf.Upload.Dir,
		CreateDateDir:    Conf.Upload.CreateDateDir,
		KeepOriginalName: Conf.Upload.KeepOriginalName,
	}
}
