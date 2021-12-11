package services

import (
	// "bytes"
	"fmt"
	// "log"

	// "github.com/Azure/azure-storage-blob-go/azblob"
)

type DetailService struct {
	ReportBlob []byte
}

func (s *DetailService) SearchBlob(reportName string) {
	fileName := reportName + ".json"
	blobURL := containerURL.NewBlockBlobURL(fileName)
	fmt.Println(blobURL)
	/*
	// Here's how to download the blob
	blobURL.Download()
	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)

	// NOTE: automatically retries are performed if the connection fails
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	// read the body into a buffer
	downloadedData := bytes.Buffer{}
	_, err = downloadedData.ReadFrom(bodyStream)
	if err != nil {
		log.Fatalln(err)
	}
	*/
}
