package main

import (
	"fmt"
	"lession02/dao"
	"log"
)

func main() {
	//错误应该调用方处理, 除非错误无关紧要,记下日志,逻辑继续执行,db错误应该是重要错误,调用方需要处理
	ok, err := dao.UpdateProduct()
	if err != nil {
		log.Fatal("db 执行失败, 错误, %v", err)
	}

	if ok {
		//...
		fmt.Println("db ok")
	}
}