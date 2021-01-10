package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/brianvoe/gofakeit"
)

// Code derived from AWS examples: https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#pkg-examples

type dynamo struct {
	*dynamodb.DynamoDB
}

func NewDynamoDBSession(endpoint, region string) *dynamo {
	cfg := &aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(region), // this should be same as in the ~/.aws/config
	}

	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess, cfg)

	return &dynamo{DynamoDB: db}
}

func main() {
	db := NewDynamoDBSession("http://localhost:8000", "us-west")
	gofakeit.Seed(0)

	result, err := db.createTable()

	if err != nil {
		log.Println("create table err:", err)
	}

	log.Println(result)

	input := &dynamodb.ListTablesInput{}
	tables, err := db.ListTables(input)

	if err := handleErr(err); err != nil {
		log.Println("list table error: ", err)
		return
	}

	log.Println("tables:", tables)
	email := gofakeit.Email()

	_, err = db.putItem(email, gofakeit.Name())

	if err := handleErr(err); err != nil {
		log.Println("put item error: ", err)
		return
	}

	item, err := db.getItem(email)

	if err := handleErr(err); err != nil {
		log.Println("get item error: ", err)
		return
	}

	log.Println(item)
}

func (db *dynamo) createTable() (*dynamodb.CreateTableOutput, error) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Email"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Email"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("Contacts"),
	}

	return db.CreateTable(input)
}

func (db *dynamo) putItem(email string, name string) (*dynamodb.PutItemOutput, error) {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
			"Name": {
				S: aws.String(name),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("Contacts"),
	}

	return db.PutItem(input)
}

func (db *dynamo) getItem(email string) (*dynamodb.GetItemOutput, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String("Contacts"),
	}

	return db.GetItem(input)
}

func handleErr(err error) error {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				return fmt.Errorf(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			case dynamodb.ErrCodeLimitExceededException:
				return fmt.Errorf(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return fmt.Errorf(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return fmt.Errorf(aerr.Error())
			}
		}
	}

	return nil
}
