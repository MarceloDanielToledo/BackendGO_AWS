package routers

import (
	"backendgo_aws/awsgo"
	"backendgo_aws/bd"
	"backendgo_aws/models"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func GetImage(ctx context.Context, fileType string, request events.APIGatewayProxyRequest, claim models.Claim) models.ResponseAPI {
	var response models.ResponseAPI
	response.StatusCode = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		response.Message = "You must send the parameter ID"
		return response
	}

	profile, err := bd.GetProfile(ID)
	if err != nil {
		response.Message = fmt.Sprintf("Error getting profile: %s", err)
		response.StatusCode = 500
		return response
	}

	var filename string
	switch fileType {
	case "A":
		filename = profile.Avatar
	case "B":
		filename = profile.Banner
	}

	fmt.Println("Filename: " + filename)

	svc := s3.NewFromConfig(awsgo.Cfg)

	file, err := downloadFromS3(ctx, svc, filename)
	if err != nil {
		response.Message = fmt.Sprintf("Error downloading file: %s", err)
		response.StatusCode = 500
		return response
	}

	response.CustomResp = &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       file.String(),
		Headers: map[string]string{
			"Content-Type":        "applitacion/octet-stream",
			"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
		},
	}

	response.StatusCode = 200
	return response
}

func downloadFromS3(ctx context.Context, svc *s3.Client, filename string) (*bytes.Buffer, error) {
	bucket := ctx.Value(models.Key("bucketName")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()
	fmt.Println("Downloaded file: " + filename)
	file, err := io.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(file)
	return buffer, nil
}
