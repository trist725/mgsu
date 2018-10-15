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
{{range .Messages}}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [{{.CamelCaseName}}] begin

{{if .Comment}}{{.Comment}}{{end}}
type {{.CamelCaseName}} struct {
{{range .Fields}}{{if .Comment}}{{.Comment}}{{end}}
	{{.CamelCaseName}} {{.TypeName}}
{{end}}
}

func New_{{.CamelCaseName}}() *{{.CamelCaseName}} {
	m := &{{.CamelCaseName}}{
	{{.NewBody}}
	}
	return m
}

func (m {{.CamelCaseName}}) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"{{.CamelCaseName}}\":%s}", string(ba))
}

func (m *{{.CamelCaseName}}) Reset() {
	{{.ResetBody}}
}

func (m {{.CamelCaseName}}) Clone() *{{.CamelCaseName}} {
	{{.CloneBody}}
}

func Clone_{{.CamelCaseName}}_Slice(dst []*{{.CamelCaseName}}, src []*{{.CamelCaseName}}) []*{{.CamelCaseName}} {
	for _, i := range dst {
		Put_{{.CamelCaseName}}(i)
	}
	dst = []*{{.CamelCaseName}}{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (m *{{.CamelCaseName}}) Unmarshal(dAtA []byte) error {
	return json.Unmarshal(dAtA, m)
}

func (m {{.CamelCaseName}}) Marshal() (dAtA []byte, err error) {
	return json.Marshal(m)
}

{{if .ID}}
func ({{.CamelCaseName}}) Json() {
}

func ({{.CamelCaseName}}) MessageID() protocol.MessageID {
	return "{{.ID}}"
}

func {{.CamelCaseName}}_MessageID() protocol.MessageID {
	return "{{.ID}}"
}
{{end}}

var g_{{.CamelCaseName}}_Pool = sync.Pool{}

func Get_{{.CamelCaseName}}() *{{.CamelCaseName}} {
	m, ok := g_{{.CamelCaseName}}_Pool.Get().(*{{.CamelCaseName}})
	if !ok {
		m = New_{{.CamelCaseName}}()
	} else {
		if m == nil {
			m = New_{{.CamelCaseName}}()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_{{.CamelCaseName}}(i interface{}) {
	if m, ok := i.(*{{.CamelCaseName}}); ok && m != nil {
		g_{{.CamelCaseName}}_Pool.Put(i)
	}
}

{{if .ID}}
func init() {
	Protocol.Register(
		&{{.CamelCaseName}}{},
		func() protocol.IMessage { return Get_{{.CamelCaseName}}() },
		func(msg protocol.IMessage) { Put_{{.CamelCaseName}}(msg) },
	)
}
{{end}}
// message [{{.CamelCaseName}}] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
{{end}}
`

var msgRegexp = regexp.MustCompile(`\s*@msg(?:\s*=\s*(\d+))?\s*`)

type messageField struct {
	md *pbplugin.Descriptor
	dp *descriptor.FieldDescriptorProto

	Name          string
	TypeName      string
	JsonName      string
	CamelCaseName string
	Comment       string
}

type message struct {
	g *pbplugin.Generator

	ID            string
	Name          string
	CamelCaseName string
	Type          string
	Comment       string
	Fields        []*messageField
}

func (m *message) NewBody() string {
	src := ""
	for _, f := range m.Fields {
		typeName, _ := m.g.GoType(f.md, f.dp)
		switch *f.dp.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("%s:New_%s(),\n", f.CamelCaseName, pbplugin.GoTypeToName(typeName))
			} else {
				src += fmt.Sprintf("%s:%s{},\n", f.CamelCaseName, typeName)
			}

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("%s:[]byte{},\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf("%s:[][]byte{},\n", f.CamelCaseName)
			}

		default:
			if pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("%s:%s{},\n", f.CamelCaseName, f.TypeName)
			}
		}
	}
	return src
}

func (m *message) ResetBody() string {
	src := ""
	for _, f := range m.Fields {
		typeName, _ := m.g.GoType(f.md, f.dp)
		switch *f.dp.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("m.%s.Reset()\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf(`
				for _, i := range m.%s {
					Put_%s(i)
				}
				m.%s = %s{}
			   	`, f.CamelCaseName, pbplugin.GoTypeToName(typeName), f.CamelCaseName, typeName)
			}

		case descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_UINT32,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_UINT64,
			descriptor.FieldDescriptorProto_TYPE_SINT32,
			descriptor.FieldDescriptorProto_TYPE_SINT64,
			descriptor.FieldDescriptorProto_TYPE_ENUM:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("m.%s = 0\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf("m.%s = %s{}\n", f.CamelCaseName, f.TypeName)
			}

		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("m.%s = false\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf("m.%s = []%s{}\n", f.CamelCaseName, pbplugin.GoTypeToName(typeName))
			}

		case descriptor.FieldDescriptorProto_TYPE_STRING:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("m.%s = \"\"\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf("m.%s = []%s{}\n", f.CamelCaseName, pbplugin.GoTypeToName(typeName))
			}

		case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
			descriptor.FieldDescriptorProto_TYPE_FLOAT,
			descriptor.FieldDescriptorProto_TYPE_FIXED32,
			descriptor.FieldDescriptorProto_TYPE_FIXED64,
			descriptor.FieldDescriptorProto_TYPE_SFIXED32,
			descriptor.FieldDescriptorProto_TYPE_SFIXED64:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("m.%s = 0.0\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf("m.%s = []%s{}\n", f.CamelCaseName, pbplugin.GoTypeToName(typeName))
			}

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("m.%s = []byte{}\n", f.CamelCaseName)
			} else {
				src += fmt.Sprintf("m.%s = [][]%s{}\n", f.CamelCaseName, pbplugin.GoTypeToName(typeName))
			}
		}
	}
	return src
}

func (m *message) CloneBody() string {
	src := fmt.Sprintf(`n, ok := g_%s_Pool.Get().(*%s)
	if !ok || n == nil {
		n = &%s{}
	}

`, m.CamelCaseName, m.CamelCaseName, m.CamelCaseName)

	for _, f := range m.Fields {
		typeName, _ := m.g.GoType(f.md, f.dp)
		//goTypeToName := pbplugin.GoTypeToName(typeName)
		switch *f.dp.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("n.%s = m.%s.Clone()\n", f.CamelCaseName, f.CamelCaseName)
			} else {
				src += fmt.Sprintf(
					`
if len(m.%s) > 0 {
	for _, i := range m.%s {
		if i != nil {
			n.%s = append(n.%s, i.Clone())
		} else {
			n.%s = append(n.%s, nil)
		}
	}
} else {
	n.%s = %s{}
}

`,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					typeName)
			}

		case descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_UINT32,
			descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_UINT64,
			descriptor.FieldDescriptorProto_TYPE_SINT32,
			descriptor.FieldDescriptorProto_TYPE_SINT64,
			descriptor.FieldDescriptorProto_TYPE_ENUM,
			descriptor.FieldDescriptorProto_TYPE_BOOL,
			descriptor.FieldDescriptorProto_TYPE_STRING,
			descriptor.FieldDescriptorProto_TYPE_DOUBLE,
			descriptor.FieldDescriptorProto_TYPE_FLOAT,
			descriptor.FieldDescriptorProto_TYPE_FIXED32,
			descriptor.FieldDescriptorProto_TYPE_FIXED64,
			descriptor.FieldDescriptorProto_TYPE_SFIXED32,
			descriptor.FieldDescriptorProto_TYPE_SFIXED64:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf("n.%s = m.%s\n", f.CamelCaseName, f.CamelCaseName)
			} else {
				src += fmt.Sprintf(
					`
if len(m.%s) > 0 {
								n.%s = make(%s, len(m.%s))
								copy(n.%s, m.%s)
							} else {
								n.%s = %s{}
							}
`,
					f.CamelCaseName,
					f.CamelCaseName,
					f.TypeName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.TypeName)
			}

		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			if !pbplugin.IsRepeated(f.dp) {
				src += fmt.Sprintf(
					`
if len(m.%s) > 0 {
								n.%s = make([]byte, len(m.%s))
								copy(n.%s, m.%s)
							} else {
								n.%s = []byte{}
							}

`,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName)
			} else {
				src += fmt.Sprintf(
					`
	if len(m.%s) > 0 {
		for _, b := range m.%s {
			if len(b) > 0 {
				nb := make([]byte, len(b))
				copy(nb, b)
				n.%s = append(n.%s, nb)
			} else {
				n.%s = append(n.%s, []byte{})
			}
		}
	} else {
		n.%s = [][]byte{}
	}

`,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName,
					f.CamelCaseName)
			}
		}
	}

	src += "\nreturn n"

	return src
}

type enumValue struct {
	ed *pbplugin.EnumDescriptor
	dp *descriptor.EnumValueDescriptorProto

	Name     string
	TypeName string
	Value    int32
	Comment  string
}

func (ev *enumValue) Prefix() string {
	if ev.ed.Parent() != nil {
		return ev.ed.Prefix()
	}
	return ""
}

type enum struct {
	g  *pbplugin.Generator
	ed *pbplugin.EnumDescriptor

	Name          string
	CamelCaseName string
	Comment       string
	Values        []*enumValue
}

func (e *enum) Prefix() string {
	return e.ed.Prefix()
}

type protoFile struct {
	PackageName string
	Messages    []*message
	Enums       []*enum
}

func parse(g *pbplugin.Generator, fd *pbplugin.FileDescriptor) *protoFile {
	pf := &protoFile{}

	pf.PackageName = fd.PackageName()

	for _, md := range fd.Messages() {
		m := &message{
			g:       g,
			Name:    pbplugin.CamelCaseSlice(md.TypeName()),
			Comment: g.GetComments(md.Path()),
		}
		m.CamelCaseName = pbplugin.CamelCase(m.Name)

		if matches := msgRegexp.FindStringSubmatch(m.Comment); len(matches) > 0 {
			m.ID = m.CamelCaseName
			//m.Comment = strings.Trim(m.Comment, matches[1])
		}

		for j, fdp := range md.GetField() {
			f := &messageField{
				md:       md,
				dp:       fdp,
				Name:     g.GetFieldName(md, fdp),
				JsonName: fdp.GetJsonName(),
				Comment:  g.GetComments(fmt.Sprintf("%s,%d,%d", md.Path(), 2, j)),
			}
			f.CamelCaseName = pbplugin.CamelCase(f.Name)
			f.TypeName, _ = g.GoType(md, fdp)

			m.Fields = append(m.Fields, f)
		}

		pf.Messages = append(pf.Messages, m)
	}

	for _, ed := range fd.Enums() {
		name := ed.GetName()
		if name == "MSG" {
			// 跳过叫MSG的枚举
			continue
		}

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

		for j, edp := range ed.GetValue() {
			v := &enumValue{
				ed:       ed,
				dp:       edp,
				Name:     *edp.Name,
				TypeName: e.CamelCaseName,
				Value:    *edp.Number,
				Comment:  g.GetComments(fmt.Sprintf("%s,2,%d", ed.Path(), j)),
			}

			if ed.Parent() != nil {
				v.Name = fmt.Sprintf("%s_%s", e.CamelCaseName, *edp.Name)
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
	return "myjson-go"
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
	p.PrintImport("", "fmt")
	p.PrintImport("", "sync")
	p.PrintImport("protocol", "github.com/trist725/mgsu/network/protocol/json")

	p.P("var _ = fmt.Sprintf")
	p.P("var _ *sync.Pool")
	p.P("var _ = protocol.PH")
}

var (
	tpl *template.Template
)

func init() {
	var err error
	tpl, err = template.New("myjson-go").Parse(t)
	if err != nil {
		panic(err)
	}
}

func main() {
	req := command.Read()
	p := newPlugin()
	res := pbplugin.Generate(req, p, ".myjson.go")
	command.Write(res)
}
