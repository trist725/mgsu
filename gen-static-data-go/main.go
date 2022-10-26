package main

import (
	"flag"
)

var (
	flagConfigFilePath = flag.String("c", "./sd.json", "config file path, default=./sd.json")

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
