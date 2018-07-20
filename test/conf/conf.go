package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	MongoURL = "localhost:27017"
	MongoDB  = "GMS"
)

type Gms struct {
	Price int
}

type Mgo struct {
	Database string
	URL      string
}

type Config struct {
	Mod string
	Mgo *Mgo
	Gms *Gms
}

var Conf Config

func init() {
	err := ParseConf("./conf/conf.json", &Conf)
	if err != nil {
		log.Fatal("[Parse configuraion]", err)
	}
}

// 解析 .json 类型配置文件
func ParseConf(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)

	return err
}
