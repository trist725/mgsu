package main

import (
	"io/ioutil"

	pbplugin "github.com/trist725/mgsu/pb/plugin"

	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/command"
	"github.com/trist725/mgsu/pb/plugin/golang/pbex"
)

func main() {
	req := command.Read()

	reqParams := map[string]string{}
	if req.GetParameter() != "" {
		reqParams = pbplugin.ParseRequestParameterString(req.GetParameter())
	}

	p := pbex.New()

	if templateFilePath, ok := reqParams["tpl"]; ok {
		templateContent, err := ioutil.ReadFile(templateFilePath)
		if err != nil {
			panic(err)
		}
		p.MustLoadTemplate(pbex.PluginName, string(templateContent), golang.TemplateFunc)
	} else {
		p.MustLoadTemplate(pbex.PluginName, pbex.DefaultTemplate, golang.TemplateFunc)
	}

	if version, ok := reqParams["version"]; ok && version != "" {
		p.SetVersion(version)
	}

	resp := command.GeneratePlugin(req, p, ".pbex.go")
	command.Write(resp)
}
