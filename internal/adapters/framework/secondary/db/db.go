package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) *Adapter {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatalf("db ping failed: %v", errPing)
	}

	return &Adapter{db: db}
}

func (dbAdapter Adapter) CloseDbConnection() {
	err := dbAdapter.db.Close()
	if err != nil {
		log.Fatalf("db close failed: %v", err)
	}
}

func (dbAdapter Adapter) AddToHistory(res int32, operation string) error {
	queryString, args, err := sq.
		Insert("arithmetic_history").
		Columns("date", "answer", "operation").
		Values(time.Now(), res, operation).
		ToSql()
	if err != nil {
		fmt.Printf("db close failed: %v\n", err)
		return err
	}

	_, errExec := dbAdapter.db.Exec(queryString, args...)
	if errExec != nil {
		fmt.Printf("db query execution failed: %v\n", err)
		return errExec
	}

	return nil
}
