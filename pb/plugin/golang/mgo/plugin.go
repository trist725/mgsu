package mgo

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/trist725/mgsu/pb/plugin/golang"
	"github.com/trist725/mgsu/pb/plugin/golang/generator"
)

var (
	collectionRegexp = regexp.MustCompile(`@collection(?:\s*=\s*(\d+))?`)
	msgRegexp        = regexp.MustCompile(`@msg(?:\s*=\s*(\S+))?`)
	mapKeyRegexp     = regexp.MustCompile(`@map_key(?:\s*=\s*(\S+))?`)
	sliceRegexp      = regexp.MustCompile(`@slice(?:\s*=\s*(\S+))?`)
)

type mgo struct {
	*golang.PluginSuper
}

func New() *mgo {
	return &mgo{
		PluginSuper: golang.NewPluginSuper(),
	}
}

func (p *mgo) Name() string {
	return "mgo-go"
}

func (p *mgo) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *mgo) GenerateImports(file *generator.FileDescriptor) {

}

func (p *mgo) Generate(fd *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	p.AddImport("msg")

	jsonPkg := p.NewImport("encoding/json")
	syncPkg := p.NewImport("sync")
	mongodbPkg := p.NewImport("github.com/trist725/mgsu/db/mongodb")
	mgoPkg := p.NewImport("github.com/globalsign/mgo")

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

		//log.Logger().Debug("message=%v", message)
		//log.Logger().Debug("message.Comment=[%s]", message.Comment)

		if matches := collectionRegexp.FindStringSubmatch(message.Comment); len(matches) > 0 {
			//log.Logger().Debug("matches=%v", matches)
			message.ID = message.Name
			//log.Logger().Debug("message.ID=%s", message.ID)
			if !mongodbPkg.IsUsed() {
				mongodbPkg.Use()
				p.AddImport(generator.GoImportPath(mongodbPkg.Location()))
			}
			if !mgoPkg.IsUsed() {
				mgoPkg.Use()
				p.AddImport(generator.GoImportPath(mgoPkg.Location()))
			}
		}

		if matches := mapKeyRegexp.FindStringSubmatch(message.Comment); len(matches) > 1 && matches[1] != "" {
			message.MapKey = matches[1]
			//log.Logger().Debug("message.MapKey=%s", message.MapKey)
		}

		if matches := sliceRegexp.FindStringSubmatch(message.Comment); len(matches) > 0 {
			message.Slice = true
			//log.Logger().Debug("message.Slice=%v", message.Slice)
		}

		if matches := msgRegexp.FindStringSubmatch(message.Comment); len(matches) > 0 {
			//log.Logger().Debug("matches=%v", matches)
			if len(matches) > 1 && matches[1] != "" {
				message.Msg = matches[1]
			} else {
				message.Msg = message.Name
			}
		}

		for commentIndex, fdp := range md.GetField() {
			field := newField(p.Generator, md, fdp, commentIndex)

			//log.Logger().Debug("field.Comment=%s", field.Comment)

			if matches := msgRegexp.FindStringSubmatch(field.Comment); len(matches) > 0 {
				if message.Msg == "" {
					message.Msg = message.Name
				}
				//log.Logger().Debug("matches=%v", matches)

				if len(matches) > 1 && matches[1] != "" {
					field.Msg = matches[1]
				} else {
					field.Msg = field.Name
				}
				if field.IsMessage() && field.IsRepeated() {
					if field.IsMap {
						field.MsgRepeatedMessageGoType = fmt.Sprintf("map[%s]*msg.%s", field.KeyType, field.ValueTypeToName)
					} else {
						field.MsgRepeatedMessageGoType = fmt.Sprintf("[]*msg.%s", field.GoTypeToName)
					}
				}
			}

			p.RecordTypeUse(fdp.GetTypeName())
			message.Fields = append(message.Fields, field)
		}

		p.RecordTypeUse(generator.CamelCaseSlice(md.TypeName()))
		file.Messages = append(file.Messages, message)
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
