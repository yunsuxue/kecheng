package main

import (
	"fmt"
	"lession02/dao"
	"log"
)

func main() {
	//错误应该调用方处理, 除非错误无关紧要,记下日志,逻辑继续执行,db错误应该是重要错误,调用方需要处理
	result, err := dao.UpdateProduct()
	if err != nil {
		//如果查看路径信息,可以循环unwrap
		/*for err != nil {
			fmt.Println(err)
			err = errors.Unwrap(err)
		}*/
		log.Fatalf("db 执行失败, 错误: %v", err)
	}

	if result != nil {
		//...
		fmt.Println("db ok")
	}
}