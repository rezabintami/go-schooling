package googlestorage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type Connection struct {
	BucketName string
	PrivateKey string
	IAMEmail   string
	ExpTime    int
}

// Upload ...
func (conn *Connection) Upload(url, name string, file multipart.File) (res string, err error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	fmt.Println("Bucket :", conn.BucketName)
	if err != nil {
		return res, err
	}

	err = write(client, conn.BucketName, name, url, file)

	return res, err
}

func write(client *storage.Client, bucket, object, url string, file multipart.File) error {
	ctx := context.Background()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	_, err := io.Copy(wc, file)
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

func (conn *Connection) GetPresignedUrl(fileUrl string) (signedUrl string, err error) {
	defaultExpTime := time.Duration(conn.ExpTime)
	expTime := time.Now().Add(defaultExpTime * time.Second)

	opts := &storage.SignedURLOptions{
		Method:         "GET",
		GoogleAccessID: conn.IAMEmail,
		PrivateKey:     []byte(conn.PrivateKey),
		Expires:        expTime,
	}

	fmt.Println("URL : ", fileUrl)
	signedUrl, err = storage.SignedURL(conn.BucketName, fileUrl, opts)
	if err != nil {
		return signedUrl, err
	}

	return signedUrl, nil
}
