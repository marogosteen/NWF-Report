package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type UploadService struct {
	FileName   string
	ReportBlob []byte
}

func (s *UploadService) Upload() {
	blobURL := containerURL.NewBlockBlobURL(s.FileName)
	azblob.UploadBufferToBlockBlob(
		ctx,
		s.ReportBlob,
		blobURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16,
		},
	)
}

// azure blob strageにUploadせず、JsonをLocalに書き込む。
func (s *UploadService) SampleUpload() {
	var i map[string]interface{}
	err := json.Unmarshal(s.ReportBlob, &i)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(s.FileName)
	if err != nil {
		log.Fatalln(err)
	}
	json.NewEncoder(file).Encode(i)
}
