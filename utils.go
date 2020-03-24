package main

import (
	"path/filepath"
	"runtime"
)

func rootProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)
	return basepath
}
