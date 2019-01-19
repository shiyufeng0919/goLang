package database

/*
访问数据库

GO语言没有内置的驱动支持任何的数据库，但是定义了database/sql接口。
用户可以基于驱动接口开发相应数据库驱动

1。database/sql接口

(1)sql.Register

这个存在于database/sql的函数是用来注册数据库驱动的。

(2)driver.Driver :数据库驱动接口，返回一个DB的conn连接(这个conn只能应用在一个goroutine里，不能用在多个goroutine里)

(3)driver.Conn:数据库连接的接口定义(这个conn只能应用在一个goroutine里，不能用在多个goroutine里)

(4)driver.Stmt:Stmt是一种准备好的状态，和Conn相关联，而且只能应用于一个goroutine不能应用于多个goroutine

(5)driver.Tx:事务处理一般就两个过程，递交和回滚

(6)driver.Execer:conn可选择实现的接口

(7)driver.Result:执行update/insert等操作返回的结果定义

(8)driver.Rows:执行查询返回的结果集接口定义

(9)driver.RowsAffected:RowsAffected实现为一个int64的别名，但是它实现了Result接口

(10)driver.Value:Value就是一空接口，它可容纳任何数据

(11)driver.ValueConverter:定义了如何把一个普通值转化成driver.value接口

(12)driver.Valuer:Valuer接口定义了返回一个driver.Value的方式

(13)database/sql :在dtabase/sql/driver提供的接口基础上定义了一些更高阶的方法，用以简化DB操作，同时内部还建议性地实现一个conn pool
*/

/*
数据库：

#####SQLite数据库：
是一个开源的嵌入式关系数据库。实现自包容，零配置，支持事务的sql数据库引擎。

其特点是高度便携、使用方便、结构紧凑、高效、可靠。与其他数据库管理系统不同，SQLite的安装和运行非常简单，
在大多数情况下，只要确保SQLite的二进制文件存在，即可开始创建、连接和使用数据库。

如果你正在寻找一个嵌入式数据库项目或解决方案，SQLite是绝对值得考虑。SQLite可以说是开源的Access。


#####PostgreSQL数据库：
PostgreSQL 是一个自由的对象-关系数据库服务器（数据库管理系统），它在灵活的 BSD-风格许可证下发行。
它提供了相对其他开放源代码数据库系统（比如 MySQL和 Firebird），以及对专有系统（比如 Oracle、Sybase、IBM 的 DB2 和 Microsoft SQL Server）的一种选择。

PostgreSQL和MySQL比较，更加庞大，因为它是为替代Oracle而设计的。所以在企业应用中采用PostgreSQL是一个明智的选择。
现在MySQL被Oracle收购之后，有传闻Oracle正在逐步的封闭MySQL，鉴于此，将来我们也许会选择PostgreSQL而不是MySQL作为项目的后端数据库。

#####beedb库
使用beedb库进行ORM开发:http://wiki.jikexueyuan.com/project/go-web-programming/05.5.html

#####NoSql数据库
流行的Nosql主要有：redis / mongoDB / Cassandra / Membase （这些DB均具有高性能，高并发读写等特点）

Redis数据库
redis是一个key-value存储系统，和Memcached类似，它支持存储的value类型相对更多，包括string（字符串）、list（链表）、set（集合）和zset（有序集合）

Mongodb数据库
MongoDB 是一个高性能、开源、无模式的文档型数据库，是介于关系数据库和非关系数据库之间的产品，在非关系数据库当中功能最丰富，又最像关系数据库.
它支持的数据结构非常松散，采用类似json的bjson格式来存储数据，因此可以存储比较复杂的数据类型。
*/
