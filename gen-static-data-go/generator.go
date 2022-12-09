package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tealeg/xlsx"
	"github.com/trist725/mgsu/util"
)

type fieldMeta struct {
	Column          int    // 字段在xlsx中的列数
	Index           int    // 字段在结构体中的序列
	XlsxName        string // 字段在xlsx中的名字
	Name            string // 字段名
	TypeName        string // 字段类型名
	Type            int    // 字段类型
	ElementTypeName string // 字段元素类型名, 如果字段是数组或者二维数组, 则记录数组的元素类型, 否则等同于TypeName
	ElementType     int    // 字段元素类型, 如果字段是数组或者二维数组, 则记录数组的元素类型, 否则等同于Type
	Comment         string // 字段注释
	//IsArray         bool   // 是否数组
	//Is2DArray       bool   // 是否二维数组
}

func (fm fieldMeta) String() string {
	return fmt.Sprintf("%d: %s\t%s\t%s", fm.Column, fm.Name, fm.TypeName, fm.Comment)
}

func (fm *fieldMeta) checkAndFormat() error {
	fm.TypeName = strings.TrimLeft(fm.TypeName, "#")
	if fm.TypeName[0] == '*' {
		return fmt.Errorf("不支持指针类型")
	}

	fm.ElementTypeName = fm.TypeName

	words := strings.Split(fm.TypeName, "_")
	switch len(words) {
	case 1:
		if isGoBaseType(words[0]) {
			fm.TypeName = getGoBaseType(words[0])
			fm.Type = fieldTypeBase
			fm.ElementTypeName = fm.TypeName
			fm.ElementType = fieldTypeBase
		} else {
			fm.Type = fieldTypeStruct
			fm.ElementType = fieldTypeStruct
		}

	case 2:
		if !isGoArray(words[1]) {
			return fmt.Errorf("错误的次要修饰类型, 只能是arr, array, 2darr, 2darray")
		}

		fm.ElementTypeName = words[0]

		if isGoBaseType(fm.ElementTypeName) {
			fm.ElementTypeName = getGoBaseType(fm.ElementTypeName)
			fm.ElementType = fieldTypeBase
		} else {
			fm.ElementType = fieldTypeStruct
		}

		switch getGoArray(words[1]) {
		case "array":
			fm.TypeName = "[]" + fm.ElementTypeName
			fm.Type = fieldTypeArray
		case "2d_array":
			fm.TypeName = "[][]" + fm.ElementTypeName
			fm.Type = fieldType2dArray
		}

	default:
		return fmt.Errorf("错误的类型, 只能是基础类型, 对象类型, 数组或二维数组")
	}

	fm.Name = util.FormatFieldName(fm.XlsxName)

	fm.Comment = strings.TrimLeft(fm.Comment, "#")

	return nil
}

type staticDataMeta struct {
	ExcelFilePath     string
	ExcelFileBaseName string

	TotalColumnNum int
	ValidColumnNum int

	Name       string
	FieldMetas []fieldMeta

	///////////////////////////////////////////////////////////////////////
	// 以下字段记录非自动生成代码
	ImportExtend  string
	StructExtend  string
	StructNew     string
	StructClone   string
	StructLoad    string
	ManagerExtend string
	ManagerNew    string
	ManagerLoad   string
	Check         string
	AfterLoadAll  string
	Extend        string
	///////////////////////////////////////////////////////////////////////
}

func newStaticDataMeta() *staticDataMeta {
	sdm := &staticDataMeta{
		ImportExtend:  "\n",
		StructExtend:  "\n",
		StructNew:     "\n",
		StructClone:   "\n",
		StructLoad:    "\n",
		ManagerExtend: "\n",
		ManagerNew:    "\n",
		ManagerLoad:   "\n",
		Check:         "\n",
		AfterLoadAll:  "\n",
		Extend:        "\n",
	}
	return sdm
}

func (sdm staticDataMeta) String() string {
	// ba, _ := json.Marshal(sdm)
	// return string(ba)

	output := fmt.Sprintf("%s\n", sdm.Name)
	for _, fm := range sdm.FieldMetas {
		output += fm.String()
		output += "\n"
	}
	output += fmt.Sprintf("totalColumnNum: %d\nvalidColumnNum: %d", sdm.TotalColumnNum, sdm.ValidColumnNum)
	return output
}

func (sdm staticDataMeta) Import() string {
	i := `import "encoding/json"
import "fmt"
import "log"
`

	needImportTime := false
	for _, fm := range sdm.FieldMetas {
		if fm.TypeName == "time.Duration" {
			needImportTime = true
			break
		}
	}

	if needImportTime {
		i += `import "time"`
	}

	i += `
	import "github.com/tealeg/xlsx"
	import "github.com/trist725/mgsu/util"
`

	return i
}

func (sdm staticDataMeta) CloneFieldSourceCode() (sourceCode string) {
	for _, fm := range sdm.FieldMetas {
		switch fm.Type {
		case fieldTypeStruct:
			sourceCode += fmt.Sprintf(`
	n.%s = sd.%s.Clone()
`, fm.Name, fm.Name)

		case fieldTypeArray:
			switch fm.ElementType {
			case fieldTypeBase:
				sourceCode += fmt.Sprintf(`
	n.%s = make([]%s, len(sd.%s))
	copy(n.%s, sd.%s)
`, fm.Name, fm.ElementTypeName, fm.Name, fm.Name, fm.Name)

			case fieldTypeStruct:
				sourceCode += fmt.Sprintf(`
	n.%s = make([]%s, len(sd.%s))
	for i, e := range sd.%s {
		n.%s[i] = e.Clone()
	}
`, fm.Name, fm.ElementTypeName, fm.Name, fm.Name, fm.Name)
			}

		case fieldType2dArray:
			switch fm.ElementType {
			case fieldTypeBase:
				sourceCode += fmt.Sprintf(`
	n.%s = make([][]%s, len(sd.%s))
	for i, e := range sd.%s {
		n.%s[i] = make([]%s, len(e))
        copy(n.%s[i], e)
	}
`, fm.Name, fm.ElementTypeName, fm.Name, fm.Name, fm.Name, fm.ElementTypeName, fm.Name)

			case fieldTypeStruct:
				sourceCode += fmt.Sprintf(`
	n.%s = make([][]%s, len(sd.%s))
	for i, e := range sd.%s {
		n.%s[i] = make([]%s, len(e))
        for j, ee := range e {
			n.%s[i][j] = ee.Clone()
		}
	}
`, fm.Name, fm.ElementTypeName, fm.Name, fm.Name, fm.Name, fm.ElementTypeName, fm.Name)
			}
		}
	}
	return
}

func (sdm staticDataMeta) generate(absSourceCodeDir string, tpl *template.Template) error {
	var buffer bytes.Buffer

	sourceCodeBaseName := strings.TrimSuffix(filepath.Base(sdm.ExcelFilePath), ".xlsx")
	sourceCodeBaseName += ".sd.go"
	sourceCodePath := filepath.Join(absSourceCodeDir, sourceCodeBaseName)

	if err := util.IsDirOrFileExist(sourceCodePath); err == nil {
		old, err := ioutil.ReadFile(sourceCodePath)
		if err != nil {
			panic(err)
		}

		oldStr := string(old)

		begin := strings.Index(oldStr, "//import_extend_begin") + len("//import_extend_begin")
		end := strings.Index(oldStr, "//import_extend_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.ImportExtend = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//struct_extend_begin") + len("//struct_extend_begin")
		end = strings.Index(oldStr, "//struct_extend_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.StructExtend = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//struct_new_begin") + len("//struct_new_begin")
		end = strings.Index(oldStr, "//struct_new_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.StructNew = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//struct_clone_begin") + len("//struct_clone_begin")
		end = strings.Index(oldStr, "//struct_clone_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.StructClone = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//struct_load_begin") + len("//struct_load_begin")
		end = strings.Index(oldStr, "//struct_load_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.StructLoad = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//manager_extend_begin") + len("//manager_extend_begin")
		end = strings.Index(oldStr, "//manager_extend_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.ManagerExtend = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//manager_new_begin") + len("//manager_new_begin")
		end = strings.Index(oldStr, "//manager_new_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.ManagerNew = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//manager_load_begin") + len("//manager_load_begin")
		end = strings.Index(oldStr, "//manager_load_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.ManagerLoad = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//after_load_all_begin") + len("//after_load_all_begin")
		end = strings.Index(oldStr, "//after_load_all_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.AfterLoadAll = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//check_begin") + len("//check_begin")
		end = strings.Index(oldStr, "//check_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.Check = oldStr[begin:end]
		}

		begin = strings.Index(oldStr, "//extend_begin") + len("//extend_begin")
		end = strings.Index(oldStr, "//extend_end")
		if (begin > -1 && end > -1) && (begin != end) {
			sdm.Extend = oldStr[begin:end]
		}
	}

	if err := tpl.Execute(&buffer, sdm); err != nil {
		panicf("generate source code from template fail, %s", err)
	}

	ba, err := format.Source(buffer.Bytes())
	if err != nil {
		debugf(buffer.String())
		panicf("could't format source code, %s", err)
	}

	if err := ioutil.WriteFile(sourceCodePath, ba, 0666); err != nil {
		panic(err)
	}

	infof("-> %s", sourceCodePath)

	return nil
}

type staticDataCodeGenerator struct {
	cfg             *config
	tpl             *template.Template
	globalTpl       *template.Template
	StaticDataMetas []*staticDataMeta
}

func newStaticDataCodeGenerator(cfg *config) (sdcg *staticDataCodeGenerator, err error) {
	sdcg = &staticDataCodeGenerator{
		cfg: cfg,
	}

	sdcg.tpl, err = template.New("go").Parse(t)
	if err != nil {
		return
	}

	sdcg.globalTpl, err = template.New("go.global").Parse(globalT)
	if err != nil {
		return
	}

	return
}

func (sdcg *staticDataCodeGenerator) generate() error {
	absExcelDir, _ := filepath.Abs(sdcg.cfg.ExcelDir)
	absSourceCodeDir, _ := filepath.Abs(sdcg.cfg.SourceCodeDir)

	infof("扫描 %s", absExcelDir)

	sdcg.StaticDataMetas = []*staticDataMeta{}

	filepath.Walk(absExcelDir, func(path string, info os.FileInfo, err error) error {
		isXlsx, err := filepath.Match("*.xlsx", info.Name())
		if err != nil {
			return nil
		}
		if !isXlsx {
			return nil
		}

		dir, _ := filepath.Split(path)
		dir, _ = filepath.Abs(dir)
		if dir != absExcelDir {
			// 排除子目录下的xlsx文件
			return nil
		}

		xl, err := xlsx.OpenFile(path)
		if err != nil {
			return err
		}

		if len(xl.Sheets) == 0 {
			panicf("%s 没有分页", info.Name())
		}

		baseName := strings.TrimSuffix(info.Name(), ".xlsx")

		if len(sdcg.cfg.blackList) > 0 {
			if _, ok := sdcg.cfg.blackList[baseName]; ok {
				// 跳过黑名单的xlsx
				return nil
			}
		}

		if len(sdcg.cfg.whiteList) > 0 {
			if _, ok := sdcg.cfg.whiteList[baseName]; !ok {
				// 跳过不在白名单的xlsx
				return nil
			}
		}

		infof("<- %s", info.Name())

		// 默认取第一个sheet
		dataSheet, ok := xl.Sheet["data"]
		if !ok {
			panicf("%s 没有data分页", info.Name())
		}
		if len(dataSheet.Rows) < 3 {
			panicf("%s 的data分页的数据少于3行", info.Name())
		}

		nameRow := dataSheet.Rows[0]
		typeRow := dataSheet.Rows[1]
		commentRow := dataSheet.Rows[2]

		sdm := newStaticDataMeta()
		sdm.ExcelFilePath = filepath.Join(absExcelDir, path)
		sdm.ExcelFileBaseName = filepath.Base(sdm.ExcelFilePath)
		sdm.Name = util.FormatFieldName(baseName)

		for i := 0; i < len(typeRow.Cells); i++ {
			sdm.TotalColumnNum++

			fm := fieldMeta{
				Column: sdm.TotalColumnNum - 1,
			}

			if nameRow.Cells[i].Type() != xlsx.CellTypeString {
				continue
			}

			fm.XlsxName = nameRow.Cells[i].String()
			if fm.XlsxName == "" {
				continue
			}

			if []byte(fm.XlsxName)[0] == '#' {
				continue
			}

			//fm.Name, err = nameRow.Cells[i].String()
			//if err != nil || fm.Name == "" {
			//	continue
			//}

			if typeRow.Cells[i].Type() != xlsx.CellTypeString {
				continue
			}

			fm.TypeName = typeRow.Cells[i].String()
			if fm.TypeName == "" {
				continue
			}

			if commentRow.Cells[i].Type() != xlsx.CellTypeString {
				continue
			}

			fm.Comment = commentRow.Cells[i].String()

			sdm.ValidColumnNum++

			if err := fm.checkAndFormat(); err != nil {
				panicf("%s 格式化出错, %s", info.Name(), err)
			}

			fm.Index = sdm.ValidColumnNum

			sdm.FieldMetas = append(sdm.FieldMetas, fm)
		}

		//fmt.Println(sdm)

		sdcg.StaticDataMetas = append(sdcg.StaticDataMetas, sdm)

		return nil
	})

	for _, sdm := range sdcg.StaticDataMetas {
		// 生成每个xlsx对应的go代码
		if err := sdm.generate(absSourceCodeDir, sdcg.tpl); err != nil {
			return err
		}
	}

	var buffer bytes.Buffer
	if err := sdcg.globalTpl.Execute(&buffer, sdcg); err != nil {
		panicf("生成代码失败, %s", err)
	}

	ba, err := format.Source(buffer.Bytes())
	if err != nil {
		debugf(buffer.String())
		panicf("格式化代码失败, %s", err)
	}

	sourceFilePath := filepath.Join(absSourceCodeDir, "static_data.sd.go")
	if err := ioutil.WriteFile(sourceFilePath, ba, 0666); err != nil {
		panic(err)
	}

	infof("-> %s", sourceFilePath)

	return nil
}
