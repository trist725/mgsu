package pbplugin

import (
	"fmt"
	"go/format"
	"strings"

	"bytes"

	"github.com/gogo/protobuf/proto"
	plugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

// filenameSuffix replaces the .pb.go at the end of each filename.
func Generate(req *plugin.CodeGeneratorRequest, p Plugin, filenameSuffix string) *plugin.CodeGeneratorResponse {
	g := New()
	g.Request = req
	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GeneratePlugin(p)

	for i := 0; i < len(g.Response.File); i++ {
		g.Response.File[i].Name = proto.String(
			strings.Replace(*g.Response.File[i].Name, ".pb.go", filenameSuffix, -1),
		)
	}
	if err := FormatSource(g.Response); err != nil {
		g.Error(err)
	}
	return g.Response
}

func FormatSource(resp *plugin.CodeGeneratorResponse) error {
	for i := 0; i < len(resp.File); i++ {
		formatted, err := format.Source([]byte(resp.File[i].GetContent()))
		if err != nil {
			return fmt.Errorf("go format error: %v", err)
		}
		content := string(formatted)
		resp.File[i].Content = &content
	}
	return nil
}

func MakeFirstLowerCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}

func MakeFirstUpperCase(s string) string {
	if len(s) < 2 {
		return strings.ToUpper(s)
	}

	bts := []byte(s)

	lc := bytes.ToUpper([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}

func ParseRequestParameterString(input string) (output map[string]string) {
	output = make(map[string]string)
	kvs := strings.Split(input, ";")
	for _, kv := range kvs {
		pair := strings.Split(kv, ":")
		if len(pair) == 1 {
			output[strings.TrimSpace(kv)] = ""
		} else if len(pair) == 2 {
			output[strings.TrimSpace(pair[0])] = strings.TrimSpace(pair[1])
		}
	}
	return
}
