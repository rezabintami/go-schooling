package googlestorage

import (
	"context"
	"io"
	"os"

	"cloud.google.com/go/storage"
)

type Connection struct {
	BucketName string
	ProjectID  string
}

// Upload ...
func (conn *Connection) Upload(url, name string) (res string, err error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return res, err
	}

	err = write(client, conn.BucketName, name, url)

	return res, err
}

func write(client *storage.Client, bucket, object, url string) error {
	ctx := context.Background()
	f, err := os.Open(url)
	if err != nil {
		return err
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	_, err = io.Copy(wc, f)
	if err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (conn *Connection) Delete(name string) (err error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	err = remove(client, conn.BucketName, name)

	return err
}

func remove(client *storage.Client, bucket, object string) error {
	ctx := context.Background()
	object = object[1:]

	// client.Bucket(bucket).Objects()
	wc := client.Bucket(bucket).Object(object)
	err := wc.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

