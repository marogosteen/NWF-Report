package services

import (
	"io"
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type DetailService struct {
	Reader io.ReadCloser
}

// Download 関数を使用して BLOB をダウンロードします。 BLOB の内容がバッファーに書き込まれる。
func (s *DetailService) SearchBlob(fileName string) {
	blobURL := containerURL.NewBlockBlobURL(fileName)

	// Here's how to download the blob
	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	// NOTE: automatically retries are performed if the connection fails
	s.Reader = downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

}
