package appctx

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type appContext struct {
	db *sql.DB
}

type AppContext interface {
	GetMainDBConnection() *sql.DB
}

func NewAppContext(db *sql.DB) *appContext{
	return &appContext{
		db: db,
	}
}

func (ctx *appContext) GetMainDBConnection() *sql.DB {
	database, err := InitMysqlDB("sql6.freemysqlhosting.net", 3306, "sql6501022", "szkJpYD8aq", "sql6501022")
	if err != nil {
		panic(err)
	}

	return database
}

func InitMysqlDB(host string, port int, username, passsword, database string) (*sql.DB, error) {
	hostPort := fmt.Sprintf("%v:%d", host, port)
	connectString := dsn(database, username, passsword, hostPort)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}

// dsn connect to db
func dsn(dbName, username, password, hostname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
