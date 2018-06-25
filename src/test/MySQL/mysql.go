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
//var string cut = "CREATE TABLE IF NOT EXISTS tasks (task_id INT(11) NOT NULL AUTO_INCREMENT,subject VARCHAR(45) DEFAULT NULL,start_date DATE DEFAULT NULL,end_date DATE DEFAULT NULL,description VARCHAR(200) DEFAULT NULL,PRIMARY KEY (task_id) ) ENGINE=InnoDB;"

func create_usr_tbl(){
	db,err:=sql.Open("mysql","aiden:123456@tcp(10.10.3.227)/andrew?charset=utf8")
	if err !=nil{
		fmt.Printf("Open database error: %s\n",err)
	}
	defer db.Close()
	cut := "CREATE TABLE IF NOT EXISTS user (" +
		"user_id INT(11) NOT NULL AUTO_INCREMENT," +
		"username VARCHAR(45) DEFAULT NULL," +
		"password VARCHAR(45) DEFAULT NULL," +
		"last_login DATE DEFAULT NULL," +
		"email VARCHAR(45) DEFAULT NULL," +
		"description VARCHAR(200) DEFAULT NULL," +
		"PRIMARY KEY (user_id) ) ENGINE=InnoDB;"
	fmt.Println(cut)
	_,err=db.Query(cut)
	if(err!=nil){
		fmt.Println(err)
	}
}

func add_usr(username string,password string) (resault string) {
	db,err:=sql.Open("mysql","aiden:123456@tcp(10.10.3.227)/andrew?charset=utf8")
	if err !=nil{
		fmt.Printf("Open database error: %s\n",err)
	}
	defer db.Close()

	var user_id int
	err=db.QueryRow("select user_id from user where username=?",username).Scan(&user_id)
	if(err!=nil){
		fmt.Println(err)
		stmt,err:=db.Prepare("insert into user SET username=?,password=?")
		if(err!=nil){
			fmt.Println(err)
			return
		}
		defer stmt.Close()

		res,err:=stmt.Exec(username,password)
		if(err!=nil){
			fmt.Println(err)
		}
		fmt.Println(res.LastInsertId())
		resault="user created "
	}else{
		fmt.Println(user_id)
		resault="user already exist"
	}

	return resault
}

func check_user(username string,password string)  (result string){
	db,err:=sql.Open("mysql","aiden:123456@tcp(10.10.3.227)/andrew?charset=utf8")
	if err !=nil{
		fmt.Printf("Open database error: %s\n",err)
	}
	defer db.Close()
	var user_id int
	var pass_word string
	err=db.QueryRow("select password,user_id from user where username=?;",username).Scan(&pass_word,&user_id)
	if(err!=nil){
		result="user not exist !"
	}else
	if(password!=pass_word){
		result="password err"
	}else{
		result="login ok"
	}
	return result
}

func test2(){
	create_usr_tbl()
	res:=add_usr("aiden2","123456")
	fmt.Println(res)
	res=check_user("aiden2","12345678")
	fmt.Println(res)

}


func main(){

	test2()
	fmt.Println("finished")

}
