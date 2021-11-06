package conf

import (
	"encoding/json"
	"io/ioutil"
)

var Conf *AppConfig

type AppConfig struct {
	MySQLConfig MySQLConfig
}

type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
}

func InitAppConfig() error {
	file, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		return err
	}

	var conf AppConfig
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return err
	}

	Conf = &conf
	return nil
}
