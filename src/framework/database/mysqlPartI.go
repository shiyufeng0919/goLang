package database

import (
	_ "github.com/GO-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

/*
baas项目应用：https://github.com/jmoiron/sqlx
*/
/*
Go语言支持mysql驱动，常见:
1.https://github.com/GO-SQL-Driver/MySQL #支持database/sql，全部采用go语言编写
2.https://github.com/ziutek/mymysql #支持database/sql,也支持自定义接口，全部采用go语言编写
3.https://github.com/Philio/GoMySQL #不支持database/sql,自定义接口，全部采用go语言编写
*/
/*
https://github.com/GO-SQL-Driver/MySQL
1。拉取:go get -u github.com/go-sql-driver/mysql
2。创建表
create database mydb
use mydb
create table userinfo(uid int(10) not null auto_increment,username varchar(64) default null,created date default null,primary key(uid))
create table userdetail(uid int(10) not null default 0,intro text null,profile text null,primary key(uid))
*/
/*
##########示例1：使用database/sql接口对数据库表进行增删改查操作
*/
func MysqlDbDemo1(){
	//db, err := sql.Open("mysql", "user:password@/dbname")
	//sql.Open:打开一个注册过的数据库驱动
	db,err:=sql.Open("mysql","root:shiyufeng@/mydb?charset=utf8")
	checkErr(err)

	//1.插入数据
	//db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态
	//?形式，防止sql注入
	stmt,err:=db.Prepare("INSERT userinfo SET username=?,created=?")
	checkErr(err)

	//stmt.Exec()函数用来执行stmt准备好的sql语句
	res,err:=stmt.Exec("yufeng","2019-01-19") //传参数，与?号个数对应
	checkErr(err)

	id,err:=res.LastInsertId()
	checkErr(err)
	fmt.Println(id) //1

	//2。更新数据
	stmt,err=db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res,err=stmt.Exec("kaixin",id)
	checkErr(err)
	affect,err:=res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	//db.Query()函数用来直接执行sql返回rows结果
	rows,err:=db.Query("select * from userinfo")
	checkErr(err)
	for rows.Next(){
		var uid int
		var username string
		var created string

		err=rows.Scan(&uid,&username,&created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(created)
	}

	//删除数据
	stmt,err=db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res,err=stmt.Exec(id)
	checkErr(err)
	affect,err=res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}



/*
gosql - 基于 sqlx 封装的 Golang 数据库操作
https://javascript.ctolib.com/ilibs-gosql.html
*/