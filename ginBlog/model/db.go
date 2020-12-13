package model

import (
	"fmt"
	"ginBlog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(true)

	// Migrate the schema
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(10 * time.Second)

}
