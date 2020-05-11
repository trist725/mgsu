package rpc

const DefaultTemplate = `
var _ = model.PH
var _ = msg.PH

{{range .Enums}}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [{{.Name}}] begin

{{if .Comment}}/*
{{.Comment}}
*/{{end}}

var {{.Name}}_Slice = []int32{
	{{range .Values}}{{.Value}},
	{{end}}
	}

func {{.Name}}_Len() int {
	return len({{.Name}}_Slice)
}

{{if .Enable0}}
func Check_{{.Name}}_I(value int32) bool {
	if _, ok := {{.Name}}_name[value]; ok {
		return true
	}
	return false
}
{{else}}
func Check_{{.Name}}_I(value int32) bool {
	if _, ok := {{.Name}}_name[value]; ok && value != 0 {
		return true
	}
	return false
}
{{end}}

func Check_{{.Name}}(value {{.Name}}) bool {
	return Check_{{.Name}}_I(int32(value))
}

func Each_{{.Name}}(f func({{.Name}}) bool) {
	for _, value := range {{.Name}}_Slice {
		if !f({{.Name}}(value)) {
			break
		}
	}
}

func Each_{{.Name}}_I(f func(int32) bool) {
	for _, value := range {{.Name}}_Slice {
		if !f(value) {
			break
		}
	}
}
// enum [{{.Name}}] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
{{end}}

{{range .Messages}}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [{{.Name}}] begin
func (m *{{.Name}}) ResetEx() {
{{range .Fields}}
	{{if .IsRepeated}}
        {{if eq .DescriptorProtoType "message"}}
            {{if .IsMap}}
            {{if and (not .ValueField.IsScalar) (not .ValueField.IsString) (not .ValueField.IsEnum)}}
            for _, i := range m.{{.Name}} {
				{{if (eq .ValueTypeToName "actor.PID")}}
					
                {{else}}
					{{$packageName := SplitPackageName .ValueTypeToName}}
					{{if $packageName}}
						{{$tYpE := SplitType .ValueTypeToName}}
						{{$packageName}}.Put_{{$tYpE}}(i)
					{{else}}
						Put_{{.ValueTypeToName}}(i)
					{{end}}
                {{end}}
            }
            {{end}}
		{{else}}
            for _, i := range m.{{.Name}} {
				{{$packageName := SplitPackageName .GoTypeToName}}
				{{if $packageName}}
					{{$tYpE := SplitType .GoTypeToName}}
					{{$packageName}}.Put_{{$tYpE}}(i)
				{{else}}
					Put_{{.GoTypeToName}}(i)
				{{end}}
            }
            {{end}}
        {{end}}
        //m.{{.Name}} = {{.GoType}}{}
        m.{{.Name}} = nil
	{{else}}
	    {{if or (eq .DescriptorProtoType "int32") (eq .DescriptorProtoType "uint32") (eq .DescriptorProtoType "int64") (eq .DescriptorProtoType "uint64") }}
	        m.{{.Name}} = 0
	    {{end}}
	    {{if or (eq .DescriptorProtoType "sint32") (eq .DescriptorProtoType "sint64") (eq .DescriptorProtoType "fixed32") (eq .DescriptorProtoType "fixed64") (eq .DescriptorProtoType "enum")}}
            m.{{.Name}} = 0
        {{end}}
	    {{if or (eq .DescriptorProtoType "bool")}}
            m.{{.Name}} = false
        {{end}}
        {{if or (eq .DescriptorProtoType "string")}}
            m.{{.Name}} = ""
        {{end}}
        {{if or (eq .DescriptorProtoType "double") (eq .DescriptorProtoType "float")}}
            m.{{.Name}} = 0.0
        {{end}}
        {{if or (eq .DescriptorProtoType "bytes") }}
            m.{{.Name}} = nil
        {{end}}
        {{if or (eq .DescriptorProtoType "message") }}
			if m.{{.Name}} != nil {
				m.{{.Name}}.ResetEx()
			} else {
				{{if (eq .GoTypeToName "actor.PID")}}
                m.{{.Name}} = &actor.PID{}
                {{else}}
                {{$packageName := SplitPackageName .GoTypeToName}}
                {{if $packageName}}
                {{$tYpE := SplitType .GoTypeToName}}
                m.{{.Name}} = {{$packageName}}.Get_{{$tYpE}}()
                {{else}}
                m.{{.Name}} = Get_{{.GoTypeToName}}()
                {{end}}
                {{end}}
			}
        {{end}}
	{{end}}
{{end}}
}

func (m {{.Name}}) Clone() *{{.Name}} {
	n, ok := g_{{.Name}}_Pool.Get().(*{{.Name}})
	if !ok || n == nil {
		n = &{{.Name}}{}
	}
{{range .Fields}}
    {{if (eq .DescriptorProtoType "message") }}
        {{if .IsRepeated}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Name}} = make({{.GoType}}, len(m.{{.Name}}))
                for i, e := range m.{{.Name}} {
                    {{if .IsMap}}
                    {{if and (not .ValueField.IsScalar) (not .ValueField.IsString) (not .ValueField.IsEnum)}}
                    if e != nil {
                        n.{{.Name}}[i] = e.Clone()
                    }
                    {{else}}
                    n.{{.Name}}[i] = e
                    {{end}}
                    {{else}}
                    if e != nil {
                        n.{{.Name}}[i] = e.Clone()
                    }
                    {{end}}
                }
            } else {
                //n.{{.Name}} = {{.GoType}}{}
                n.{{.Name}} = nil
            }
        {{else}}
            if m.{{.Name}} != nil {
            	n.{{.Name}} = m.{{.Name}}.Clone()
            }
        {{end}}
    {{else if (eq .DescriptorProtoType "bytes")}}
        {{if .IsRepeated}}
            if len(m.{{.Name}}) > 0 {
                for _, b := range m.{{.Name}} {
                    if len(b) > 0 {
                        nb := make([]byte, len(b))
                        copy(nb, b)
                        n.{{.Name}} = append(n.{{.Name}}, nb)
                    } else {
                        //n.{{.Name}} = append(n.{{.Name}}, []byte{})
                        n.{{.Name}} = append(n.{{.Name}}, nil)
                    }
                }
            } else {
                //n.{{.Name}} = [][]byte{}
                n.{{.Name}} = nil
            }
        {{else}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Name}} = make([]byte, len(m.{{.Name}}))
                copy(n.{{.Name}}, m.{{.Name}})
            } else {
                //n.{{.Name}} = []byte{}
                n.{{.Name}} = nil
            }
        {{end}}
    {{else}}
        {{if .IsRepeated}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Name}} = make([]{{.GoTypeToName}}, len(m.{{.Name}}))
                copy(n.{{.Name}}, m.{{.Name}})
            } else {
                //n.{{.Name}} = []{{.GoTypeToName}}{}
                n.{{.Name}} = nil
            }
        {{else}}
            n.{{.Name}} = m.{{.Name}}
        {{end}}
    {{end}}
{{end}}
	return n
}

func Clone_{{.Name}}_Slice(dst []*{{.Name}}, src []*{{.Name}}) []*{{.Name}} {
	for _, i := range dst {
		Put_{{.Name}}(i)
	}
	if len(src) > 0 {
		dst = make([]*{{.Name}}, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*{{.Name}}{}
		dst = nil
	}
	return dst
}

func (m {{.Name}}) JsonString() string {
	ba, _ := json.Marshal(m)
	return "{{.Name}}:" + string(ba)
}

{{if .ID}}
func ({{.Name}}) RPC() {
}

func ({{.Name}}) MessageID() rpc.MessageID {
	return {{.ID}}
}

func {{.Name}}_MessageID() rpc.MessageID {
	return {{.ID}}
}

func ({{.Name}}) MessageName() string {
	return "{{.Name}}"
}

func {{.Name}}_MessageName() string {
	return "{{.Name}}"
}

func ({{.Name}}) ResponseMessageID() rpc.MessageID {
{{if .Response}}
	return {{.Response}}_MessageID()
{{else}}
	return 0
{{end}}
}

func {{.Name}}_ResponseMessageID() rpc.MessageID {
{{if .Response}}
	return {{.Response}}_MessageID()
{{else}}
	return 0
{{end}}
}

{{end}}

{{if .Response}}
func Request_{{.Response}}(pid *actor.PID, send *{{.Name}}) (*{{.Response}}, error) {
	return Request_{{.Response}}_T(pid, send, {{.ResponseTimeout}} * time.Millisecond)
}

func Request_{{.Response}}_T(pid *actor.PID, send *{{.Name}}, timeout time.Duration) (*{{.Response}}, error) {
	if pid == nil {
		return nil, fmt.Errorf("pid is nil")
	}
	f := actor1.RootContext.RequestFuture(pid, send, timeout)
	iRecv, err := f.Result()
	if err != nil {
		return nil, err
	}
	recv, ok := iRecv.(*{{.Response}})
	if !ok {
		return nil, fmt.Errorf("recv %#v is not {{.Response}}", recv)
	}
	return recv, nil
}
{{end}}

func ({{.Name}}) CanIgnore() bool {
{{if .CanIgnore}} return true {{else}} return false {{end}}
}

func New_{{.Name}}() *{{.Name}} {
	m := &{{.Name}}{
{{range .Fields}}
    {{if (eq .DescriptorProtoType "message") }}
        {{if not .IsRepeated}}
			{{if (eq .GoTypeToName "actor.PID")}}
			{{.Name}}: &actor.PID{},
			{{else}}
			{{$packageName := SplitPackageName .GoTypeToName}}
			{{if $packageName}}
			{{$tYpE := SplitType .GoTypeToName}}
			{{.Name}}: {{$packageName}}.New_{{$tYpE}}(),
			{{else}}
			{{.Name}}: New_{{.GoTypeToName}}(),
			{{end}}
			{{end}}
        {{end}}
    {{end}}
{{end}}
	}
	return m
}

var g_{{.Name}}_Pool = sync.Pool{}

func Get_{{.Name}}() *{{.Name}} {
	m, ok := g_{{.Name}}_Pool.Get().(*{{.Name}})
	if !ok {
		m = New_{{.Name}}()
	} else {
		if m == nil {
			m = New_{{.Name}}()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_{{.Name}}(i interface{}) {
	if m, ok := i.(*{{.Name}}); ok && m != nil {
		g_{{.Name}}_Pool.Put(i)
	}
}

{{if .ID}}
func init() {
	Protocol.Register(
		&{{.Name}}{},
		func() rpc.IMessage { return Get_{{.Name}}() },
		func(msg rpc.IMessage) { Put_{{.Name}}(msg) },
	)
}
{{end}}
// message [{{.Name}}] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
{{end}}
`
