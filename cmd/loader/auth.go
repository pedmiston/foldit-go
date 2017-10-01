package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Auth struct {
	DOAccessKey   string `yaml:"do_access_key"`
	DOSecretKey   string `yaml:"do_secret_key"`
	MySQLUsername string `yaml:"mysql_username"`
	MySQLPassword string `yaml:"mysql_password"`
}

var auth *Auth

func authenticate(creds string) {
	b, err := ioutil.ReadFile(creds)
	if err != nil {
		log.Fatal(err)
	}

	auth = &Auth{}
	err = yaml.Unmarshal(b, auth)
	if err != nil {
		log.Fatal(err)
	}

}
