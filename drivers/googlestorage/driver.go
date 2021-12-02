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
func (conn *Connection) Upload(url, name string, file *multipart.FileHeader) (res string, err error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return res, err
	}

	err = write(client, conn.BucketName, name, url, file)

	return res, err
}

func write(client *storage.Client, bucket, object, url string, file *multipart.FileHeader) error {
	ctx := context.Background()

	file.Filename = url
	src, err := file.Open()
	if err != nil {
		return err
	}

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	_, err = io.Copy(wc, src)
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

	wc := client.Bucket(bucket).Object(object)
	err := wc.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Connection) GetPresignedUrl(fileUrl string) (signedUrl string, err error) {
	// ctx := context.Background()

	defaultExpTime := time.Duration(conn.ExpTime)
	expTime := time.Now().Add(defaultExpTime * time.Second)
	// fmt.Println("KEY :", conn.PrivateKey)
	// fmt.Println("IAMEmail :", conn.IAMEmail)
	// fmt.Println("Expired :", conn.ExpTime)

	//! PLAN 1
	// creds, _ := google.FindDefaultCredentials(ctx, storage.ScopeReadWrite)
	// cfg, _ := google.JWTConfigFromJSON(creds.JSON)

	// sakeyFile := "key/go-ticket-325413-020379b56a5d.json"
	// saKey, err := ioutil.ReadFile(sakeyFile)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cfg, err := google.JWTConfigFromJSON(saKey)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// opts := &storage.SignedURLOptions{
	// 	Method:         "GET",
	// 	GoogleAccessID: cfg.Email,
	// 	PrivateKey:     cfg.PrivateKey,
	// 	Expires:        expTime,
	// }

	//! PLAN 2
	// test := conn.PrivateKey
	// test = strings.Replace(test, "\n", "test", -1)
	// fmt.Println("reader :", test)
	// key := []byte(conn.PrivateKey)
	// fmt.Println("KEY :", key)
	// key2 := string(key)
	// fmt.Println("KEY 2 :", key2)

	// fmt.Println("KEY 3 :", cfg.PrivateKey)
	// key4 := string(cfg.PrivateKey)
	// fmt.Println("KEY 4 :", key4)

	// bytesReader := bytes.NewReader([]byte(conn.PrivateKey))
	// bufReader := bufio.NewReader(bytesReader)
	// keyPrivate, _ := bufReader.ReadBytes('\n')
	// fmt.Println("KEY 5 :", keyPrivate)
	// key6 := string(keyPrivate)
	// fmt.Println("KEY 6 :", key6)

	// d := make([]byte, base64.StdEncoding.DecodedLen(len([]byte(conn.PrivateKey))))
	// n, _ := base64.StdEncoding.Decode(d, []byte(conn.PrivateKey))

	// d = d[:n]
	// private, _ := x509.ParsePKIXPublicKey(d)
	// fmt.Println("KEY 7 :", d)
	// fmt.Println("KEY 8 :", string(d))
	// fmt.Println("KEY 9 :", private)
	// keyPrivate, _ := bytesReader.
	// block, _ := pem.Decode([]byte(conn.PrivateKey))
	// rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// cert, err := x509.ParseCertificates(block.Bytes)
	// s, _ := ssh.ParsePrivateKey(block.Bytes)
	// fmt.Println("public key :", cfg.PrivateKey)
	opts := &storage.SignedURLOptions{
		Method:         "GET",
		GoogleAccessID: conn.IAMEmail,
		PrivateKey:     []byte(conn.PrivateKey),
		Expires:        expTime,
	}

	// c, err := credentials.NewIamCredentialsClient(ctx)
	// if err != nil {
	// 	return signedUrl, err
	// }

	//! PLAN 3
	// opts := &storage.SignedURLOptions{
	// 	Method:         "GET",
	// 	GoogleAccessID: conn.IAMEmail,
	// 	SignBytes: func(b []byte) ([]byte, error) {
	// 		req := &credentialspb.SignBlobRequest{
	// 			Payload: b,
	// 			Name:    conn.IAMEmail,
	// 		}
	// 		resp, err := c.SignBlob(ctx, req)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return resp.SignedBlob, err
	// 	},
	// 	Expires: expTime,
	// }

	fmt.Println("GetPresignedURL :", fileUrl)
	signedUrl, err = storage.SignedURL(conn.BucketName, fileUrl, opts)
	if err != nil {
		return signedUrl, err
	}

	return signedUrl, nil
}
