package main

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * xorm基本操作和高级操作
 */
func main() {

	//1. 创建数据库引擎对象
	engine, err := xorm.NewEngine("mysql", "root:yu271400@/testCms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	//2.设置映射规则
	engine.SetMapper(core.SnakeMapper{})

	//3、同步数据库表格
	engine.Sync2(new(PersonTable))

	//4.判断person表格是否存在
	personExist, err := engine.IsTableExist(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personExist {
		fmt.Println("人员表存在")
	} else {
		fmt.Println("人员表不存在")
	}

	//5.判断person表格是否为空
	personEmpty, err := engine.IsTableEmpty(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personEmpty {
		fmt.Println(" 人员表是空的 ")
	} else {
		fmt.Println(" 人员表不为空 ")
	}

	//二、条件查询
	//1.ID查询
	var person PersonTable
	// select * from person_table where id = 1
	result, err := engine.Id(1).Get(&person)
	fmt.Println(result, err)
	fmt.Println(person.PersonName)
	fmt.Println()

	//2.where多条件查询
	//var person1 PersonTable
	var person2 []PersonTable
	// select * from person_table where person_age = 26 and person_sex = 2
	engine.Where(" person_age = ? and person_sex = ?", 26, 2).Find(&person2)
	fmt.Println(person2)
	fmt.Println()

	//3.And条件查询
	var persons []PersonTable
	//select * from person_table where person_age = 26 and person_sex = 2
	err = engine.Where(" person_age = ? ", 26).And("person_sex = ? ", 2).Find(&persons)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(persons)
	fmt.Println()

	//4、Or条件查询
	var personArr []PersonTable
	//select * from person_table where person_age = 26 or person_sex = 1
	err = engine.Where(" person_age = ? ", 26).Or("person_sex = ? ", 1).Find(&personArr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(personArr)
	fmt.Println()

	//5、原生SQL语句查询支持 like语法
	var personsNative []PersonTable //模糊查询
	err = engine.SQL(" select * from person_table where person_name like 't%' ").Find(&personsNative)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(personsNative)
	fmt.Println()

	////6、排序条件查询
	var personsOrderBy []PersonTable
	//select * from person_table orderby person_age  升序排列
	//engine.OrderBy(" person_age ").Find(&personsOrderBy)
	engine.OrderBy(" person_age desc ").Find(&personsOrderBy)
	fmt.Println(personsOrderBy)
	fmt.Println()

	////7、查询特定字段
	var personsCols []PersonTable
	engine.Cols("person_name", "person_age").Find(&personsCols)
	for _, col := range personsCols {
		fmt.Println(col)
	}

	//
	////三、增加记录操作
	//personInsert := PersonTable{
	//	PersonName: "Hello",
	//	PersonAge:  18,
	//	PersonSex:  1,
	//}
	//rowNum, err := engine.Insert(&personInsert)
	//fmt.Println(rowNum) //rowNum 受影响的记录条数
	//fmt.Println()

	////四、删除操作
	//rowNum, err := engine.Delete(&personInsert)
	//fmt.Println(rowNum) //rowNum 受影响的记录条数
	//fmt.Println()

	//
	////五、更新操作
	//rowNum, err := engine.Id(9).Update(&personInsert)
	//fmt.Println(rowNum) //rowNum 受影响的记录条数
	//fmt.Println()


	////六、统计功能count
	//count, err := engine.Count(new(PersonTable))
	//fmt.Println("PersonTable表总记录条数：", count)


	////七、事务操作
	personsArray := []PersonTable{
		PersonTable{
			PersonName: "Jack",
			PersonAge:  28,
			PersonSex:  1,
		},
		PersonTable{
			PersonName: "Mali",
			PersonAge:  28,
			PersonSex:  1,
		},
		PersonTable{
			PersonName: "Ruby",
			PersonAge:  28,
			PersonSex:  1,
		},
	}

	session := engine.NewSession()
	session.Begin()
	for i := 0; i < len(personsArray); i++ {
		_, err = session.Insert(personsArray[i])
		if err != nil {
			session.Rollback()
			session.Close()
		}
	}
	err = session.Commit()
	session.Close()
	if err != nil {
		panic(err.Error())
	}

}

/**
 * 人员结构表
 */
type PersonTable struct {
	Id         int64     `xorm:"pk autoincr"`   //主键自增
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
