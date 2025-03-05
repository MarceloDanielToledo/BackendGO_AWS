package secretmanager

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"backendgo_aws/models"
	"backendgo_aws/awsgo"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret
	fmt.Println("> Get Secret: ", secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if  err != nil {
		fmt.Println("> Error in GetSecret: ", err)
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)
	fmt.Println("> Secret: ", dataSecret)
	return dataSecret, nil

}