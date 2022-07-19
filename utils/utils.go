package utils

import (
	"fmt"
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
	fmt.Println(RootPath())
	err := godotenv.Load(RootPath() + "/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
