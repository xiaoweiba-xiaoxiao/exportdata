package queryoprate

import "bytes"

type Sqlfunc func(configbyte []byte,str string,agrs ...interface{})([]string,[]map[string]interface{},error)

type Esfunc func(configbyte []byte,buf bytes.Buffer,agrs ...interface{})([]string,[]map[string]interface{},error)

type Do func(data Data,args ...interface{})(interface{},error)

type Data interface{
	ToMap() ([]string,[]map[string]interface{},error)
}

type defaultSqlData struct{
	configByte []byte
    args []interface{}
	str string
	datafunc Sqlfunc
}

type defaultEsdata struct{
	configByte []byte
	buf bytes.Buffer
	args []interface{}
	datafunc Esfunc
}
