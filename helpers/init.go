package helpers

import (
	"path"
	"runtime"
)

func InitResource() {
	InitMysql()
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	InitLogger(root+"/logs/all.log", "debug")
}
