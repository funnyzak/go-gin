package resource

import (
	"embed"
)

//go:embed static
var StaticFS embed.FS

//go:embed template
var TemplateFS embed.FS

func IsTemplateFileExist(name string) bool {
	_, err := TemplateFS.Open(name)
	return err == nil
}
