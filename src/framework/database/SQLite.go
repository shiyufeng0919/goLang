package database

import (
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

/*
###使用SQLite数据库:
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.3.md

###SQLite菜鸟教程：http://www.runoob.com/sqlite/sqlite-attach-database.html

https://github.com/mattn/go-sqlite3 支持database/sql接口，
基于cgo(关于cgo的知识请参看官方文档或者本书后面的章节)写的

SQLite 是一个开源的嵌入式关系数据库，实现自包容、零配置、支持事务的SQL数据库引擎。其特点是高度便携、使用方便、结构紧凑、高效、可靠。
与其他数据库管理系统不同，SQLite 的安装和运行非常简单，在大多数情况下,只要确保SQLite的二进制文件存在即可开始创建、连接和使用数据库。
如果您正在寻找一个嵌入式数据库项目或解决方案，SQLite是绝对值得考虑。SQLite可以说是开源的Access。

Step1:go get https://github.com/mattn/go-sqlite3
Step2:创建表结构
命令行：
1.$sqlite3  //进入sqlite3
2.附加数据库：
sqlite> attach database sqlite3.db as 'sqlite3';
3.查看所有database
sqlite> .database
4.create table userinfo and userdetail
CREATE TABLE sqlite3.userinfo (
	`uid` INTEGER PRIMARY KEY AUTOINCREMENT,
	`username` VARCHAR(64) NULL,
	`department` VARCHAR(64) NULL,
	`created` DATE NULL
);
CREATE TABLE sqlite3.userdetail (
	`uid` INT(10) NULL,
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
);
5.查看表结构
$.schema userinfo
6.查看所有表
$.tables
7.执行
  SQLiteDemo()
8.格式化输出格式
sqlite> .header on    ##目的显示表头
sqlite> .mode column
sqlite> select * from userinfo; //查询表数据
9.退出sqlite
.exit
*/
func SQLiteDemo(){
	/*
	打开sqlite3的驱动，注:
	1.sqlite3.db必须在src目录下
	2.sqlite3.db需要附加到sqlite数据库中：attach database sqlite3.db as 'sqlite3';
	  并在指定数据库(别名)下创建表结构
	*/
	//db,err:=sql.Open("sqlite3","sqlite3.db") #与下方等价
	db,err:=sql.Open("sqlite3","./sqlite3.db")
	checkErr(err)

	//插入数据
	stmt,err:=db.Prepare("insert into userinfo(username,department,created) values (?,?,?)")
	checkErr(err)

	res,err:=stmt.Exec("yufeng","blockchain dept","2017-10-23")
	checkErr(err)
	res,err=stmt.Exec("kaixin","blockchain","2017-10-23")
	checkErr(err)

	id,err:=res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt,err=db.Prepare("update userinfo set department=? where uid=?")
	checkErr(err)

	res,err=stmt.Exec("baas",id)
	checkErr(err)
	affect,err:=res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows,err:=db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next(){
		var uid int
		var username string
		var dept string
		var created time.Time

		err=rows.Scan(&uid,&username,&dept,&created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(dept)
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

/*
实际项目应用：fabric-ca使用sqlite数据库
*/


