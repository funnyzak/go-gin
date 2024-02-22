package mygin

import (
	"fmt"
	"go-gin/pkg/utils"
	"go-gin/pkg/utils/file"
	"go-gin/pkg/utils/image"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AttachmentUpload struct {
	URLPrefix        string   // URLPrefix is the prefix of the uploaded file url
	MaxSize          int64    // MaxSize is the max file size, default is 2MB
	AllowTypes       []string // AllowTypes is the allowed file types
	FormName         string   // FormName is the form name for the file, default is "file"
	StoreDir         string   // StoreDir is the directory to store the uploaded file
	CreateDateDir    bool     // CreateDateDir is the flag to create date directory
	KeepOriginalName bool     // KeepOriginalName is the flag to keep the original file name
}

type AttachmentUploadedFile struct {
	Url          string `json:"url"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	Size         int64  `json:"size"`
	MiMe         string `json:"mime"`
	Width        int    `json:"width"`     // Width is the width of the image
	Height       int    `json:"height"`    // Height is the height of the image
	Ext          string `json:"ext"`       // Ext is the file extension, eg: .jpg
	MD5          string `json:"md5"`       // MD5 is the md5 of the file
	SavePath     string `json:"save_path"` // SavePath is the path to save the file
}

// Upload uploads the file, and returns the uploaded file info.
func (a *AttachmentUpload) Upload(c *gin.Context, subDir ...string) (*AttachmentUploadedFile, error) {
	result := &AttachmentUploadedFile{}
	form_file, err := c.FormFile(a.FormName)
	if err != nil {
		return result, err
	}
	if form_file.Size > a.MaxSize {
		return result, fmt.Errorf("file size too large")
	}
	form_file_ext := strings.ToLower(filepath.Ext(form_file.Filename)) // eg: .jpg
	form_file_fileilename := form_file.Filename
	form_file_fileize := form_file.Size
	form_file_mime := form_file.Header.Get("Content-Type")

	if len(a.AllowTypes) > 0 && !utils.InArrayString(form_file_mime, a.AllowTypes) {
		return result, fmt.Errorf("file type not allowed")
	}

	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")

	saveName := utils.GenHexStr(32) + form_file_ext
	if a.KeepOriginalName {
		saveName = form_file_fileilename
	}

	savePath := a.StoreDir
	url := a.URLPrefix
	for _, v := range subDir {
		if v != "" {
			savePath = path.Join(savePath, v)
			url = fmt.Sprintf("%s/%s", url, v)
		}
	}

	if a.CreateDateDir {
		savePath = path.Join(savePath, year+month+day)
		url = fmt.Sprintf("%s/%s", url, year+month+day)
	}
	url = fmt.Sprintf("%s/%s", url, saveName)

	if err := file.MkdirAllIfNotExists(savePath, os.ModePerm); err != nil {
		return result, err
	}

	if err := c.SaveUploadedFile(form_file, path.Join(savePath, saveName)); err != nil {
		return result, err
	}

	md5, _ := file.FileMD5(path.Join(savePath, saveName))
	w, h, _ := image.GetImageSize(path.Join(savePath, saveName))

	result.Url = url
	result.Name = saveName
	result.OriginalName = form_file_fileilename
	result.Size = form_file_fileize
	result.MiMe = form_file_mime
	result.Ext = form_file_ext
	result.MD5 = md5
	result.Width = w
	result.Height = h
	result.SavePath = savePath
	return result, nil
}

func NewAttachmentUpload() *AttachmentUpload {
	return &AttachmentUpload{
		URLPrefix:  "/upload",
		MaxSize:    1024 * 1024 * 2,
		AllowTypes: []string{"image/jpeg", "image/png", "image/gif", "image/jpg"},
		FormName:   "file",
		StoreDir:   "./upload",
	}
}
