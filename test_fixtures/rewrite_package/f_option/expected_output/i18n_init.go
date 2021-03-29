package input_files

import (
	"path/filepath"

	goi18n "github.com/EverlongProject/go-i18n/i18n"
	i18n "github.com/EverlongProject/i18n4go/i18n"
)

var T goi18n.TranslateFunc

func init() {
	T = i18n.Init(filepath.Join("test_fixtures", "rewrite_package", "f_option", "input_files"), i18n.GetResourcesPath())
}
