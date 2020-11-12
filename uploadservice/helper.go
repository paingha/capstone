package main

import (
	"bytes"
	"io"
	"io/ioutil"

	"bitbucket.com/irb/api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
)

func streamToByte(stream io.Reader) []byte {
	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(stream)
	return fileBuffer.Bytes()
}

func writeLocal(fileParam models.FileParam) {
	logrus.Infof("Writing file %s ... ", fileParam.Name)
	file := streamToByte(fileParam.File)
	err := ioutil.WriteFile("../api/files/public"+"/"+fileParam.Name, file, 0644)
	if err != nil {
		handleError(err, "An error occured while writing file.")
	}
}

func uploadFile(fileParam models.FileParam) {
	logrus.Infof("Uploading file %s ... ", fileParam.Name)
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretKey, cfg.AWSSessionToken),
	})
	if err != nil {
		handleError(err, "AWS session error.")
	}
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.AWSS3Bucket),
		Key:    aws.String(fileParam.Name),
		Body:   fileParam.File,
	})
	if err != nil {
		handleError(err, "AWS file upload error.")
	}
}
