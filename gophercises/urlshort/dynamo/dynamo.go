package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

// PutRecord inserts a new url shortening
func PutRecord(svc *dynamodb.DynamoDB, key string, value string) bool {
	record := Record{ShortURL: key, URL: value}
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"shorturl": {
				S: aws.String(record.ShortURL),
			},
			"url": {
				S: aws.String(record.URL),
			},
		},
		TableName:              aws.String("urlshortener"),
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				fmt.Println(dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return false
	}
	return true
}
