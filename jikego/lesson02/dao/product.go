package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
)

var codePath string
var codeLine int

//db发生错误应该调用方处理更符合逻辑, 此处可以使用 go 1.13新加的%w, 加上错误信息,让链路清晰, 或者使用 pkg/errors 包包装
func UpdateProduct() (sql.Result, error) {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/golab?charset=utf8")
	if err != nil {
		_, codePath, codeLine, _ = runtime.Caller(0)
		return nil, fmt.Errorf(" db open error,UpdateProduct fail,path[%s],line[%d]:%w", codePath, codeLine, err)
	}
	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	if err != nil {
		//可以加上文件和行号信息
		_, codePath, codeLine, _ = runtime.Caller(0)
		return nil, fmt.Errorf(" sql Prepare fail,path[%s],line[%d]:%w", codePath, codeLine, err)
	}

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	if err != nil {
		//可以加上文件和行号信息
		_, codePath, codeLine, _ = runtime.Caller(0)
		return nil, fmt.Errorf(" sql Exec fail,path[%s],line[%d]:%w", codePath, codeLine, err)
	}

	//...
	return res, nil
}
