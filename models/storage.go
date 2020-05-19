package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

var (
	o orm.Ormer
)

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:dev@(127.0.0.1:3306)/auth_jwt?charset=utf8", 30)

	orm.RunCommand()

	o = orm.NewOrm()
}

func GetORM() orm.Ormer {
	return o
}
