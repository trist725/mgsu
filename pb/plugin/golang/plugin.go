package golang

import (
	"text/template"

	"github.com/trist725/mgsu/pb/plugin/golang/generator"
)

type PluginSuper struct {
	*generator.Generator
	generator.PluginImports

	Template *template.Template
}

func NewPluginSuper() *PluginSuper {
	return &PluginSuper{}
}

func (p *PluginSuper) MustLoadTemplate(templateName string, templateContent string, templateFunc template.FuncMap) {
	var err error
	if templateFunc != nil {
		p.Template, err = template.New(templateName).Funcs(templateFunc).Parse(templateContent)
	} else {
		p.Template, err = template.New(templateName).Parse(templateContent)
	}
	if err != nil {
		panic(err)
	}
}
