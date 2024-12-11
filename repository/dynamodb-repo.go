package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/clrajapaksha/to-do-list-app/entities"
)

type dynamoDBRepo struct {
	tableName string
}

func NewDynamoDBRepository() TaskRepository {
	return &dynamoDBRepo{
		tableName: "Task",
	}
}

func createDynamoDBClient() *dynamodb.DynamoDB {

	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
			},
		),
	)
	return dynamodb.New(sess)

}

func (repo *dynamoDBRepo) Save(task *entities.Task) (*entities.Task, error) {
	dynamoDBClient := createDynamoDBClient()
	attributeValue, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		return nil, err
	}

	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (repo *dynamoDBRepo) FindAll() ([]entities.Task, error) {
	dynamoDBClient := createDynamoDBClient()

	params := &dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}
	result, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}
	var tasks []entities.Task = []entities.Task{}
	for _, i := range result.Items {
		task := entities.Task{}
		err = dynamodbattribute.UnmarshalMap(i, &task)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *dynamoDBRepo) FindByID(id string) (*entities.Task, error) {
	dynamoDBClient := createDynamoDBClient()

	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	task := entities.Task{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &task)
	if err != nil {
		panic(err)
	}
	return &task, nil

}

func (repo *dynamoDBRepo) Delete(task *entities.Task) error {
	return nil
}
