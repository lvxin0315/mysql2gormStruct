package sql_to_struct

import "io/ioutil"

//写文件
func WriteFile(filePath string, content []byte) error {
	return ioutil.WriteFile(filePath, content, 0666)
}

