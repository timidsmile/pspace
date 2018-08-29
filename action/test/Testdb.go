package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // mysql 驱动
	"github.com/timidsmile/pspace/components"
	"github.com/timidsmile/pspace/model"
)

func TestdbAction(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "master!"
	}

	// 获取通用数据库对象`*sql.DB`以使用其函数
	mysqlDB := components.Db.DB()

	// Ping
	mysqlDB.Ping()

	// 指定连接池数目
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)

	db := components.Db

	has1 := db.HasTable(&model.UserBasic{}) // 检查模型`UserBasic`表是否存在
	has2 := db.HasTable("user_third_basic") // 检查表`users_basic`是否存在

	if has1 == false && has2 == false {
		// 新建表
		db.CreateTable(&model.UserBasic{})
		db.CreateTable(&model.UserThirdBasic{})
	}

	if has1, has2 := db.HasTable(&model.UserBasic{}), db.HasTable("user_third_basic"); has1 == false && has2 == false {
		// 新建表
		db.CreateTable(&model.UserBasic{})
		db.CreateTable(&model.UserThirdBasic{})
	}

	// -----------------------------------------------插入-----------------------------------------------

	/*user1 := &model.UserBasic{
		UserID:    123,
		UserName:  "jinzhu",
		Mobile:    "test",
		Email:     "test",
		Passwd:    "test",
		NickName:  "test",
		AvatarUrl: "test",
		Status:    1,
	}

	user2 := &model.UserBasic{
		UserID:    124,
		UserName:  "jinzhu2",
		Mobile:    "test",
		Email:     "test",
		Passwd:    "test",
		NickName:  "test",
		AvatarUrl: "test",
	}

	components.Db.Create(&user1)
	components.Db.Create(&user2)
	*/
	// -----------------------------------------------查询-----------------------------------------------
	user := &model.UserBasic{}
	/*
			// 获取第一条记录，按主键排序
			db.First(&user)
			// SELECT * FROM user_basic ORDER BY id LIMIT 1;

			// 获取最后一条记录，按主键排序
			db.Last(&user)
			// SELECT * FROM user_basic ORDER BY id DESC LIMIT 1;

			// 获取所有记录
			db.Find(&user)
			// SELECT * FROM user_basic;

			// 使用主键获取记录
			db.First(&user, 10)
			// SELECT * FROM user_basic WHERE id = 10;

			// 获取第一个匹配记录
			db.Where("name = ?", "jinzhu").First(&user)
			// SELECT * FROM user_basic WHERE name = 'jinzhu' limit 1;

			// 获取所有匹配记录
			db.Where("name = ?", "jinzhu").Find(&user)
			// SELECT * FROM user_basic WHERE name = 'jinzhu';

			db.Where("user_name <> ?", "jinzhu").Find(&user)

			// IN
			db.Where("user_name in (?)", []string{"jinzhu", "jinzhu2"}).Find(&user)

			// LIKE
			db.Where("user_name LIKE ?", "%jin%").Find(&user)

			// AND
			db.Where("user_name = ? AND user_id >= ?", "jinzhu", 120).Find(&user)

			// Struct
			db.Where(&model.UserBasic{UserName: "jinzhu", NickName: "jz"}).First(&user)
			// SELECT * FROM user_basic WHERE name = "jinzhu" AND nick_name = 'jz' LIMIT 1;

		// Map
		db.Where(map[string]interface{}{"user_name": "jinzhu", "nick_name": "jz"}).Find(&user)
		// SELECT * FROM user_basic WHERE name = "jinzhu" AND age = 20;

		// 主键的Slice
		db.Where([]int64{1, 2, 3}).Find(&user)
		// SELECT * FROM user_basic WHERE id IN (1, 2, 3);

		db.Not("user_name", "jinzhu").First(&user)
		// SELECT * FROM user_basic WHERE name <> "jinzhu" LIMIT 1;

		// Not In
		db.Not("user_name", []string{"jinzhu", "jinzhu2"}).Find(&user)
		// SELECT * FROM user_basic WHERE name NOT IN ("jinzhu", "jinzhu2");

		// Not In slice of primary keys
		db.Not([]int64{1, 2, 3}).First(&user)
		// SELECT * FROM user_basic WHERE id NOT IN (1,2,3);

		db.Not([]int64{}).First(&user)
		// SELECT * FROM user_basic;

		// Plain SQL
		db.Not("name = ?", "jinzhu").First(&user)
		// SELECT * FROM user_basic WHERE NOT(name = "jinzhu");

		// Struct
		db.Not(model.UserBasic{UserName: "jinzhu"}).First(&user)
		// SELECT * FROM user_basic WHERE name <> "jinzhu";

		// or 条件查询
		db.Where("id = ?", 1).Or("id = ?", 2).Find(&user)
		// SELECT * FROM user_basic WHERE role = 'admin' OR role = 'super_admin';

		// Struct
		db.Where("name = 'jinzhu'").Or(model.UserBasic{UserName: "jinzhu2"}).Find(&user)
		// SELECT * FROM user_basic WHERE user_name = 'jinzhu' OR user_name = 'jinzhu2';

		// Map
		db.Where("user_name = 'jinzhu'").Or(map[string]interface{}{"user_name": "jinzhu 2"}).Find(&user)

		// 使用select 指定要从数据库检索的字段，默认情况下，将选择所有字段;
		db.Select("user_name, nick_name").Find(&user)
		// SELECT user_name, nick_name FROM user_basic;

		db.Select([]string{"user_name", "nick_name"}).Find(&user)
		// SELECT user_name, nick_name FROM user_basic;

		// order
		db.Order("user_id desc, id").Find(&user)
		// SELECT * FROM user_basic ORDER BY age desc, name;
	*/

	// Multiple orders
	db.Order("id desc").Order("user_id").Find(&user)
	// SELECT * FROM user_basic ORDER BY age desc, name;
	return

	// -----------------------------------------------删除表----------------------------------------------

	// 删除模型`UserBasic`的表
	db.DropTable(&model.UserBasic{})

	// 删除表`user_third_basic`
	db.DropTable("user_third_basic")

	// 删除模型`UserBasic`的表和表`user_third_basic`
	db.DropTableIfExists(&model.UserBasic{}, "user_third_basic")

	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("welcome %s !\n", value)))
	return
}
