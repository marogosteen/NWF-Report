package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"

	"github.com/nwf-report/repositories"
)

type UploadService struct {
	FileName   string
	ResultBlob []byte
}

func (s *UploadService) Upload() {
	blobURL := repositories.ContainerURL.NewBlockBlobURL(s.FileName)
	azblob.UploadBufferToBlockBlob(
		repositories.Ctx,
		s.ResultBlob,
		blobURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16,
		},
	)
}

func (s *UploadService) SampleUpload() {
	var i map[string]interface{}
	err := json.Unmarshal(s.ResultBlob, &i)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(s.FileName)
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(file).Encode(i)
}
