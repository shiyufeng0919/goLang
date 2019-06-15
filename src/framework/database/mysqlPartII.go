package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

/*
Baas应用：https://github.com/jmoiron/sqlx

go get github.com/jmoiron/sqlx
*/
var schema=`
CREATE TABLE person(
  first_name text,
  last_name text,
  email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email	string
}
type Place struct {
	Country string
	City sql.NullString
	TelCode int
}

func JmoironSqlxDemo(){
	db,err:=sqlx.Connect("mysql","root:shiyufeng@/mydb?charset=utf8")
	if err !=nil{
		fmt.Println(err)
	}

	db.MustExec(schema)

	tx:=db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)","yufeng","kaixin","yufeng@sina.com")
    tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)","Beijing","010")

	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)")
	tx.Commit()

	people:=[]Person{}
	db.Select(&people,"SELECT * FROM person ORDER BY first_name ASC")
	for _,value:=range people{
		fmt.Println(value)
	 }
	}