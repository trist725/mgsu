package mgo

const DefaultTemplate = `
var _ = json.Marshal
var _ = msg.PH

{{range .Messages}}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [{{.Name}}] begin

func New_{{.Name}}() *{{.Name}} {
	m := &{{.Name}}{
{{range .Fields}}
    {{if (eq .DescriptorProtoType "message") }}
        {{if not .IsRepeated}}
            {{.Name}}: Get_{{.GoTypeToName}}(),
        {{end}}
    {{end}}
{{end}}
	}
	return m
}

func (m {{.Name}}) JsonString() string {
	ba, _ := json.Marshal(m)
	return "{{.Name}}:" + string(ba)
}

func (m *{{.Name}}) ResetEx() {
{{range .Fields}}
    {{if .IsRepeated}}
        {{if eq .DescriptorProtoType "message"}}
            {{if .IsMap}}
            {{if and (not .ValueField.IsScalar) (not .ValueField.IsString) (not .ValueField.IsEnum)}}
            for _, i := range m.{{.Name}} {
                Put_{{.ValueTypeToName}}(i)
            }
            {{end}}
            {{else}}
            for _, i := range m.{{.Name}} {
                Put_{{.GoTypeToName}}(i)
            }
            {{end}}
        {{end}}
        // m.{{.Name}} = {{.GoType}}{}
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
				m.{{.Name}} = Get_{{.GoTypeToName}}()
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
        {{if .IsRepeated -}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Name}} = make({{.GoType}}, len(m.{{.Name -}}))
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
                // n.{{.Name}} = {{.GoType}}{}
                n.{{.Name}} = nil
            }
        {{else}}
            if m.{{.Name}} != nil {
            	n.{{.Name}} = m.{{.Name}}.Clone()
            }
        {{end}}
    {{else if (eq .DescriptorProtoType "bytes")}}
        {{if .IsRepeated -}}
            if len(m.{{.Name}}) > 0 {
                for _, b := range m.{{.Name}} {
                    if len(b) > 0 {
                        nb := make([]byte, len(b))
                        copy(nb, b)
                        n.{{.Name}} = append(n.{{.Name}}, nb)
                    } else {
                        // n.{{.Name}} = append(n.{{.Name}}, []byte{})
                        n.{{.Name}} = append(n.{{.Name}}, nil)
                    }
                }
            } else {
                // n.{{.Name}} = [][]byte{}
                n.{{.Name}} = nil
            }
        {{else -}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Name}} = make([]byte, len(m.{{.Name}}))
                copy(n.{{.Name}}, m.{{.Name}})
            } else {
                // n.{{.Name}} = []byte{}
                n.{{.Name}} = nil
            }
        {{end}}
    {{else}}
        {{if .IsRepeated -}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Name}} = make([]{{.GoTypeToName}}, len(m.{{.Name}}))
                copy(n.{{.Name}}, m.{{.Name}})
            } else {
                // n.{{.Name}} = []{{.GoTypeToName}}{}
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
		// dst = []*{{.Name}}{}
		dst = nil
	}
	return dst
}

{{if .ID}}
func (sc SimpleClient)FindOne_{{.Name}}(ctx context.Context, query interface{}, opts ...options.FindOptions) (one *{{.Name}}, err error) {
	one = Get_{{.Name}}()
	err = sc.cli.Database.Collection(Tbl{{.Name}}).Find(ctx, query, opts...).One(one)
	if err != nil {
		Put_{{.Name}}(one)
		return nil, err
	}
	return
}

func (sc SimpleClient)FindSome_{{.Name}}(ctx context.Context, query interface{}, opts ...options.FindOptions) (some []*{{.Name}}, err error) {
	some = []*{{.Name}}{}
	err = sc.cli.Database.Collection(Tbl{{.Name}}).Find(ctx, query, opts...).All(&some)
	if err != nil {
		return nil, err
	}
	return
}

func (sc SimpleClient)UpdateSome_{{.Name}}(ctx context.Context, selector interface{}, update interface{}, opts ...options.UpdateOptions) (result *qmgo.UpdateResult, err error) {
	return sc.cli.Database.Collection(Tbl{{.Name}}).UpdateAll(ctx, selector, update, opts...)
}

func (sc SimpleClient)Upsert_{{.Name}}(ctx context.Context, selector interface{}, update interface{}, opts ...options.UpsertOptions) (result *qmgo.UpdateResult, err error) {
	return sc.cli.Database.Collection(Tbl{{.Name}}).Upsert(ctx, selector, update, opts...)
}

func (sc SimpleClient)UpsertByObjID_{{.Name}}(ctx context.Context, objID interface{}, update interface{}, opts ...options.UpsertOptions) (result *qmgo.UpdateResult, err error) {
	return sc.cli.Database.Collection(Tbl{{.Name}}).UpsertId(ctx, objID, update, opts...)
}

func (m {{.Name}}) Insert(ctx context.Context, opts ...options.InsertOneOptions) (result *qmgo.InsertOneResult, err error) {
	return SC.cli.Database.Collection(Tbl{{.Name}}).InsertOne(ctx, m, opts...)
}

func (m {{.Name}}) Update(ctx context.Context, selector interface{}, update interface{}, opts ...options.UpdateOptions) (err error) {
	return SC.cli.Database.Collection(Tbl{{.Name}}).UpdateOne(ctx, selector, update, opts...)
}

func (m {{.Name}}) Upsert(ctx context.Context, selector interface{}, opts ...options.UpsertOptions) (result *qmgo.UpdateResult, err error) {
	return SC.cli.Database.Collection(Tbl{{.Name}}).Upsert(ctx, selector, m, opts...)
}

func (m {{.Name}}) UpdateByObjID(ctx context.Context, opts ...opts.UpdateOptions) (err error) {
	return SC.cli.Database.Collection(Tbl{{.Name}}).UpdateId(context.Background(), m.ObjID, bson.D{ {"$set", m} }, opts...)
}

func (m {{.Name}}) RemoveByObjID(ctx context.Context, opts ...opts.RemoveOptions) error {
	return SC.cli.Database.Collection(Tbl{{.Name}}).RemoveId(context.Background(), m.ObjID, opts...)
}
{{end}}

{{if .Msg}}
func (m {{.Name}}) ToMsg(n *msg.{{.Msg}}) *msg.{{.Msg}} {
    if n == nil {
        n = msg.Get_{{.Msg}}()
    }
{{range .Fields}}
{{if .Msg}}
    {{if .IsMessage }}
        {{- if .IsRepeated -}}
            if len(m.{{.Name}}) > 0 {
				{{if .IsMap}}
					{{- if and (not .ValueField.IsScalar) (not .ValueField.IsString) (not .ValueField.IsEnum) -}}
						n.{{.Msg}} = make({{.MsgRepeatedMessageGoType}}, len(m.{{.Name}}))
						for i, e := range m.{{.Name}} {
							if e != nil {
								n.{{.Msg}}[i] = e.ToMsg(n.{{.Msg}}[i])
							} else {
								n.{{.Msg}}[i] = msg.Get_{{.ValueTypeToName}}()
							}
						}
					{{- else if .ValueField.IsEnum -}}
						n.{{.Msg}} = make({{.MsgRepeatedMessageGoType}}, len(m.{{.Name}}))
						for i, e := range m.{{.Name}} {
							n.{{.Msg}}[i] = msg.{{.ValueTypeToName}}(e)
						}
					{{- else -}}
						n.{{.Msg}} = make({{.GoType}}, len(m.{{.Name}}))
						for i, e := range m.{{.Name}} {
							n.{{.Msg}}[i] = e
						}
					{{- end -}}
				{{- else -}}
					{{- if and (not .IsScalar) (not .IsString) (not .IsEnum) -}}
						n.{{.Msg}} = make({{.MsgRepeatedMessageGoType}}, len(m.{{.Name}}))
						for i, e := range m.{{.Name}} {
							if e != nil {
								n.{{.Msg}}[i] = e.ToMsg(n.{{.Msg}}[i])
							} else {
								n.{{.Msg}}[i] = msg.Get_{{.GoTypeToName}}()
							}
						}
					{{- else -}}
						n.{{.Msg}} = make({{.GoType}}, len(m.{{.Name}}))
						for i, e := range m.{{.Name}} {
							n.{{.Msg}}[i] = e
						}
					{{- end -}}
				{{- end -}}
            } else {
                // n.{{.Msg}} = {{.GoType}}{}
                n.{{.Msg}} = nil
            }
        {{- else -}}
            if m.{{.Msg}} != nil {
                n.{{.Msg}} = m.{{.Name}}.ToMsg(n.{{.Msg}})
            }
        {{end}}
    {{else if .IsBytes}}
        {{if .IsRepeated -}}
            if len(m.{{.Name}}) > 0 {
                for _, b := range m.{{.Name -}} {
                    if len(b) > 0 {
                        nb := make([]byte, len(b))
                        copy(nb, b)
                        n.{{.Msg}} = append(n.{{.Msg}}, nb)
                    } else {
                        // n.{{.Msg}} = append(n.{{.Msg}}, []byte{})
                        n.{{.Msg}} = append(n.{{.Msg}}, nil)
                    }
                }
            } else {
                // n.{{.Msg}} = [][]byte{}
                n.{{.Msg}} = nil
            }
        {{- else -}}
            if len(m.{{.Name}}) > 0 {
                n.{{.Msg}} = make([]byte, len(m.{{.Name}}))
                copy(n.{{.Msg}}, m.{{.Name}})
            } else {
                // n.{{.Msg}} = []byte{}
                n.{{.Msg}} = nil
            }
        {{- end -}}
    {{else}}
        {{if .IsRepeated -}}
			if len(m.{{.Name}}) > 0 {
			{{if .IsEnumSlice -}}
				for _, e := range m.{{.Name}} {
					n.{{.Msg}} = append(n.{{.Msg}}, msg.{{.GoTypeToName}}(e))  
				}
			{{- else -}}
				n.{{.Msg}} = make([]{{.GoTypeToName}}, len(m.{{.Name}}))
				copy(n.{{.Msg}}, m.{{.Name}})
			{{- end -}}
            } else {
                // n.{{.Msg}} = []{{.GoTypeToName}}{}
                n.{{.Msg}} = nil
            }
		{{else if .IsEnum}}
			n.{{.Msg}} = msg.{{.GoType}}(m.{{.Name}})
        {{else}}
            n.{{.Msg}} = m.{{.Name}}
        {{- end}}
    {{end}}
{{end}}
{{end}}
    return n
}
{{end}}

{{$Msg := .}}
{{$MsgName := .Msg}}

{{range .Maps}}

type {{.Name}} map[{{.Key}}]*{{$Msg.Name}}

func To{{.Name}}(m map[{{.Key}}]*{{$Msg.Name}}) *{{.Name}} {
	if m == nil {
		return nil
	}
	return (*{{.Name}})(&m)
}

func New{{.Name}}() (m *{{.Name}}) {
	m = &{{.Name}}{}
	return
}

func (m *{{.Name}}) Get(key {{.Key}}) (value *{{$Msg.Name}}, ok bool) {
	value, ok = (*m)[key]
	return
}

func (m *{{.Name}}) Set(key {{.Key}}, value *{{$Msg.Name}}) {
	(*m)[key] = value
}

func (m *{{.Name}}) Add(key {{.Key}}) (value *{{$Msg.Name}}) {
	value = Get_{{$Msg.Name}}()
	(*m)[key] = value
	return
}

func (m *{{.Name}}) Remove(key {{.Key}}) (removed bool) {
	if _, ok := (*m)[key]; ok {
		delete(*m, key)
		return true
	}
	return false
}

func (m *{{.Name}}) RemoveOne(fn func(key {{.Key}}, value *{{$Msg.Name}}) (removed bool)) {
	for key, value := range *m {
		if fn(key, value) {
			delete(*m, key)
			break
		}
	}
}

func (m *{{.Name}}) RemoveSome(fn func(key {{.Key}}, value *{{$Msg.Name}}) (removed bool)) {
	left := map[{{.Key}}]*{{$Msg.Name}}{}
	for key, value := range *m {
		if !fn(key, value) {
			left[key] = value
		}
	}
	*m = left
}

func (m *{{.Name}}) Each(f func(key {{.Key}}, value *{{$Msg.Name}}) (continued bool)) {
	for key, value := range *m {
		if !f(key, value) {
			break
		}
	}
}

func (m {{.Name}}) Size() int {
	return len(m)
}

func (m {{.Name}}) Clone() (n *{{.Name}}) {
	if m.Size() == 0 {
		return nil
	}
	n = To{{.Name}}(make(map[{{.Key}}]*{{$Msg.Name}}, m.Size()))
	for k, v := range m {
		if v != nil {
			(*n)[k] = v.Clone()
		} else {
			(*n)[k] = nil
		}
	}
	return n
}

func (m *{{.Name}}) Clear() {
	*m = *New{{.Name}}()
}

{{if $MsgName}}
func (m {{.Name}}) ToMsg(n map[{{.Key}}]*msg.{{$MsgName}}) map[{{.Key}}]*msg.{{$MsgName}} {
	if m.Size() == 0 {
		return nil
	}
	n = make(map[{{.Key}}]*msg.{{$MsgName}}, m.Size())
	for k, v := range m {
		if v != nil {
			n[k] = v.ToMsg(nil)
		} else {
			n[k] = msg.Get_{{$Msg.Name}}()
		}
	}
	return n
}
{{end}}

{{end}}

{{if .Slice}}
type {{.Name}}Slice []*{{.Name}}

func New{{.Name}}Slice() *{{.Name}}Slice {
	return &{{.Name}}Slice{}
}

func To{{.Name}}Slice(s []*{{.Name}}) *{{.Name}}Slice {
	return (*{{.Name}}Slice)(&s)
}

func (s *{{.Name}}Slice) Add() *{{.Name}} {
	return s.AddOne(Get_{{.Name}}())
}

func (s *{{.Name}}Slice) AddOne(newOne *{{.Name}}) *{{.Name}}  {
	*s = append(*s, newOne)
	return newOne
}

func (s *{{.Name}}Slice) RemoveOne(fn func(index int, element *{{.Name}}) (removed bool)) {
	for i, e := range *s {
		if fn(i, e) {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}

func (s *{{.Name}}Slice) RemoveSome(fn func(index int, element *{{.Name}}) (removed bool)) {
	var left []*{{.Name}}
	for i, e := range *s {
		if !fn(i, e) {
			left = append(left, e)
		}
	}
	*s = left
}

func (s {{.Name}}Slice) Each(fn func(index int, element *{{.Name}}) (continued bool)) {
	for i, e := range s {
		if !fn(i, e) {
			break
		}
	}
}

func (s {{.Name}}Slice) Size() int {
	return len(s)
}

func (s {{.Name}}Slice) Clone() (n *{{.Name}}Slice) {
	if s.Size() == 0 {
		return nil
	}
	n = To{{.Name}}Slice(make([]*{{.Name}}, s.Size()))
	for i, e := range s {
		if e != nil {
			(*n)[i] = e.Clone()
		} else {
			(*n)[i] = nil
		}
	}
	return n
}

func (s *{{.Name}}Slice) Clear() {
	*s = *New{{.Name}}Slice() 
}

{{if .Msg}}
func (s {{.Name}}Slice) ToMsg(n []*msg.{{.Name}}) []*msg.{{.Name}} {
	if s.Size() == 0 {
		return nil
	}
	n = make([]*msg.{{.Name}}, s.Size())
	for i, e := range s {
		if e != nil {
			n[i] = e.ToMsg(nil)
		} else {
			n[i] = msg.Get_{{.Name}}()
		}
	}
	return n
}
{{end}}
{{end}}

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

// collection [{{.Name}}] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
{{end}}
`
