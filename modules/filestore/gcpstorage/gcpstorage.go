package gcpstorage

import (
	"context"
	"strings"

	"cloud.google.com/go/storage"

	"github.com/spaceuptech/space-cloud/utils"
)

// GCPStorage holds the GCPStorage client
type GCPStorage struct {
	client *storage.Client
}

// Init initializes a GCPStorage client
func Init(connection, endpoint string) (*GCPStorage, error) {
	ctx := context.TODO()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.Bucket(endpoint).Create(ctx, connection, nil); err != nil {
		if !strings.Contains(err.Error(), "409") {
			return nil, err
		}
	}
	return &GCPStorage{client}, nil
}

// GetStoreType returns the file store type
func (g *GCPStorage) GetStoreType() utils.FileStoreType {
	return utils.GCPStorage
}

// Gracefully close the GCPStorage module
func (g *GCPStorage) Close() error {
	return g.Close()
}
