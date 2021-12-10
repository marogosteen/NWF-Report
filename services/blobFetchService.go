package services

import (
	"fmt"
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type BlobFetchService struct {
	ReportBlob [][]byte
}

func (b *BlobFetchService) FetchBlobList() {
	// List the container that we have created above
	fmt.Println("Listing the blobs in the container:")
	for marker := (azblob.Marker{}); marker.NotDone(); {
		// Get a result segment starting with the blob indicated by the current Marker.
		listBlob, err := containerURL.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		// ListBlobs returns the start of the next segment; you MUST use this to get
		// the next segment (after processing the current result segment).
		marker = listBlob.NextMarker

		// Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
		for _, blobInfo := range listBlob.Segment.BlobItems {
			fmt.Print("    Blob name: " + blobInfo.Name + "\n")
		}
	}
}
