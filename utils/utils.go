package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	_, f, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(f)
)

type MapItem struct {
	Key, Value interface{}
}

type MapSlice []MapItem

func (ms MapSlice) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func RootPath() string {
	return basePath[:len(basePath)-5]
}

func SetEnv() {
	if os.Getenv("APP_ENV") != "production" {
		fmt.Println("Loading .env file")

		path := filepath.Join(RootPath(), ".env")
		fmt.Println("filepath", filepath.Join(RootPath(), ".env"))

		err := godotenv.Load(path)

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func BuildMapSliceByMap(m map[string]interface{}) MapSlice {
	ms := make(MapSlice, 0)
	for k, v := range m {
		ms = append(ms, MapItem{Key: k, Value: v})
	}
	return ms
}

func ParseUint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	return uint(i), err
}
