package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MsDB struct {
	*sql.DB
}

var DB *sql.DB

func init() {
	dsn := fmt.Sprintf("root:@tcp(localhost:3306)/zerokk_go_blog?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}
	DB = db
}

func (d *MsDB) QueryOne(model interface{}, sql string, args ...interface{}) error {
	rows, err := d.Query(sql, args...)
	if err != nil {
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	vals := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for k := range vals {
		scans[k] = &vals[k]
	}
	if rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return err
		}
	}
	var result = make(map[string]interface{})
	elem := reflect.ValueOf(model).Elem()
	for index, val := range columns {
		result[val] = string(vals[index])
	}
	for i := 0; i < elem.NumField(); i++ {
		structField := elem.Type().Field(i)
		fieldInfo := structField.Tag.Get("orm")
		v := result[fieldInfo]
		t := structField.Type
		switch t.String() {
		case "int":
			s := v.(string)
			vInt, _ := strconv.Atoi(s)
			elem.Field(i).Set(reflect.ValueOf(vInt))
		case "string":
			elem.Field(i).Set(reflect.ValueOf(v.(string)))
		case "int64":
			s := v.(string)
			vInt64, _ := strconv.ParseInt(s, 10, 64)
			elem.Field(i).Set(reflect.ValueOf(vInt64))
		case "int32":
			s := v.(string)
			vInt32, _ := strconv.ParseInt(s, 10, 32)
			elem.Field(i).Set(reflect.ValueOf(vInt32))
		case "time.Time":
			s := v.(string)
			t, _ := time.Parse(time.RFC3339, s)
			elem.Field(i).Set(reflect.ValueOf(t))
		}
	}
	return nil
}
