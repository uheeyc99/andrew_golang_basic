package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func test(){
	//fmt.Println("aaaq")
	db,err:=sql.Open("mysql","aiden:123456@tcp(10.10.3.227)/andrew?charset=utf8")
	if err !=nil{
		fmt.Printf("Open database error: %s\n",err)
	}
	defer db.Close()
	err=db.Ping()
	if err!=nil{
		fmt.Println("Ping err",err)
	}


}

func main(){
	for i:=0;i<1000000;i++{
		test()
	}
	fmt.Println("finished")

}
