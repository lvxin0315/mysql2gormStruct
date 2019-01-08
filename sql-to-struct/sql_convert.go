package sql_to_struct

import (
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type SqlConvert struct {
	//转换的sql
	SqlString string
	//查询的db信息
	DbUser string
	DbPassword string
	DbHost string
	DbPort int
	DbDatabase string
	//转换使用的db信息
	ConvertDbUser string
	ConvertDbPassword string
	ConvertDbHost string
	ConvertDbPort int
	ConvertDbDatabase string
	convertTable string //转换的临时表名
}

func (sc *SqlConvert) createTable () error {
	UUID := uuid.NewV4()
	sc.convertTable = "tmp" + strings.Replace(UUID.String(),"-","_",4)
	sql := fmt.Sprintf("CREATE TABLE  %s.%s AS (%s)",sc.ConvertDbDatabase,sc.convertTable,sc.SqlString)
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		sc.DbUser,
		sc.DbPassword,
		sc.DbHost,
		sc.DbPort,
		sc.DbDatabase))
	defer db.Close()
	if err != nil{
		return err
	}
	db.Raw(sql).Row()
	return db.Error
}

func (sc *SqlConvert)Convert(structName, packageName, savePath string) (error) {
	//创建临时table
	err := sc.createTable()
	if err != nil{
		return err
	}
	//使用table_convert转换临时table
	tableConvert := TableConvert{
		sc.ConvertDbUser,
		sc.ConvertDbPassword,
		sc.ConvertDbHost,
		sc.ConvertDbPort,
		sc.ConvertDbDatabase,
	}
	return tableConvert.Convert(sc.convertTable,structName,packageName,savePath)
}