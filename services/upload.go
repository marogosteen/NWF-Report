package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/nwf-report/config"
)

var (
	ctx          context.Context
	containerURL azblob.ContainerURL
)

func init() {
	cfg := config.Config
	accountName := cfg.Account
	accountKey := cfg.Key
	containerName := cfg.ContainerName

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// From the Azure portal, get your storage account blob service URL endpoint.
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL = azblob.NewContainerURL(*URL, p)

	// Create the container
	ctx = context.Background()

}

type UploadService struct {
	Title     string
	BestEpoch int
	BestLoss  float64
	Observed  []float64
	Predicted []float64
}

func (s *UploadService) Upload() {
	b, err := json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}

	fileName := s.Title + ".json"
	blobURL := containerURL.NewBlockBlobURL(fileName)

	var i map[string]interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		log.Fatalln(err)
	}

	azblob.UploadBufferToBlockBlob(
		ctx,
		b,
		blobURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16,
		},
	)
}
