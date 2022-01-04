package tools

import (
	"bytes"
	"fmt"
	"encoding/json"
)

func ToBuffer(a interface{})(bytes.Buffer,error){
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(a)
	if err != nil {
		err = fmt.Errorf("to buffer failed:%v",err)
	}
	return buf,err
}