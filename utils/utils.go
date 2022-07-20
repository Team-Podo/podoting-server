package utils

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func RootPath() string {
	return basePath[:len(basePath)-5]
}

func SetEnv() {
	err := godotenv.Load(filepath.Join(RootPath(), ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
