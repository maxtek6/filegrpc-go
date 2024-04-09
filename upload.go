package filegrpc

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/maxtek6/filegrpc-go/services/upload"
	"google.golang.org/grpc"
)

type UploadClient struct {
	client upload.UploadClient
}

func NewUploadClient(conn *grpc.ClientConn) (*UploadClient, error) {
	if conn == nil {
		return nil, errors.New("filegrpc: nil connection")
	}
	return &UploadClient{
		client: upload.NewUploadClient(conn),
	}, nil
}

func (client *UploadClient) Upload(ctx context.Context, name string, reader io.Reader) ([]byte, error) {
	initResponse, err := client.client.Init(ctx, &upload.InitRequest{
		Name: name,
	})
	if err != nil {
		return nil, fmt.Errorf("filegrpc: failed to initialize upload: %v", err)
	}
	stream, err := client.client.Send(ctx)
	if err != nil {
		return nil, fmt.Errorf("filegrpc: failed to create upload stream: %v", err)
	}
	streaming := true
	for streaming {
		buffer := make([]byte, 1024*1024*4)
		readSize, err := reader.Read(buffer)
		if err != nil {
			return nil, fmt.Errorf("filegrpc: failed to read upload buffer: %v", err)
		}
		stream.Send(&upload.SendRequest{
			UploadId: initResponse.UploadId,
			Chunk:    buffer,
		})
		if readSize < len(buffer) {
			streaming = false
		}
	}
	sendResponse, err := stream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("filegrpc: failed to complete upload: %v", err)
	}
	return sendResponse.Metadata, nil
}
