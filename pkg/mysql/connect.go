package mysql

import (
	"context"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New(ctx context.Context, db_url string) *gorm.DB {
	var err error
	dialector := mysql.Open(db_url)

	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		connect(dialector, 100)
	}

	db.Logger = db.Logger.LogMode(logger.Silent)

	return db
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}}); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--

			connect(dialector, count)
			return
		}

		panic(err.Error())
	}
}
