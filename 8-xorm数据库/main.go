package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

func main() {

	//xorm

	//1. 创建数据库引擎对象 // orcale  sqlite3
	// root 为用户名
	// : 分割用户名和密码
	// password
	// @/ 固定格式
	// elmcms 要连接的数据库名

	engine, err := xorm.NewEngine("mysql", "root:yu271400@/testCms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	//2. 数据库引擎关闭
	defer engine.Close()

	engine.Logger().Info(" 连接数据库成功")

	//数据库引擎设置
	engine.ShowSQL(true)                     //设置显示SQL语句
	engine.Logger().SetLevel(core.LOG_DEBUG) //设置日志级别
	engine.SetMaxOpenConns(10)               //设置最大连接数
	//engine.SetMaxIdleConns(2)
	//将结构体定义自动映射成为数据库表格
	//engine.Sync(new(Person))
	engine.Sync2(new(Person))

	//engine.SetMapper(core.SnakeMapper{}) //驼峰式规则映射
	//engine.Sync2(new(QianfengTable))

	//完全一致映射规则
	engine.SetMapper(core.SameMapper{})
	engine.Sync2(new(PersonTable))

	//core.SnakeMapper{}
	//core.SameMapper{}
	//core.GonicMapper{}

	//查询表的所有数据
	session := engine.Table("user")
	count, err := session.Count()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(count)

	//使用原生sql语句进行查询
	result, err := engine.Query("select  * from user")
	if err != nil {
		panic(err.Error())
	}
	for key, value := range result {
		fmt.Println(key, value)
	}
}

type Person struct {
	Age  int
	Name string
}

type QianfengTable struct {
	QianfengName    string
	QianfengPhone   string
	QianfengAddress string
}

type PersonTable struct {
	PersonName string
	PersonAge  int
	PersonSex  int
}

/**
 *
 */
func OrmMapping() {

}
