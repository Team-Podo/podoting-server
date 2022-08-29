package utils

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func CheckFileExtension(fileHeader *multipart.FileHeader) bool {
	extension := strings.ToLower(filepath.Ext(fileHeader.Filename))
	return extension == ".jpg" || extension == ".jpeg" || extension == ".png"
}

func CDNUrlByFilePath(path string) string {
	return os.Getenv("CDN_URL") + "/" + path
}
