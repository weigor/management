package common

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbPath string
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

func init() {
	flag.StringVar(&dbPath, "conf", "./config/db.toml", "")
}

type Orm struct {
	*gorm.DB
}

func MysqlInit() *Orm {
	var conf Config
	_, err := toml.DecodeFile(dbPath, &conf)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Dbname, conf.Mysql.Timeout)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	defer db.Close()
	return &Orm{
		DB: db,
	}
}
