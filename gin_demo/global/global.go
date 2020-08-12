package global

import (
	"learn_tools/gin_demo/conf"
	"learn_tools/gin_demo/library/database/orm"

	"github.com/jinzhu/copier"
)

var (
	My *orm.MyORM
)

func Init() {
	My = orm.NewMySQL(conf.Conf.DB)
}

func Close() {
	if My != nil {
		My.Close()
	}

}

func Copy(toValue interface{}, fromValue interface{}) error {
	return copier.Copy(toValue, fromValue)
}
