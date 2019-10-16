package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File

	UseGoogleUrlCheck bool
	GoogleApiKey      string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/configs.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadGoogleService()

}

func LoadGoogleService() {
	sec, err := Cfg.GetSection("google")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	UseGoogleUrlCheck = sec.Key("UseGoogleUrlCheck").MustBool(false)
	GoogleApiKey = sec.Key("GoogleApiKey").MustString("")

}
