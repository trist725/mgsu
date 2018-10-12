// 本文件由gen_static_data_go生成
// 请勿修改！！！

package sd

import "log"
import "path/filepath"

var (
	TestMgr  = newTestManager()
	Test2Mgr = newTest2Manager()
)

func LoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = TestMgr.Load(filepath.Join(absExcelDir, "test.xlsx")) && success

	success = Test2Mgr.Load(filepath.Join(absExcelDir, "test2.xlsx")) && success

	return
}

func AfterLoadAll(excelDir string) (success bool) {
	absExcelDir, err := filepath.Abs(excelDir)
	if err != nil {
		log.Println(err)
		return false
	}

	success = true

	success = TestMgr.AfterLoadAll(filepath.Join(absExcelDir, "test.xlsx")) && success

	success = Test2Mgr.AfterLoadAll(filepath.Join(absExcelDir, "test2.xlsx")) && success

	return
}
