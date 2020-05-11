package main

import (
	"io/ioutil"

	"github.com/trist725/mgsu/pb/plugin"
	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/command"
	"github.com/trist725/mgsu/pb/plugin/golang/mgo"
)

func main() {
	req := command.Read()

	reqParams := map[string]string{}
	if req.GetParameter() != "" {
		reqParams = pbplugin.ParseRequestParameterString(req.GetParameter())
	}

	p := mgo.New()

	if templateFilePath, ok := reqParams["tpl"]; ok {
		templateContent, err := ioutil.ReadFile(templateFilePath)
		if err != nil {
			panic(err)
		}
		p.MustLoadTemplate("mgo-go", string(templateContent), golang.TemplateFunc)
	} else {
		p.MustLoadTemplate("mgo-go", mgo.DefaultTemplate, golang.TemplateFunc)
	}

	resp := command.GeneratePlugin(req, p, ".mgo.go")
	command.Write(resp)
}
