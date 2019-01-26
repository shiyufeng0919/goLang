package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"strconv"
	"github.com/syndtr/goleveldb/leveldb/util"
)
/*
golang使用高性能的KV库leveldb
https://github.com/syndtr/goleveldb/blob/master/README.md

LevelDB是Google开源的持久化KV单机数据库，具有很高的随机写，顺序读/写性能，但是随机读的性能很一般。
也就是说，LevelDB很适合应用在查询较少，而写很多的场景。key和value都是任意长度的字节数组
entry（即一条K-V记录）默认是按照key的字典顺序存储的，当然开发者也可以重载这个排序函数。
提供的基本操作接口：Put()、Delete()、Get()、Batch()。支持批量操作以原子操作进行，
可以创建数据全景的snapshot(快照)，并允许在快照中查找数据。可以通过前向（或后向）迭代器遍历数据（迭代器会隐含的创建一个snapshot）
自动使用Snappy压缩数据。非关系型数据模型（NoSQL），不支持sql语句，也不支持索引。一次只允许一个进程访问一个特定的数据库
没有内置的C/S架构，但开发者可以使用LevelDB库自己封装一个server。

Step1:go get github.com/syndtr/goleveldb/leveldb
Step2:创建并打开DB，CRUD
*/
func LevelDB(){
	//创建并打开数据库(以项目golang为基准，创建一个目录为path/to/db的目录，并打开,写入)
	//db,err:=leveldb.OpenFile("path/to/db",nil)
	                  //golang/src/source/leveldb下存储leveldb
	db,err:=leveldb.OpenFile("src/source/leveldb",nil)
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	defer db.Close() //关闭数据库

	//写入5条数据
	db.Put([]byte("key1"),[]byte("value1"),nil)
	db.Put([]byte("key2"),[]byte("value2"),nil)
	db.Put([]byte("key3"),[]byte("value3"),nil)
	db.Put([]byte("key4"),[]byte("value4"),nil)
	db.Put([]byte("key5"),[]byte("value5"),nil)

	//循环遍历数据
	fmt.Println("循环遍历数据######")
	iter:=db.NewIterator(nil,nil)
	for iter.Next(){
		fmt.Printf("key:%s,value:%s\n",iter.Key(),iter.Value())
	}
	iter.Release()

	//读取某条数据
	fmt.Println("读取某条数据######")
	data,_:=db.Get([]byte("key2"),nil)
	fmt.Printf("key2:%s\n",data)

	//批量写入数据
	fmt.Println("批量写入数据######")
	batch:=new(leveldb.Batch)
	batch.Put([]byte("key6"),[]byte(strconv.Itoa(10000)))
	batch.Put([]byte("key7"),[]byte(strconv.Itoa(20000)))
	batch.Delete([]byte("key4"))
	db.Write(batch,nil)

	//查找数据
	fmt.Println("查找数据######")
	key:="key7"
	iter=db.NewIterator(nil,nil)
	for ok:=iter.Seek([]byte(key));ok;ok=iter.Next(){
		fmt.Printf("key:%s,value:%s\n",iter.Key(),iter.Value())
	}
	iter.Release()

	//按key范围遍历数据
	fmt.Println("按key范围遍历数据######")
	iter=db.NewIterator(&util.Range{Start:[]byte("key3"),Limit:[]byte("key7")},nil)
	for iter.Next(){
		fmt.Printf("key:%s,value:%s\n",iter.Key(),iter.Value())
	}
    iter.Release()
}