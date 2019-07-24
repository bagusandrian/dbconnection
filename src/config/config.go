package config

import (
	"fmt"
	"log"

	gcfg "gopkg.in/gcfg.v1"
)

type (
	Config struct {
		Database struct {
			DBMaster string
			DBSlave  string
		}
	}
)

// FilePathList list of possible config file relative path to binary location
func FilePathList() []string {
	return []string{
		"files/etc/dbconnection/",
	}
}

// ReadConfig read `*.ini` configuration file and save it to variable of `*Config` type
func ReadConfig() *Config {
	var (
		cfg     Config
		err     error
		environ string
	)

	path := FilePathList()

	for _, val := range path {
		file := fmt.Sprintf("%sdbconnection.ini", val)
		log.Printf("%s\n", file)
		err := gcfg.ReadFileInto(&cfg, file)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("[dbconnection] Cannot load config env:%s :%+v\n ", environ, err)
	}

	log.Printf("[dbconnection] Config load success, using \"%s\".\n", environ)

	return &cfg
}
