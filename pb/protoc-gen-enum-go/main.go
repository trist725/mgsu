package main

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity/command"

	"github.com/trist725/mgsu/pb/pbplugin"
)

const t = `
{{range .Enums}}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [{{.CamelCaseName}}] begin

{{if .Comment}}{{.Comment}}{{end}}
type {{.CamelCaseName}} int32

const (
{{range .Values}}
{{if .Comment}}{{.Comment}}{{end}}
	{{.Name}} {{.TypeName}} = {{.Value}}{{end}}
)

var {{.CamelCaseName}}_name = map[int32]string{
	{{range .Values}}{{.Value}}:"{{.Name}}",
	{{end}}
	}

var {{.CamelCaseName}}_value = map[string]int32{
	{{range .Values}}"{{.Name}}":{{.Value}},
	{{end}}
	}

var {{.CamelCaseName}}_Slice = []int32{
	{{range .Values}}{{.Value}},
	{{end}}
	}

func (x {{.CamelCaseName}}) String() string {
	if name, ok := {{.CamelCaseName}}_name[int32(x)]; ok {
		return name
	}
	return ""
}

func {{.CamelCaseName}}_Len() int {
	return len({{.CamelCaseName}}_Slice)
}

{{if .Enable0}}
func Check_{{.CamelCaseName}}_I(value int32) bool {
	if _, ok := {{.CamelCaseName}}_name[value]; ok {
		return true
	}
	return false
}
{{else}}
func Check_{{.CamelCaseName}}_I(value int32) bool {
	if _, ok := {{.CamelCaseName}}_name[value]; ok && value != 0 {
		return true
	}
	return false
}
{{end}}

func Check_{{.CamelCaseName}}(value {{.CamelCaseName}}) bool {
	return Check_{{.CamelCaseName}}_I(int32(value))
}

func Each_{{.CamelCaseName}}(f func({{.CamelCaseName}}) bool) {
	for _, value := range {{.CamelCaseName}}_Slice {
		if !f({{.CamelCaseName}}(value)) {
			break
		}
	}
}

func Each_{{.CamelCaseName}}_I(f func(int32) bool) {
	for _, value := range {{.CamelCaseName}}_Slice {
		if !f(value) {
			break
		}
	}
}
// enum [{{.CamelCaseName}}] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
{{end}}
`

var enableZeroRegexp = regexp.MustCompile(`\s*@enable_0(?:\s*=\s*(\w+))?\s*`)

type enumValue struct {
	ed *pbplugin.EnumDescriptor
	dp *descriptor.EnumValueDescriptorProto

	Name     string
	TypeName string
	Value    int32
	Comment  string
}

type enum struct {
	g  *pbplugin.Generator
	ed *pbplugin.EnumDescriptor

	Name          string
	CamelCaseName string
	Comment       string
	Values        []*enumValue
	Enable0       string
}

func (e *enum) Prefix() string {
	return e.ed.Prefix()
}

type protoFile struct {
	PackageName string
	Enums       []*enum
}

func parse(g *pbplugin.Generator, fd *pbplugin.FileDescriptor) *protoFile {
	pf := &protoFile{}

	pf.PackageName = fd.PackageName()

	for _, ed := range fd.Enums() {
		name := ed.GetName()
		if ed.Parent() != nil {
			name = ed.Prefix() + name
		}

		e := &enum{
			g:             g,
			ed:            ed,
			Name:          name,
			CamelCaseName: pbplugin.CamelCase(name),
			Comment:       g.GetComments(ed.Path()),
		}

		if matches := enableZeroRegexp.FindStringSubmatch(e.Comment); len(matches) > 0 {
			//if matches[1] == "true" {
			e.Enable0 = "true"
			//}
		}

		for j, edp := range ed.GetValue() {
			v := &enumValue{
				ed:       ed,
				dp:       edp,
				Name:     *edp.Name,
				TypeName: e.CamelCaseName,
				Value:    *edp.Number,
				Comment:  g.GetComments(fmt.Sprintf("%s,2,%d", ed.Path(), j)),
			}

			e.Values = append(e.Values, v)
		}

		pf.Enums = append(pf.Enums, e)
	}

	return pf
}

type plugin struct {
	*pbplugin.Generator
}

func newPlugin() *plugin {
	p := &plugin{}
	return p
}

func (plugin) Name() string {
	return "enum-go"
}

func (p *plugin) Init(g *pbplugin.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(fd *pbplugin.FileDescriptor) {
	var code bytes.Buffer
	tpl.Execute(&code, parse(p.Generator, fd))
	s := code.String()
	p.P(s)
}

func (p *plugin) GenerateImports(fd *pbplugin.FileDescriptor) {
	// p.PrintImport("", "sync")
	// p.PrintImport("", "github.com/name5566/leaf/db/mongodb")
	// p.PrintImport("", "gopkg.in/mgo.v2")

	// p.P("var _ *sync.Pool")
	// p.P("var _ *mongodb.DialContext")
	// p.P("var _ *mgo.DBRef")
}

var (
	tpl *template.Template
)

func init() {
	var err error
	tpl, err = template.New("enum-go").Parse(t)
	if err != nil {
		panic(err)
	}
}

func main() {
	req := command.Read()
	p := newPlugin()
	res := pbplugin.Generate(req, p, ".enum.go")
	command.Write(res)
}
