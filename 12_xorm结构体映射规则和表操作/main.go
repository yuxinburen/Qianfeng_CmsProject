package main

import (
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"fmt"
)

func main() {
	//1. 创建数据库引擎对象
	engine, err := xorm.NewEngine("mysql", "root:yu271400@/testCms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	//设置名称映射规则：驼峰式命名映射规则
	engine.SetMapper(core.SnakeMapper{})
	engine.Sync2(new(UserTable))

	//engine.SetMapper(core.SameMapper{})
	engine.Sync2(new(StudentTable))

	//对特定词支持性更友好
	//engine.SetMapper(core.GonicMapper{})
	engine.Sync2(new(PersonTable))

	personEmpty, err := engine.IsTableEmpty(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personEmpty {
		fmt.Println(" 人员表是空的 ")
	} else {
		fmt.Println(" 人员表不为空 ")
	}

	//判断表结构是否存在
	studentExist, err := engine.IsTableExist(new(StudentTable))
	if err != nil {
		panic(err.Error())
	}
	if studentExist {
		fmt.Println("学生表存在")
	} else {
		fmt.Println("学生表不存在")
	}

}

type UserTable struct {
	UserId   int64  `xorm:"pk autoincr"`
	UserName string `xorm:"varchar(32)"` //用户名
	UserAge  int64  `xorm:"default 1"`   //用户年龄
	UserSex  int64  `xorm:"default 0"`   //用户性别
}

/**
 * 学生表
 */
type StudentTable struct {
	Id          int64  `xorm:"pk autoincr"` //主键 自增
	StudentName string `xorm:"varchar(24)"` //
	StudentAge  int    `xorm:"int default 0"`
	StudentSex  int    `xorm:"index"` //sex为索引
}

/**
 * 人员结构表
 */
type PersonTable struct {
	ID         int64     `xorm:"pk autoincr"`   //主键自增  id
	PersonName string    `xorm:"varchar(24)"`   //可变字符
	PersonAge  int       `xorm:"int default 0"` //默认值
	PersonSex  int       `xorm:"notnull"`       //不能为空
	City       CityTable `xorm:"-"`             //不映射该字段
}

type CityTable struct {
	CityName      string
	CityLongitude float32
	CityLatitude  float32
}
