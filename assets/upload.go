package assets

import (
	"io/ioutil"
	"log"

	"github.com/minio/minio-go/v7"
)

func createBucket(bucketName string) {
	err := MINIO_CLIENT.MakeBucket(CTX, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := MINIO_CLIENT.BucketExists(CTX, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func UploadAssets(dirName string, bucketName string, contentType string) {

	createBucket(bucketName)

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := dirName + fileName

		info, err := MINIO_CLIENT.FPutObject(CTX, bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s to %s\n", fileName, info.Bucket)
	}
}
