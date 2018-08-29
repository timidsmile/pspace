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

	if Db, err = gorm.Open("mysql", dbCfg); err != nil {
		return err
	}

	if cfg.Debug {
		Db.LogMode(true)
	}

	return nil
}
