package services

import (
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type DeleteService struct {
}

func (s *DeleteService) Delete(fileName string) {
	blobURL := containerURL.NewBlockBlobURL(fileName)
	// Delete the blob we created earlier.
	_, err := blobURL.Delete(ctx, azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{})
	if err != nil {
		log.Fatalln(err)
	}
}
