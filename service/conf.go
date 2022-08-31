package service

import (
	_ "embed"
	"encoding/json"
	"log"
)

var Conf struct {
	GRPCPort   string
	BasePrefix string
}

//go:embed conf.json
var data []byte

func init() {
	err := json.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
