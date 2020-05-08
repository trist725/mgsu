package main

import (
	"io/ioutil"

	"gitee.com/nggs/tools/pbplugin"
	"gitee.com/nggs/tools/pbplugin/golang"
	"gitee.com/nggs/tools/pbplugin/golang/command"
	"gitee.com/nggs/tools/pbplugin/golang/pbex"
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
