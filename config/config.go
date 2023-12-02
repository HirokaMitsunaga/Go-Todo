package config

import (
	"Go-Todo/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)
type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

//外部から読み込みできるように宣言
var Config ConfigList


//init関数はmain関数実行の前に実行される
func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

//config.iniの情報を読み込む
func LoadConfig (){
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
	
}