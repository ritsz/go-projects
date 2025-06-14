package common

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../..")
)

func FileNameWithoutExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
