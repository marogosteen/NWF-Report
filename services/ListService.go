package services

import (
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type ListService struct {
	ReportList []string
}

// TODO 一度に1セグメントを呼ぶべきではない。Pager実装（5000件）
// ContainerURL で ListBlobs メソッドを使用して、コンテナー内のファイルの一覧を取得します。 
// ListBlobs は、指定された マーカー から開始して単一セグメントの BLOB (最大 5,000) を返します。 
// 空のマーカーを使用して最初から列挙を開始します。 BLOB 名は辞書式順序で返されます。 
// セグメントの取得後、それを処理し、前に返されたマーカーを渡す ListBlobs を再度呼び出します。
func (s *ListService) Fetch() {
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
			s.ReportList = append(s.ReportList, blobInfo.Name)
		}
	}
}
