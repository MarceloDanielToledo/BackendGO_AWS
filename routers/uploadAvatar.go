package routers

import (
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, fileType string, request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400

	idUser := claim.ID.Hex()
	var fileName string
	var user models.User
	bucket := aws.String(ctx.Value(models.Key("bucket")).(string))

	switch fileType {
	case "A":
		fileName = "avatars/" + idUser + ".jpg"
		user.Avatar = fileName
	case "B":
		fileName = "banners/" + idUser + ".jpg"
		user.Banner = fileName
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		response.Message = "Error parsing media type"
		response.StatusCode = 500
		return response
	}
	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			response.Message = "Error decoding base64"
			response.StatusCode = 500
			return response
		}
		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			response.Message = "Error reading the first part"
			response.StatusCode = 500
			return response
		}

		if err != io.EOF {
			if p.FileName() != "" {
				buff := bytes.NewBuffer(nil)
				if _, err := io.Copy(buff, p); err != nil {
					response.Message = "Error copying the file"
					response.StatusCode = 500
					return response
				}
				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1"),
				})
				if err != nil {
					response.Message = "Error creating the session"
					response.StatusCode = 500
					return response
				}

				if err != nil {
					response.Message = "Error creating the session"
					response.StatusCode = 500
					return response
				}

				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(fileName),
					Body:   &readSeeker{buff},
				})
				if err != nil {
					response.Message = "Error uploading the file"
					response.StatusCode = 500
					return response
				}

				status, err := bd.UpdateProfile(user, idUser)
				if err != nil || !status {
					response.Message = "Error updating the user"
					response.StatusCode = 400
					return response
				}

			}

		}

	} else {
		response.Message = "Content-Type is not multipart"
		response.StatusCode = 500
	}
	response.StatusCode = 200
	response.Message = "File uploaded successfully"
	return response
}
