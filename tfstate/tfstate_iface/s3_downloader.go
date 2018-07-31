package tfstate_iface

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3DownloaderAPI interface {
	Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error)
}
