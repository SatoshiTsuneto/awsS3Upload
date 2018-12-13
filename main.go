package main

import "awsS3Upload/s3Uploader"

func main() {
	// S3の設定
	s3Info := s3Uploader.S3UploadInfo{
		AccessKeyId:     "アクセスキー",
		SecretAccessKey: "シークレットキー",
		Region:          "ap-northeast-1", // 東京リージョン
		BucketName:      "バケット名",
	}

	// アップロードするファイルのディレクトリと名前
	filePath := "ファイルパス"
	fileName := "ファイル名"

	// ファイルのアップロード
	s3Uploader.FileUploadToS3(s3Info, filePath, fileName)
}
