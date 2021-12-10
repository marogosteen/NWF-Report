package repositories

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"

	"github.com/nwf-report/config"
)

var (
	Ctx          context.Context
	ContainerURL azblob.ContainerURL
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
	ContainerURL = azblob.NewContainerURL(*URL, p)

	// Create the container
	Ctx = context.Background()
}
