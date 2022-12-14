package dydb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"lambda-store-nir/service/application/domain"
	"lambda-store-nir/service/application/repositories"
)

type DocumentRepository struct {
	AwsSession *session.Session
	TableName  string
}

func NewDocumentRepository(awsSession *session.Session, tableName string) repositories.DocumentRepository {
	var c repositories.DocumentRepository = &DocumentRepository{
		AwsSession: awsSession,
		TableName:  tableName,
	}
	return c
}

func (s *DocumentRepository) Save(document domain.Document) error {

	item, err := dynamodbattribute.MarshalMap(document)

	if err != nil {
		return err
	}

	svc := dynamodb.New(s.AwsSession)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(s.TableName),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}
