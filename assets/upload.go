package assets

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/minio/minio-go/v7"
)

func createBucket(ctx context.Context, minioClient *minio.Client, bucketName string) {
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func UploadAssets(ctx context.Context, minioClient *minio.Client, dirName string, bucketName string, contentType string) {

	createBucket(ctx, minioClient, bucketName)

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := dirName + fileName

		info, err := minioClient.FPutObject(ctx, bucketName, fileName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s to %s\n", fileName, info.Bucket)
	}
}
