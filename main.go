package main

import (
	_"bufio"
	_"bytes"
	_"context"
	_ "database/sql"
	_"encoding/json"
	"fmt"
	_"io/ioutil"
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/xiaoweiba-xiaoxiao/exportdata/queryoprate"
	"github.com/xiaoweiba-xiaoxiao/exportdata/tools"
	conf "github.com/xiaoweiba-xiaoxiao/goconfig/config"
	"github.com/xiaoweiba-xiaoxiao/gologs/logs"
)

var (
	logger logs.Logs
	exclename string
	configByte []byte
	sheetNames []string
	cellSlice  [][]string
	dataSlice  [][]map[string]interface{}
)


func initconfig(file string)(error){
	var err error
	configByte,err = conf.LoadYaml(file)
	return err
}

func initlogger(){
	logfile := gjson.GetBytes(configByte,"log.path").String()
	if logfile == ""{
		logfile = "./export.log"
	}
	logger = logs.NewLogger(logs.FILE,logfile)
}

func initxlsx(){
	exclename = gjson.GetBytes(configByte,"excel.name").String()
}



func main(){
	file := "./config.yaml"
	err := initconfig(file)
	if err != nil {
		panic(err)
	}
	initlogger()
	initxlsx()
	sheets := gjson.GetBytes(configByte,"excel.sheetsname").Array()
	for _,sheet :=  range sheets{
		args := []interface{}{}
		sheetname := sheet.Get("name").String()
		sheetNames = append(sheetNames, sheetname)
		sqlstr := sheet.Get("sqlstr").String()
		argsjosn := sheet.Get("args").Array()
		for _,arg :=range argsjosn{
			if argint,err := strconv.Atoi(arg.Raw);err == nil{
				args=append(args, argint)
				continue
			}
			if argfloat,err := strconv.ParseFloat(arg.Raw,32); err == nil{
				args = append(args, argfloat)
				continue
			}
			args = append(args,arg.Raw)
		}
		data := queryoprate.NewSqlData(sqlstr,args,configByte)
		colums,maps,err := data.ToMap()
		if err != nil {
		   err = fmt.Errorf("%v sheet unmashal []map[string]interface faild:%v",sheetname,err)
		   logger.Error(err)
		   continue
	    }
		if len(maps)==0{
			err = fmt.Errorf("%v,has no data",sheetname)
			logger.Info(err)
			continue
		}
		cellSlice = append(cellSlice, colums)
		dataSlice = append(dataSlice, maps)
	}
	tools.OutPut(exclename,sheetNames,cellSlice,dataSlice)
}