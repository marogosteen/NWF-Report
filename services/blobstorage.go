package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/nwf-report/config"
)

func Hoge() {
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
	containerURL := azblob.NewContainerURL(*URL, p)

	// Create the container
	fmt.Printf("Creating a container named %s\n", containerName)
	ctx := context.Background() // This example uses a never-expiring context
	// _, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a file to test the upload and download.
	fmt.Printf("Creating a dummy file to test the upload and download\n")
	data := []byte("hello world this is a blob\n")
	fileName := "hoge.txt"
	err = ioutil.WriteFile(fileName, data, 0700)

	// Here's how to upload a blob.
	blobURL := containerURL.NewBlockBlobURL(fileName)
	file, err := os.Open(fileName)

	// You can use the low-level Upload (PutBlob) API to upload files. Low-level APIs are simple wrappers for the Azure Storage REST APIs.
	// Note that Upload can upload up to 256MB data in one shot. Details: https://docs.microsoft.com/rest/api/storageservices/put-blob
	// To upload more than 256MB, use StageBlock (PutBlock) and CommitBlockList (PutBlockList) functions.
	// Following is commented out intentionally because we will instead use UploadFileToBlockBlob API to upload the blob
	// _, err = blobURL.Upload(ctx, file, azblob.BlobHTTPHeaders{ContentType: "text/plain"}, azblob.Metadata{}, azblob.BlobAccessConditions{})
	// handleErrors(err)

	// The high-level API UploadFileToBlockBlob function uploads blocks in parallel for optimal performance, and can handle large files as well.
	// This function calls StageBlock/CommitBlockList for files larger 256 MBs, and calls Upload for any file smaller
	fmt.Printf("Uploading the file with blob name: %s\n", fileName)
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
}
