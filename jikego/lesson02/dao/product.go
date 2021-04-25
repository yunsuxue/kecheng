package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//db发生错误应该调用方处理更符合逻辑, 此处可以使用 go 1.13新加的%w, 加上错误信息,让链路清晰, 或者使用 pkg/errors 包包装
func UpdateProduct() (bool, error) {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/golab?charset=utf8")
	if err != nil {
		return false, fmt.Errorf(" db open error,UpdateProduct fail:%w", err)
	}
	defer db.Close()

	//...
	return true, nil
}
