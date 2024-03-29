// 本文件由gen-static-data-go生成
// 请遵照提示添加修改！！！

package sd

import "encoding/json"
import "fmt"
import "log"

import "github.com/tealeg/xlsx"
import "github.com/trist725/mgsu/util"

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展import代码
//import_extend_begin
//import_extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////

const (
	G_C_Depose_Frag string = "5" // 普通品质分解获得碎片数
	G_C_xxx_Frag    string = "3" // 普通品质分解获得碎片数
)

type Global struct {
	ID int64 `excel_column:"0" excel_name:"id"` // 公共配置表ID

	Name string `excel_column:"1" excel_name:"name"` // 配置名

	Value string `excel_column:"2" excel_name:"value"` // 配置值

	Desc string `excel_column:"3" excel_name:"desc"` // 配置说明

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin
	//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func NewGlobal() *Global {
	sd := &Global{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin
	//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd Global) String() string {
	ba, _ := json.Marshal(sd)
	return string(ba)
}

func (sd Global) Clone() *Global {
	n := NewGlobal()
	*n = sd

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin
	//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return n
}

func (sd *Global) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
type GlobalManager struct {
	dataArray []*Global
	dataMap   map[int64]*Global

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager扩展字段
	//manager_extend_begin
	//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func newGlobalManager() *GlobalManager {
	mgr := &GlobalManager{
		dataArray: []*Global{},
		dataMap:   make(map[int64]*Global),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager的New代码
	//manager_new_begin
	//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *GlobalManager) Load(data []byte, fileName string) (success bool) {
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

		sd := NewGlobal()
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
		//struct_load_begin
		//struct_load_end
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
		//manager_load_begin
		//manager_load_end
		//////////////////////////////////////////////////////////////////////////////////////////////////
	}

	return
}

func (mgr GlobalManager) Size() int {
	return len(mgr.dataArray)
}

func (mgr GlobalManager) Get(id int64) *Global {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr GlobalManager) Each(f func(sd *Global) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *GlobalManager) each(f func(sd *Global) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr GlobalManager) findIf(f func(sd *Global) bool) *Global {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr GlobalManager) FindIf(f func(sd *Global) bool) *Global {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr GlobalManager) check(fileName string, row int, sd *Global) error {
	if _, ok := mgr.dataMap[sd.ID]; ok {
		return fmt.Errorf("%s 第%d行的id重复", fileName, row)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加检查代码
	//check_begin
	//check_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return nil
}

func (mgr *GlobalManager) AfterLoadAll(fileName string) (success bool) {
	success = true
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加加载后处理代码
	//after_load_all_begin
	//after_load_all_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展代码
//extend_begin
//extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////
