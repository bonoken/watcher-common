package example

import (
	"github.com/bonoken/watcher-common/common/config"
	"log"
)

var (
	Config *Properties
)

func mainTest() {

	// read config.yml
	if err := config.ReadConfig(&Config, "./config", "config"); err != nil {
		log.Fatal("config reading error")
	}

}

type Properties struct {
	Project struct {
		Url string
		Env string
	}
	Log struct {
		Level string
		Path  string
		File  string
	}
	Database struct {
		Dialect    string
		Connection string
	}
	HttpPort   string
	Encryption struct {
		Key string
		Iv  string
	}
	Interface struct {
		AwsCollector string
	}
}
