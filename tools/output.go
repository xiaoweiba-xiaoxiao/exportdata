package tools

import (
	"github.com/tealeg/xlsx"
)

func OutPut(filename string,sheets []string,cellslice [][]string, dataslice [][]map[string]interface{}) {
	lengthSheet := len(sheets)
	lengthCell := len(cellslice)
	lengthData := len(dataslice)
	if !(lengthSheet == lengthCell && lengthData == lengthCell) {
		return
	}
	of := xlsx.NewFile()
	for i,sheetname := range sheets{
		count := 0
		sheet, _ := of.AddSheet(sheetname)
		row, _ := sheet.AddRowAtIndex(count)
		for _, header := range cellslice[i] {
			cell := row.AddCell()
			cell.SetValue(header)
		}
		count++
		for _, data := range dataslice[i] {
			row, _ := sheet.AddRowAtIndex(count)
			for _, header := range cellslice[i] {
				cell := row.AddCell()
				cell.SetValue(data[header])
			}
			count++
		}
	}	
	of.Save(filename)
}