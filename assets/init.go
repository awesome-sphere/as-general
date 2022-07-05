package assets

import (
	"context"
	"log"

	"github.com/awesome-sphere/as-general/utils"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio() {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := utils.GetenvOr("MINIO_ACCESS_KEY", "minio")
	secretAccessKey := utils.GetenvOr("MINIO_SECRET_KEY", "minio123")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient)

	UploadAssets(ctx, minioClient, "assets/trailers/", "trailers", "video/mp4")
	UploadAssets(ctx, minioClient, "assets/posters/", "posters", "image/jpeg")
}
