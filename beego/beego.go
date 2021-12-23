package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "mariadb:mariadb@tcp(localhost:3306)/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "test"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "berny"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
