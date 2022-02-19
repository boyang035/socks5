package config

import (
	"path/filepath"
)

var RootPath,_ = filepath.Abs(".")
var ShadowPath = RootPath + "\\bin\\shadowsocks2-win64.exe"
