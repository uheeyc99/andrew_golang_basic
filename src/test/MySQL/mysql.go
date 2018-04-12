package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"time"
)

func test(){
	fmt.Println("a: " + time.Now().String())
	db,err:=sql.Open("mysql","aiden:123456@tcp(10.10.3.227)/andrew?charset=utf8")
	if err !=nil{
		fmt.Printf("Open database error: %s\n",err)
	}

	fmt.Println("b: " + time.Now().String())

	defer db.Close()
	err=db.Ping()
	fmt.Println("c: " + time.Now().String())

	if err!=nil{
		fmt.Println("Ping err",err)
	}else{
		fmt.Println("pong " + time.Now().String())
	}


}

func main(){
	for i:=0;i<1000000;i++{
		test()
	}
	fmt.Println("finished")

}
