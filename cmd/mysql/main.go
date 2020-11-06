package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"management/model"
)

var (
	dbPath string
	do     string
)

type Mysql struct {
	Username string
	Password string
	Host     string
	Port     int
	Dbname   string
	Timeout  string
}
type Config struct {
	Mysql *Mysql
}

type Orm struct {
	*gorm.DB
	*Config
}

func init() {
	flag.StringVar(&dbPath, "db", "", "")
}

func MysqlInit(s string) *Orm {
	var conf Config
	var dsn string
	_, err := toml.DecodeFile(dbPath, &conf)
	if err != nil {
		panic(err)
	}
	if s == "create" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, "information_schema", conf.Mysql.Timeout)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Dbname, conf.Mysql.Timeout)
	}
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return &Orm{
		DB:     db,
		Config: &conf,
	}
}

func tables() []interface{} {
	return []interface{}{
		&model.User{},
	}
}

func main() {
	flag.StringVar(&do, "do", "", "do")
	flag.Parse()

	switch do {
	case "create":
		db := MysqlInit("create")
		createDbSQL := "CREATE DATABASE IF NOT EXISTS " + db.Mysql.Dbname + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
		err := db.Exec(createDbSQL).Error
		if err != nil {
			panic("创建失败：" + err.Error() + " sql:" + createDbSQL)
		}
		defer db.Close()
	case "drop":
		db := MysqlInit("drop")
		dropDbSQL := "DROP DATABASE IF EXISTS " + db.Mysql.Dbname + ";"
		err := db.Exec(dropDbSQL).Error
		if err != nil {
			panic("删除失败：" + err.Error() + " sql:" + dropDbSQL)
		}
		defer db.Close()

	case "migrate":
		db := MysqlInit("migrate")
		table := tables()
		for _, v := range table {
			if db.HasTable(v) { //判断表是否存在
				db.AutoMigrate(v) //存在就自动适配表，也就说原先没字段的就增加字段
			} else {
				db.CreateTable(v) //不存在就创建新表
			}
		}
		defer db.Close()
	default:
		panic(fmt.Sprintf("nothing to do:%s", do))
	}
}
