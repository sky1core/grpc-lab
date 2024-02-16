package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return err // 연결 실패 시 에러 반환
	}

	return nil // 성공 시 nil 반환
}
