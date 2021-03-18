package config

import (
	"github.com/go-ini/ini"
	"log"
)

type app struct {
	JwtSecret string
	PageSize  int
}

var AppSetting = &app{}

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &database{}

type Logger struct {
	Level      int8
	Path       string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var LoggerSetting = &Logger{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	mapTo("logger", LoggerSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
