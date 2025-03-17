package util

import (
	"path/filepath"
	"runtime"
)

func GetRoot() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(b), "../..")
}
