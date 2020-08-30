/*
handles the image operations
*/

package image

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

// s3 bucket name
var bucket = "bucket.shubhamdhanera.com"

var ses *session.Session

// init method will intialize the session
func init() {
	ses = session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("us-west-2")},
		Profile: "dhanera",
	}))
}

// HandleFile ...
func HandleFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	imageID, _ := uuid.NewUUID()
	reader := bytes.NewReader(buffer)

	// uploading image to s3
	err := upload(reader, imageID.String(), http.DetectContentType(buffer))
	if err != nil {
		return "", err
	}
	return imageID.String(), nil
}

// upload ...
// uploading image to s3
func upload(fileReader *bytes.Reader, imageID string, fileType string) error {

	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(imageID),
		Body:        fileReader,
		ContentType: aws.String(fileType),
	}

	svc := s3.New(ses)
	_, err := svc.PutObject(params)

	if err != nil {
		log.Println("error while updloading on the s3:", err)
	} else {
		fmt.Println("file uploaded to s3")
	}
	return err
}
