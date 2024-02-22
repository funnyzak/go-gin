package singleton

import (
	"go-gin/pkg/mygin"
	"go-gin/pkg/utils/file"
	"os"
)

var (
	AttachmentUpload *mygin.AttachmentUpload
)

func LoadUpload() {
	AttachmentUpload = &mygin.AttachmentUpload{
		URLPrefix:        Conf.Upload.URLPrefix,
		MaxSize:          Conf.Upload.MaxSize,
		AllowTypes:       Conf.Upload.AllowTypes,
		FormName:         "file",
		StoreDir:         Conf.Upload.Dir,
		CreateDateDir:    Conf.Upload.CreateDateDir,
		KeepOriginalName: Conf.Upload.KeepOriginalName,
	}
	file.MkdirAllIfNotExists(Conf.Upload.Dir, os.ModePerm)
}
