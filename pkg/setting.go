package pkg

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPORT     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/env.ini")
	if err != nil {
		log.Fatalf("load setting ini error %v", err)
	}

}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("bebug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fatal to get setction 'server' %v", err)
	}
	HTTPPORT = sec.Key("HTTP_PORT").MustInt(8765)
	ReadTimeOut = time.Duration(sec.Key("Read_TimeOut").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("Write_TimeOut").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("fatal to get setction 'app' %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_Secret").MustString("!)&@%&*$E^$")
}
