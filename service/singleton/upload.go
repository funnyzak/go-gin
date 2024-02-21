package singleton

import (
	"go-gin/pkg/mygin"
)

var (
	AttchmentUpload *mygin.AttchmentUpload
)

func LoadUpload() {
	AttchmentUpload = &mygin.AttchmentUpload{
		BaseURL:          Conf.Site.BaseURL + Conf.Upload.VirtualPath,
		MaxSize:          Conf.Upload.MaxSize,
		AllowTypes:       Conf.Upload.AllowTypes,
		FormName:         "file",
		StoreDir:         Conf.Upload.Dir,
		CreateDateDir:    Conf.Upload.CreateDateDir,
		KeepOriginalName: Conf.Upload.KeepOriginalName,
	}
}
