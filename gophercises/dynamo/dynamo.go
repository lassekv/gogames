package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Record A returned record
type Record struct {
	ShortURL string
	URL      string
}

// CreateClient Creates a client for quering dynamodb
func CreateClient() (svc *dynamodb.DynamoDB, ok bool) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewSharedCredentials("", "default")})
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		fmt.Println("unable to fetch credentials")
		ok = false
		return nil, ok
	}
	ok = true
	svc = dynamodb.New(sess)
	return svc, ok
}

// GetRecord returns one record with the given key
func GetRecord(svc *dynamodb.DynamoDB, key string) (Record, bool) {

	dynamoKey := map[string]*dynamodb.AttributeValue{
		"shorturl": &dynamodb.AttributeValue{
			S: aws.String(key),
		}}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("urlshortener"),
		Key:       dynamoKey,
	})

	if err != nil {
		fmt.Println(err.Error())
		return Record{}, false
	}

	record := Record{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &record)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return record, true
}
