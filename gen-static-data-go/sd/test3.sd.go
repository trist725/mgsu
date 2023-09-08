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

type Test3 struct {
	ID int64 `excel_column:"1" excel_name:"id"` // 船只ID

	Name string `excel_column:"3" excel_name:"name"` // 名字

	Icon string `excel_column:"4" excel_name:"icon"` // 头像资源

	Quality int `excel_column:"5" excel_name:"quality"` // 品质

	Type int `excel_column:"6" excel_name:"type"` // 类型

	Str int `excel_column:"7" excel_name:"str"` // 船只初始攻击

	SailorStr int `excel_column:"8" excel_name:"sailor_str"` // 水手初始攻击

	Def int `excel_column:"9" excel_name:"def"` // 船只初始耐久

	Sailors int `excel_column:"10" excel_name:"sailors"` // 船只搭载水手数量

	SailorLimit int `excel_column:"11" excel_name:"sailor_limit"` // 出港保底水手数

	Power int `excel_column:"12" excel_name:"power"` // 船只初始推进力

	Carrying int `excel_column:"13" excel_name:"carrying"` // 船只初始载货量

	StrLimit int `excel_column:"14" excel_name:"str_limit"` // 船只攻击改造上限值

	DefLimit int `excel_column:"15" excel_name:"def_limit"` // 船只耐久改造上限值

	SailorsLimit int `excel_column:"16" excel_name:"sailors_limit"` // 船只搭载水手数量改造上限值

	PowerLimit int `excel_column:"17" excel_name:"power_limit"` // 船只推进力改造上限值

	CarryingLimit int `excel_column:"18" excel_name:"carrying_limit"` // 船只载货量改造上限值

	ItemTypeID int `excel_column:"19" excel_name:"item_type_id"` // 材料表中对应专属订造材料id

	MakeMoney int `excel_column:"20" excel_name:"make_money"` // 订造价格

	TransformationMoney int `excel_column:"21" excel_name:"transformation_money"` // 改造价格

	SellMoney int `excel_column:"22" excel_name:"sell_money"` // 出售价格

	FormulaShipID int `excel_column:"23" excel_name:"formula_ship_id"` // 对应合成表中船图纸id

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin
	//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func NewTest3() *Test3 {
	sd := &Test3{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin
	//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd Test3) String() string {
	ba, _ := json.Marshal(sd)
	return string(ba)
}

func (sd Test3) Clone() *Test3 {
	n := NewTest3()
	*n = sd

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin
	//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return n
}

func (sd *Test3) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
type Test3Manager struct {
	dataArray []*Test3
	dataMap   map[int64]*Test3

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager扩展字段
	//manager_extend_begin
	//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func newTest3Manager() *Test3Manager {
	mgr := &Test3Manager{
		dataArray: []*Test3{},
		dataMap:   make(map[int64]*Test3),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager的New代码
	//manager_new_begin
	//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *Test3Manager) Load(data []byte, fileName string) (success bool) {
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

		sd := NewTest3()
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

func (mgr Test3Manager) Size() int {
	return len(mgr.dataArray)
}

func (mgr Test3Manager) Get(id int64) *Test3 {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr Test3Manager) Each(f func(sd *Test3) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *Test3Manager) each(f func(sd *Test3) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr Test3Manager) findIf(f func(sd *Test3) bool) *Test3 {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr Test3Manager) FindIf(f func(sd *Test3) bool) *Test3 {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr Test3Manager) check(fileName string, row int, sd *Test3) error {
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

func (mgr *Test3Manager) AfterLoadAll(fileName string) (success bool) {
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
