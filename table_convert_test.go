package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	"mysql2gormStruct/sql-to-struct"
	"testing"
)

var dbUser = "jcc"
var dbPassword = "jcc"
var dbHost = "192.168.1.205"
var dbPort = 3306
var dbDatabase = "erp_205"

//查询db所有表
func getTables() []string {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sql := fmt.Sprintf("SELECT TABLE_NAME as tn FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA LIKE '%s' ", dbDatabase)
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var tables []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		//fmt.Println(rows.Columns())
		tables = append(tables, table)
	}
	return tables
}

func TestTableConvert(t *testing.T) {
	tc := sql_to_struct.TableConvert{
		DbDatabase:dbDatabase,
		DbHost:dbHost,
		DbPassword:dbPassword,
		DbPort:dbPort,
		DbUser:dbUser,
	}

	tables := getTables()
	for index := range tables{
		err := tc.Convert(tables[index],sql_to_struct.CamelString(tables[index]),"testStruct","testStruct/" + sql_to_struct.CamelString(tables[index]) + ".go")
		if err != nil{
			t.Error(err)
			break
		}
	}
}

func TestSqlConvert(t *testing.T) {
	sql := "SELECT g_name,g_man,g_address FROM viooma_guest"
	sc := sql_to_struct.SqlConvert{
		SqlString:sql,
		DbDatabase:dbDatabase,
		DbHost:dbHost,
		DbPassword:dbPassword,
		DbPort:dbPort,
		DbUser:dbUser,
		ConvertDbDatabase:"qing",
		ConvertDbHost:dbHost,
		ConvertDbPassword:dbPassword,
		ConvertDbPort:dbPort,
		ConvertDbUser:dbUser,
	}
	err := sc.Convert("TestSqlStruct","testStruct","testStruct/TestSqlStruct.go")
	if err != nil{
		t.Error(err)
	}
}
