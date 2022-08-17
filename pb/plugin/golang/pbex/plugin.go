package pbex

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/generator"
)

var (
	msgRegexp      = regexp.MustCompile(`@msg(?:\s*=\s*(\d+))?`)
	responseRegexp = regexp.MustCompile(`\s*@response(?:\s*=\s*(\S+))?\s*`)
)

func init() {
	generator.RegisterPlugin(New())
}

type pbex struct {
	*golang.PluginSuper
	version string
}

func New() *pbex {
	return &pbex{
		PluginSuper: golang.NewPluginSuper(),
		version:     "v1",
	}
}

func (p *pbex) Name() string {
	return PluginName
}

func (p *pbex) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *pbex) GenerateImports(file *generator.FileDescriptor) {

}

func (p *pbex) Generate(fd *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	jsonPkg := p.NewImport("encoding/json")
	syncPkg := p.NewImport("sync")
	reflectPkg := p.NewImport("reflect")

	protocolImportPath := fmt.Sprintf("github.com/trist725/mgsu/network/protocol/protobuf/%s", p.version)
	protocolPkg := p.NewImport(protocolImportPath)

	file := newFile(fd)
	file.Protocol = p.version

	for _, md := range file.FileDescriptor.Messages() {
		if md.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}

		if !jsonPkg.IsUsed() {
			jsonPkg.Use()
			p.AddImport(generator.GoImportPath(jsonPkg.Location()))
		}

		if !syncPkg.IsUsed() {
			syncPkg.Use()
			p.AddImport(generator.GoImportPath(syncPkg.Location()))
		}

		message := newMessage(p.Generator, md)

		if matches := msgRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 && matches[1] != "" {
			message.ID = matches[1]
			if !protocolPkg.IsUsed() {
				protocolPkg.Use()
				p.AddImport(generator.GoImportPath(protocolPkg.Location()))
			}
		}

		nameWords := strings.Split(message.Name, "_")
		if len(nameWords) > 1 && nameWords[0] == "C2S" {
			if matches := responseRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 {
				message.Response = matches[1]
			}
		}

		for commentIndex, fdp := range md.GetField() {
			field := golang.NewField(p.Generator, md, fdp, commentIndex)
			p.RecordTypeUse(fdp.GetTypeName())
			if !field.IsOneof {
				message.Fields = append(message.Fields, field)
			}
		}
		for _, oneof := range md.GetOneofDecl() {
			message.Oneofs = append(message.Oneofs, generator.CamelCase(oneof.GetName()))
			if !reflectPkg.IsUsed() {
				reflectPkg.Use()
				p.AddImport(generator.GoImportPath(reflectPkg.Location()))
			}
		}

		p.RecordTypeUse(generator.CamelCaseSlice(md.TypeName()))
		file.Messages = append(file.Messages, message)
		file.messagesByName[message.Name] = message
	}

	for _, message := range file.Messages {
		if message.Response != "" {
			continue
		}
		nameWords := strings.Split(message.Name, "_")
		if len(nameWords) <= 1 || nameWords[0] != "C2S" {
			continue
		}
		response := fmt.Sprintf("S2C_%s", nameWords[1])
		if _, ok := file.messagesByName[response]; ok {
			message.Response = response
		}
	}

	for commentIndex1, ed := range fd.Enums() {
		enum := golang.NewEnum(p.Generator, ed, commentIndex1)

		for commentIndex2, edp := range ed.GetValue() {
			enumValue := golang.NewEnumValue(p.Generator, enum.Name, edp, commentIndex1, commentIndex2)
			enum.Values = append(enum.Values, enumValue)
		}

		p.RecordTypeUse(enum.Name)
		file.Enums = append(file.Enums, enum)
	}

	var code bytes.Buffer
	err := p.Template.Execute(&code, file)
	if err != nil {
		panic(err)
	}
	s := code.String()
	p.P(s)
}

func (p *pbex) SetVersion(version string) {
	p.version = version
}
