package boot

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	g "main/global"
)

func DataBaseSetup() {
	MysqlInit()
	RedisInit()
}

func MysqlInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		g.Config.Database.Mysql.Username,
		g.Config.Database.Mysql.Password,
		g.Config.Database.Mysql.Addr,
		g.Config.Database.Mysql.Port,
		g.Config.Database.Mysql.DBName,
	)
	sqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
		g.Logger.Fatal("connect mysql failed")
		panic(err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		g.Logger.Fatal("connect mysql failed")
		panic(err)
	}
	g.Mdb = db
	g.Logger.Info("connect mysql successfully")
}

func RedisInit() {

}
