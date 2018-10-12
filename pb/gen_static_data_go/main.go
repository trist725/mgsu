package main

import (
	"flag"
)

var (
	flagConfigFilePath = flag.String("c", "./gen_static_data_go.json", "config file path, default=./gen_static_data_go.json")

	gCfg *config
)

func init() {
	flag.Parse()

	gCfg = newConfig()

	gCfg.load(*flagConfigFilePath)

	gCfg.check()
}

func main() {
	sdcg, err := newStaticDataCodeGenerator(gCfg)
	if err != nil {
		panic(err)
	}

	if err := sdcg.generate(); err != nil {
		panic(err)
	}
}
