package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
type Userinfo struct {
	Uid        int `orm:"pk"` // 如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid     int `orm:"pk"`
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"` // 设置一对多关系
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "mariadb:mariadb@tcp(localhost:3306)/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Post), new(Tag))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	// insert
	var userinfo Userinfo
	userinfo.Username = "zxxx"
	userinfo.Departname = "zxxx"

	id, err := o.Insert(&userinfo)
	if err == nil {
		fmt.Println(id, err)
	}
}
