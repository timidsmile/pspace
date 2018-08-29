package components

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/timidsmile/pspace/setting"
)

var (
	Db *gorm.DB
)

const dbTemplate = "%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

func dsn(db setting.DbConfig) string {
	return fmt.Sprintf(dbTemplate, db.User, db.Host, db.Port, db.Name)
}

func InitDb(cfg *setting.Config) error {
	var err error
	dbCfg := dsn(cfg.PspaceDb)

	fmt.Println(dbCfg)

	if Db, err = gorm.Open("mysql", dbCfg); err != nil {
		return err
	}

	// 全局开启日志
	Db.LogMode(true)

	// 全局禁用表名复数
	Db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	return nil
}
