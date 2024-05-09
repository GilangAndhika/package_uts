package package_uts

import (
	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "ats",
}

var MongoConn = atdb.MongoConnect(MongoInfo)