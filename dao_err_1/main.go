package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go_study/dao_err_1/dao"
)


func main() {
	_, err := dao.GetUser(1)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("查询无结果 reason: %v", err.Error())
	}
	// do something
}


