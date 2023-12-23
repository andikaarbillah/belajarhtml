package setupDB

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "26527y25tw"
	dbname = "test"
)

var psqlInfo = fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable",host,port,user,password,dbname)

func ConnectDb() *sql.DB{
	db, err := sql.Open(user,psqlInfo)
	if err != nil{
		panic(err)
	}
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	return db
}



