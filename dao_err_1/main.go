package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type data struct {
	info string
}

func main() {
	_, err := dao()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("查询无结果 reason: %v", err.Error())
	}
	// do something
}

func dao() (data, error) {
	return data{}, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("sql: %s, err: %v", "custome err info", "err info"))
}
