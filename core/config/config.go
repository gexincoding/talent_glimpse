package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var TalentGlimpseConfig Config

type Config struct {
	MySQLConfig MySQLConfig `json:"MySQLConfig"`
	RedisConfig RedisConfig `json:"RedisConfig"`
}

type MySQLConfig struct {
	Name string `json:"Name"`
	Host string `json:"Host"`
	User string `json:"User"`
	Pass string `json:"Pass"`
}

type RedisConfig struct {
	Addr string `json:"Addr"`
	Pass string `json:"Pass"`
}

func Init() error {
	conf, err := ioutil.ReadFile("core/config/conf.json")
	if err != nil {
		log.Println(fmt.Errorf("read conf err: %v", err))
		return err
	}
	err = json.Unmarshal(conf, &TalentGlimpseConfig)
	if err != nil {
		log.Println(fmt.Errorf("conf bind json err: %v", err))
		return err
	}
	return nil
}
