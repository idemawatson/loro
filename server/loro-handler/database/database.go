package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"

	"loro-handler/config"
)

func GetTable(tableName string) *dynamoClient {
	cfg := aws.NewConfig().WithRegion(config.GetEnv("REGION"))
	db := dynamo.New(session.Must(session.NewSession()), cfg)
	return &dynamoClient{table: db.Table(config.GetEnv(tableName))}
}

type dynamoClient struct {
	table DynamoClient
}

type DynamoClient interface {
	Get(name string, value interface{}) QueryIface
}

type Query struct {
	dynamo.Query
}

type QueryIface interface {
	Range(name string, op dynamo.Operator, values ...interface{}) QueryIface
	Index(name string) QueryIface
	One(out interface{}) error
}
