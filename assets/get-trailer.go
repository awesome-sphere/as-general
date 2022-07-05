package assets

import (
	"encoding/base64"
	"log"

	"github.com/minio/minio-go/v7"
)

func getTrailerObject(fileName string) *minio.Object {
	obj, err := MINIO_CLIENT.GetObject(CTX, "trailers", fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Printf("Successfully retrieved %s", fileName)
	return obj
}

func GetTrailer(fileName string) string {
	if obj := getTrailerObject(fileName); obj != nil {
		fileInfo, _ := obj.Stat()
		buffer := make([]byte, fileInfo.Size)
		obj.Read(buffer)
		defer obj.Close()
		return base64.StdEncoding.EncodeToString(buffer)
	}
	return ""
}