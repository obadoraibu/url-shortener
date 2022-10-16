package main

import (
	"flag"
	"io/ioutil"
	"log"
	"obadoraibu/url-shortener/internal/app/apiserver"

	"gopkg.in/yaml.v3"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("Error reading yaml file:  #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
