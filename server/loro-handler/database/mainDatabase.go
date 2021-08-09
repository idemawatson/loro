package database

import (
	"github.com/guregu/dynamo"
)

var sharedInstance dynamo.Table = newMainTableAccess()
var TABLENAME = "MAIN_TABLE"

func newMainTableAccess() dynamo.Table {
	return GetTable(TABLENAME)
}

func GetMainTable() dynamo.Table {
	return sharedInstance
}
