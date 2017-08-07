package config

import "encoding/json"
import "os"

type Configuration struct {
	HostName string
	Port     string
}

func AppConfig() Configuration {

	file, ok := os.Open("config/httpd.json")
	defer file.Close()

	if ok != nil {
		panic("Can't open config file: httpd.json")
	}

	decoder := json.NewDecoder(file)
	cfg := Configuration{}

	if ok := decoder.Decode(&cfg); ok != nil {
		panic(ok)
	}
	return cfg
}
