package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type  Conf struct{
	TaskName string `json:"TaskName"`
	TaskHost string `json:"TaskHost"`
	TaskPort string `json:"TaskPort"`
}

var _cfg *Conf=nil

func ParseConf(path string) (*Conf,error){
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	reader:=bufio.NewReader(file)
	decoder:=json.NewDecoder(reader)
	if err=decoder.Decode(&_cfg); err!=nil{
		return nil,err
	}
	return _cfg,nil
}
