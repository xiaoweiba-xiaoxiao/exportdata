package tools

import (
	"testing"
)

func TestOutPut(t *testing.T){
	filename := "test1.xlsx"
	sheets := []string{"天晴","阴天","雨天"}
    cells  := [][]string{
		      {"城市","温度","最高温度"},
			  {"城市","温度","最高温度"},
			  {"城市","温度","最高温度"},
			}
	dataslice := [][]map[string]interface{}{
		{{"城市":"杭州","温度":"16-25度","最高温度":"25度",}},
		{{"城市":"湛江","温度":"21-29度","最高温度":"29度",}},
		{{"城市":"北海","温度":"25-31度","最高温度":"31度",}},
	}
	OutPut(filename,sheets,cells,dataslice)
}