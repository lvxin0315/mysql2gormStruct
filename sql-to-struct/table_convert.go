package sql_to_struct

import (
	"github.com/Shelnutt2/db2struct"
	"fmt"
)

type TableConvert struct {
	DbUser string
	DbPassword string
	DbHost string
	DbPort int
	DbDatabase string
}

//获取被转换结构体内容
func (tc *TableConvert)GetConvertContent(tableName, structName, packageName string) ([]byte, error) {
	columnDataTypes, err := db2struct.GetColumnsFromMysqlTable(
		tc.DbUser,
		tc.DbPassword,
		tc.DbHost,
		tc.DbPort,
		tc.DbDatabase,
		tableName)
	if err != nil {
		fmt.Println("Error in selecting column data information from mysql information schema")
		return []byte{}, err
	}
	struc, err := db2struct.Generate(
		*columnDataTypes,
		tableName,
		structName,
		packageName,
		true,
		true,
		false)
	if err != nil {
		fmt.Println("Error in creating struct from json: " + err.Error())
		return []byte{}, err
	}
	return struc, nil
}


//转换结构体并保存
func (tc *TableConvert)Convert(tableName, structName, packageName, savePath string) (error) {
	content,err := tc.GetConvertContent(tableName,structName,packageName)
	if err != nil{
		return err
	}
	err = WriteFile(savePath,content)
	if err != nil{
		return err
	}
	return nil
}