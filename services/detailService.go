package services

import (
	"bytes"
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type DetailService struct {
	ReportBlob []byte
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
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	// TODO bodyStream.Close()
	// read the body into a buffer
	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)
	if err != nil {
		log.Fatalln(err)
	}

	s.ReportBlob = downloadedData.Bytes()
}
