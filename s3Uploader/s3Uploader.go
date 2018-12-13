package s3Uploader

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

type S3UploadInfo struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
	BucketName      string
}

func FileUploadToS3(s3Info S3UploadInfo, filePath, fileName string) {
	// ファイルの読み込み
	file, err := os.Open(fmt.Sprintf("%s%s", filePath, fileName))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// S3アップロードの準備
	cre := credentials.NewStaticCredentials(
		s3Info.AccessKeyId,
		s3Info.SecretAccessKey,
		"")
	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(s3Info.Region),
	})

	// ファイルのアップロード
	_, err = cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Info.BucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
}
