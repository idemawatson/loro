package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"

	"loro-handler/config"
)

func GetTable(tableName string) dynamo.Table {
	cfg := aws.NewConfig().WithRegion(config.GetEnv("REGION"))
	db := dynamo.New(session.Must(session.NewSession()), cfg)
	return db.Table(config.GetEnv(tableName))
}
