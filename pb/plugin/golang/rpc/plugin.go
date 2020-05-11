package rpc

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/generator"
	//"github.com/trist725/mgsu/pb/plugin/log"
)

var (
	msgRegexp             = regexp.MustCompile(`\s*@msg(?:\s*=\s*(\d+))?\s*`)
	responseRegexp        = regexp.MustCompile(`\s*@response(?:\s*=\s*(\S+))?\s*`)
	responseTimeoutRegexp = regexp.MustCompile(`\s*@response_timeout(?:\s*=\s*(\d+))?\s*`)
	canIgnoreRegexp       = regexp.MustCompile(`\s*@can_ignore(?:\s*=\s*(\d+))?\s*`)
)

type rpc struct {
	*golang.PluginSuper
}

func New() *rpc {
	return &rpc{
		PluginSuper: golang.NewPluginSuper(),
	}
}

func (p *rpc) Name() string {
	return "rpc-go"
}

func (p *rpc) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *rpc) GenerateImports(file *generator.FileDescriptor) {

}

func (p *rpc) Generate(fd *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	p.AddImport("model")
	p.AddImport("msg")

	jsonPkg := p.NewImport("encoding/json")
	syncPkg := p.NewImport("sync")
	rpcPkg := p.NewImport("gitee.com/nggs/microservice/rpc")
	actorPkg := p.NewImport("gitee.com/nggs/protoactor-go/actor")
	myactorPkg := p.NewImport("gitee.com/nggs/actor")
	timePkg := p.NewImport("time")

	file := newFile(fd)

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

		if matches := msgRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 {
			message.ID = matches[1]
			if !rpcPkg.IsUsed() {
				rpcPkg.Use()
				p.AddImport(generator.GoImportPath(rpcPkg.Location()))
			}
		}

		nameWords := strings.Split(message.Name, "_")
		if len(nameWords) > 1 && nameWords[0] == "C2S" {
			if matches := responseRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 {
				message.Response = matches[1]
				if !actorPkg.IsUsed() {
					actorPkg.Use()
					p.AddImport(generator.GoImportPath(actorPkg.Location()))
				}
				if !myactorPkg.IsUsed() {
					myactorPkg.Use()
					p.AddImport(generator.GoImportPath(myactorPkg.Location()))
				}
				if !timePkg.IsUsed() {
					timePkg.Use()
					p.AddImport(generator.GoImportPath(timePkg.Location()))
				}
			}
			if matches := responseTimeoutRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 {
				message.ResponseTimeout = matches[1]
			}
		}

		if matches := canIgnoreRegexp.FindStringSubmatch(message.Comment); len(matches) > 0 {
			message.CanIgnore = true
		}

		for commentIndex, fdp := range md.GetField() {
			field := golang.NewField(p.Generator, md, fdp, commentIndex)
			p.RecordTypeUse(fdp.GetTypeName())
			message.Fields = append(message.Fields, field)
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
			if !actorPkg.IsUsed() {
				actorPkg.Use()
				p.AddImport(generator.GoImportPath(actorPkg.Location()))
			}
			if !myactorPkg.IsUsed() {
				myactorPkg.Use()
				p.AddImport(generator.GoImportPath(myactorPkg.Location()))
			}
			if !timePkg.IsUsed() {
				timePkg.Use()
				p.AddImport(generator.GoImportPath(timePkg.Location()))
			}
			if matches := responseTimeoutRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 {
				message.ResponseTimeout = matches[1]
			}
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

func init() {
	generator.RegisterPlugin(New())
}
