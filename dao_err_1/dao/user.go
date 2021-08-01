package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id       int
	username string
}

func GetUser(id int) (Data, error) {
	db, err := sql.Open("mysql",
		"root:mypassword@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	data := Data{}
	row := db.QueryRow("select  from user where id=? limit 1", id)
	err = row.Scan(&data)
	if err == sql.ErrNoRows {
		return Data{}, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("sql: %s, err: %v", "custome err info", "err info"))
	}
	return data, nil
}
