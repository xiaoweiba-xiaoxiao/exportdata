package queryoprate

import (
	"bytes"
	"time"
	"github.com/tidwall/gjson"
	"github.com/xiaoweiba-xiaoxiao/exportdata/tools"
)

func (de *defaultEsdata)ToMap()([]string,[]map[string]interface{},error){
	return de.datafunc(de.configByte,de.buf,de.args...)
}


func (ds *defaultSqlData)ToMap()([]string,[]map[string]interface{},error){
	return ds.datafunc(ds.configByte,ds.str,ds.args...)
}

func NewSqlDataWithFunc(str string,agrs []interface{},config []byte,c Sqlfunc)(Data){
	return &defaultSqlData{datafunc: c,configByte: config,args: agrs,str: str}
}

func NewSqlData(str string,agrs []interface{},config []byte)(Data){
	return &defaultSqlData{datafunc: QueryData,configByte: config,args: agrs,str: str}
}

func NewEsDataWithFunc(config []byte,buf bytes.Buffer,args []interface{},c Esfunc)(Data){
	return &defaultEsdata{configByte: config,args: args,buf: buf,datafunc: c}
}

func QueryData(configByte []byte,str string,args ...interface{})([]string,[]map[string]interface{},error){
	colums := []string{}
	host := gjson.GetBytes(configByte,"database.host").String()
	port := gjson.GetBytes(configByte,"database.port").String()
	user := gjson.GetBytes(configByte,"database.user").String()
	password := gjson.GetBytes(configByte,"database.password").String()
	dbname := gjson.GetBytes(configByte,"database.database").Array()[0].String()
	datasmap := []map[string]interface{}{}
	client := tools.NewDBClient(user,password,host,port,dbname)
	rows,err := client.DB.Query(str,args...)
	if err != nil {
		return colums,datasmap,err
	}
	client.DB.Close()
	colums,err = rows.Columns()
	if err != nil {
		return colums,datasmap,err
	}
	defer rows.Close()
	length := len(colums)
	caches := []interface{}{}
	for i := 0;i < length;i++{
		var a interface{}
		caches = append(caches, &a)
	}
	for rows.Next() {
		err := rows.Scan(caches...)
		if err != nil {
			return colums,datasmap,err
		}
		rowdata := map[string]interface{}{} 
		for i,cache := range caches {
			tempv := *cache.(*interface{})
			if v,ok := tempv.([]byte);ok {
			   rowdata[colums[i]] = string(v)
			   continue
			}
			if v,ok := tempv.(int64);ok {
				rowdata[colums[i]] = v
				continue
			}
			if v,ok := tempv.(time.Time);ok {
				rowdata[colums[i]] = v.String()
				continue
			}
		    rowdata[colums[i]]=tempv
		}
		datasmap = append(datasmap, rowdata)
	}
	return colums,datasmap,nil
}

func DataDo(data Data,c Do,args ...interface{})(interface{},error){
	return c(data,args...)
}

