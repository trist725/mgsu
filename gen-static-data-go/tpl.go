package main

const t = `// 本文件由gen-static-data-go生成
// 请遵照提示添加修改！！！

package sd

{{.Import}}
//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展import代码
//import_extend_begin{{.ImportExtend}}//import_extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////

type {{.Name}} struct {
    {{range .FieldMetas}}
    {{.Name}} {{.TypeName}} ` + "`excel_column:\"{{.Column}}\" excel_name:\"{{.XlsxName}}\"`" + ` // {{.Comment}}
	{{end}}
	
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin{{.StructExtend}}//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func New{{.Name}}() *{{.Name}} {
	sd := &{{.Name}}{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin{{.StructNew}}//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd {{.Name}}) String() string {
    ba, _ := json.Marshal(sd)
    return string(ba)
}

func (sd {{.Name}}) Clone() *{{.Name}} {
    n := New{{.Name}}()
    *n = sd
	{{.CloneFieldSourceCode}}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin{{.StructClone}}//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

    return n
}

func (sd *{{.Name}}) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
type {{.Name}}Manager struct {
	dataArray []*{{.Name}}
	dataMap   map[int64]*{{.Name}}

	//////////////////////////////////////////////////////////////////////////////////////////////////
    // TODO 添加manager扩展字段
	//manager_extend_begin{{.ManagerExtend}}//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func new{{.Name}}Manager() *{{.Name}}Manager {
	mgr := &{{.Name}}Manager{
		dataArray: []*{{.Name}}{},
		dataMap:   make(map[int64]*{{.Name}}),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
    // TODO 添加manager的New代码
	//manager_new_begin{{.ManagerNew}}//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *{{.Name}}Manager) Load(data []byte, fileName string) (success bool) {
	success = true

	xl, err := xlsx.OpenBinary(data)
	if err != nil {
		log.Printf("打开 %s 失败, %s\n", fileName, err)
		return false
	}

	if len(xl.Sheets) == 0 {
		log.Printf("%s 没有分页可加载\n", fileName)
		return false
	}

	dataSheet, ok := xl.Sheet["data"]
	if !ok {
		log.Printf("%s 没有data分页\n", fileName)
		return false
	}

	if len(dataSheet.Rows) < 3 {
		log.Printf("%s 数据少于3行\n", fileName)
		return false
	}

	for i := 3; i < len(dataSheet.Rows); i++ {
		row := dataSheet.Rows[i]
		if len(row.Cells) == 0 {
		    continue
		}

		firstColumn := row.Cells[0]
		firstComment := firstColumn.String()
		if firstComment != "" {
		    if firstComment[0] == '#' {
		        // 跳过被注释掉的行
		        continue
		    }
		}

		sd := New{{.Name}}()
		err = sd.load(row)
		if err != nil {
			log.Printf("%s 加载第%d行失败, %s\n", fileName, i+1, err)
			success = false
			continue
		}

		if sd.ID == 0 {
			continue
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////
        // TODO 添加结构体加载代码
		//struct_load_begin{{.StructLoad}}//struct_load_end
		//////////////////////////////////////////////////////////////////////////////////////////////////

		if err := mgr.check(fileName, i+1, sd); err != nil {
			log.Println(err)
			success = false
			continue
		}

		mgr.dataArray = append(mgr.dataArray, sd)
		mgr.dataMap[sd.ID] = sd

		//////////////////////////////////////////////////////////////////////////////////////////////////
        // TODO 添加manager加载代码
		//manager_load_begin{{.ManagerLoad}}//manager_load_end
		//////////////////////////////////////////////////////////////////////////////////////////////////
	}

	return
}

func (mgr {{.Name}}Manager) Size() int {
	return len(mgr.dataArray)
}

func (mgr {{.Name}}Manager) Get(id int64) *{{.Name}} {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr {{.Name}}Manager) Each(f func(sd *{{.Name}}) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *{{.Name}}Manager) each(f func(sd *{{.Name}}) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr {{.Name}}Manager) findIf(f func(sd *{{.Name}}) bool) *{{.Name}} {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr {{.Name}}Manager) FindIf(f func(sd *{{.Name}}) bool) *{{.Name}} {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr {{.Name}}Manager) check(fileName string, row int, sd *{{.Name}}) error {
	if _, ok := mgr.dataMap[sd.ID]; ok {
		return fmt.Errorf("%s 第%d行的id重复", fileName, row)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加检查代码
	//check_begin{{.Check}}//check_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return nil
}

func (mgr *{{.Name}}Manager) AfterLoadAll(fileName string) (success bool) {
	success = true
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加加载后处理代码
	//after_load_all_begin{{.AfterLoadAll}}//after_load_all_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展代码
//extend_begin{{.Extend}}//extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////
`

const globalT = `// 本文件由gen-static-data-go生成
// 请勿修改！！！

package sd

import "embed"

var (
    {{range .StaticDataMetas}}{{.Name}}Mgr = new{{.Name}}Manager()
    {{end}}
)

{{range .StaticDataMetas}}
//go:embed xlsx/{{.ExcelFileBaseName -}}
{{end}}
var f embed.FS

func LoadAll() (success bool) {
	var data []byte
	success = true

    {{range .StaticDataMetas}}
	data, _ = f.ReadFile("xlsx/{{.ExcelFileBaseName}}")
	success = {{.Name}}Mgr.Load(data, "{{.ExcelFileBaseName}}") && success
    {{- end}}

	return
}

func AfterLoadAll() (success bool) {
	success = true

	{{range .StaticDataMetas}}
	success = {{.Name}}Mgr.AfterLoadAll("{{.ExcelFileBaseName}}") && success
    {{- end}}

	return
}
`
