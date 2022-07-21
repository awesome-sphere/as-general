package assets

import (
	"context"
	"log"

	"github.com/awesome-sphere/as-general/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var CTX context.Context
var MINIO_CLIENT *minio.Client

func ConnectMinio() {
	ctx := context.Background()
	endpoint := "minio-service:9000"
	accessKeyID := utils.GetenvOr("MINIO_ACCESS_KEY", "minio")
	secretAccessKey := utils.GetenvOr("MINIO_SECRET_KEY", "minio123")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to Minio")

	CTX, MINIO_CLIENT = ctx, minioClient

	UploadAssets("assets/trailers/", "trailers", "video/mp4")
	UploadAssets("assets/posters/", "posters", "image/jpeg")

}
