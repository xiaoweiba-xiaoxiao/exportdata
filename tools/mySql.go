package tools

import (
	_ "database/sql"
	"fmt"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SqlClient struct{
	DB *sqlx.DB
}

var mu sync.Mutex
var clinet SqlClient


func NewDBClient(user,password,host,port,database string) SqlClient {	
	var db *sqlx.DB
	var err error	
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",user,password,host,port,database)
	mu.Lock()
	db,err = sqlx.Open("mysql",dns)
	if err != nil {
		err = fmt.Errorf("Mysql数据库初始化失败:%v")
		panic("Mysql数据库初始化失败")
	}
	clinet.DB = db
	defer mu.Unlock()
	return clinet
}


