package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
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
		path := filepath.Join(RootPath(), ".env")
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

func CheckDateFormatInvalid(d string) bool {
	_, err := time.Parse("2006-01-02", d)
	return err != nil
}

func CheckTimeFormatInvalid(t string) bool {
	_, err := time.Parse("15:04", t)
	return err != nil
}

func GenerateOrderKey() string {
	var chars = []rune("ABCDEFGHJKMNPQRSTUVWXYZ23456789")
	var key = make([]rune, 8)
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}
	return string(key)
}

func GenerateOrderDetailKey() string {
	var chars = []rune("ABCDEFGHJKMNPQRSTUVWXYZ23456789")
	var key = make([]rune, 12)
	for i := range key {
		key[i] = chars[rand.Intn(len(chars))]
	}
	return string(key)
}
