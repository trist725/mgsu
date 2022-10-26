package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/trist725/mgsu/util"
)

type config struct {
	ExcelDir      string
	SourceCodeDir string
	Debug         bool
	BlackList     []string
	WhiteList     []string
	blackList     map[string]struct{}
	whiteList     map[string]struct{}
}

func newConfig() *config {
	cfg := &config{
		ExcelDir:      ".",
		SourceCodeDir: "./sd",
		Debug:         true,
		blackList:     map[string]struct{}{},
		whiteList:     map[string]struct{}{},
	}
	return cfg
}

func (c *config) load(cfgFilePath string) {
	if err := util.IsDirOrFileExist(cfgFilePath); err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, c); err != nil {
		panic(err)
	}

	if len(c.BlackList) > 0 {
		for _, b := range c.BlackList {
			if b == "" {
				continue
			}
			c.blackList[b] = struct{}{}
		}
	}

	if len(c.WhiteList) > 0 {
		for _, w := range c.WhiteList {
			if w == "" {
				continue
			}
			c.whiteList[w] = struct{}{}
		}
	}
}

func (c config) check() {
	if c.ExcelDir == "" {
		panicf("excel目录为空")
	}
	if c.SourceCodeDir == "" {
		panicf("源代码目录为空")
	}
}
