package aws

import (
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"path/filepath"
	"testing"
)

type AWSTestSuite struct {
	suite.Suite
	file *os.File
}

func (a *AWSTestSuite) SetupTest() {
	err := godotenv.Load(filepath.Join(utils.RootPath(), ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Init()

	a.file, _ = os.OpenFile(
		"test.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
		// 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644), // 파일 권한은 644
	)
}

func (a *AWSTestSuite) TestS3Env() {
	a.Equal("podoting-live", S3.bucket)
	a.Equal("ap-northeast-2", S3.region)
}

func (a *AWSTestSuite) TestUploadFile() {
	file, err := S3.UploadFile(a.file, "", ".txt")
	a.Nil(err)

	if file != nil {
		a.Equal("test.txt", *file.Key)
	}
}

func TestAWSTestSuite(t *testing.T) {
	suite.Run(t, new(AWSTestSuite))
}
