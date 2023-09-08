// 本文件由gen-static-data-go生成
// 请遵照提示添加修改！！！

package sd

import "encoding/json"
import "fmt"
import "log"
import "time"
import "github.com/tealeg/xlsx"
import "github.com/trist725/mgsu/util"

//////////////////////////////////////////////////////////////////////////////////////////////////
// TODO 添加扩展import代码
//import_extend_begin
//import_extend_end
//////////////////////////////////////////////////////////////////////////////////////////////////

type Test struct {
	ID int64 `excel_column:"1" excel_name:"id"` // 整型

	Tb bool `excel_column:"2" excel_name:"tb"` // sfsg

	Time time.Duration `excel_column:"4" excel_name:"time"` // 时间

	IntArray []int `excel_column:"5" excel_name:"int_array"` // 整型数组

	Int2dArray [][]int `excel_column:"6" excel_name:"int_2d_array"` // 二维整型数组

	StringArray []string `excel_column:"7" excel_name:"string_array"` // 字符串数组

	String2dArray [][]string `excel_column:"8" excel_name:"string_2d_array"` // 二维字符串数组

	Float float64 `excel_column:"9" excel_name:"float"` // 浮点

	FloatArr []float64 `excel_column:"10" excel_name:"float_arr"` // 浮点数组

	Float2dArr [][]float64 `excel_column:"11" excel_name:"float_2d_arr"` // 浮点二维数组

	Reward Reward `excel_column:"12" excel_name:"reward"` // 对象2

	RewardArr []Reward `excel_column:"13" excel_name:"reward_arr"` // 对象2数组

	Reward2darr [][]Reward `excel_column:"14" excel_name:"reward_2darr"` // 对象2二维数组

	Item Item `excel_column:"15" excel_name:"item"` // 对象

	ItemArr []Item `excel_column:"16" excel_name:"item_arr"` // 对象数组

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体扩展字段
	//struct_extend_begin
	//struct_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func NewTest() *Test {
	sd := &Test{}
	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体New代码
	//struct_new_begin
	//struct_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
	return sd
}

func (sd Test) String() string {
	ba, _ := json.Marshal(sd)
	return string(ba)
}

func (sd Test) Clone() *Test {
	n := NewTest()
	*n = sd

	n.IntArray = make([]int, len(sd.IntArray))
	copy(n.IntArray, sd.IntArray)

	n.Int2dArray = make([][]int, len(sd.Int2dArray))
	for i, e := range sd.Int2dArray {
		n.Int2dArray[i] = make([]int, len(e))
		copy(n.Int2dArray[i], e)
	}

	n.StringArray = make([]string, len(sd.StringArray))
	copy(n.StringArray, sd.StringArray)

	n.String2dArray = make([][]string, len(sd.String2dArray))
	for i, e := range sd.String2dArray {
		n.String2dArray[i] = make([]string, len(e))
		copy(n.String2dArray[i], e)
	}

	n.FloatArr = make([]float64, len(sd.FloatArr))
	copy(n.FloatArr, sd.FloatArr)

	n.Float2dArr = make([][]float64, len(sd.Float2dArr))
	for i, e := range sd.Float2dArr {
		n.Float2dArr[i] = make([]float64, len(e))
		copy(n.Float2dArr[i], e)
	}

	n.Reward = sd.Reward.Clone()

	n.RewardArr = make([]Reward, len(sd.RewardArr))
	for i, e := range sd.RewardArr {
		n.RewardArr[i] = e.Clone()
	}

	n.Reward2darr = make([][]Reward, len(sd.Reward2darr))
	for i, e := range sd.Reward2darr {
		n.Reward2darr[i] = make([]Reward, len(e))
		for j, ee := range e {
			n.Reward2darr[i][j] = ee.Clone()
		}
	}

	n.Item = sd.Item.Clone()

	n.ItemArr = make([]Item, len(sd.ItemArr))
	for i, e := range sd.ItemArr {
		n.ItemArr[i] = e.Clone()
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加结构体Clone代码
	//struct_clone_begin
	//struct_clone_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return n
}

func (sd *Test) load(row *xlsx.Row) error {
	return util.DeserializeStructFromXlsxRow(sd, row)
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
type TestManager struct {
	dataArray []*Test
	dataMap   map[int64]*Test

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager扩展字段
	//manager_extend_begin
	//manager_extend_end
	//////////////////////////////////////////////////////////////////////////////////////////////////
}

func newTestManager() *TestManager {
	mgr := &TestManager{
		dataArray: []*Test{},
		dataMap:   make(map[int64]*Test),
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO 添加manager的New代码
	//manager_new_begin
	//manager_new_end
	//////////////////////////////////////////////////////////////////////////////////////////////////

	return mgr
}

func (mgr *TestManager) Load(data []byte, fileName string) (success bool) {
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

		sd := NewTest()
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

func (mgr TestManager) Size() int {
	return len(mgr.dataArray)
}

func (mgr TestManager) Get(id int64) *Test {
	sd, ok := mgr.dataMap[id]
	if !ok {
		return nil
	}
	return sd.Clone()
}

func (mgr TestManager) Each(f func(sd *Test) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd.Clone()) {
			break
		}
	}
}

func (mgr *TestManager) each(f func(sd *Test) bool) {
	for _, sd := range mgr.dataArray {
		if !f(sd) {
			break
		}
	}
}

func (mgr TestManager) findIf(f func(sd *Test) bool) *Test {
	for _, sd := range mgr.dataArray {
		if f(sd) {
			return sd
		}
	}
	return nil
}

func (mgr TestManager) FindIf(f func(sd *Test) bool) *Test {
	for _, sd := range mgr.dataArray {
		n := sd.Clone()
		if f(n) {
			return n
		}
	}
	return nil
}

func (mgr TestManager) check(fileName string, row int, sd *Test) error {
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

func (mgr *TestManager) AfterLoadAll(fileName string) (success bool) {
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
