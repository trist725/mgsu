// 本文件由gen-static-data-go生成
// 请勿修改！！！

package sd

import "embed"

var (
	GlobalMgr = newGlobalManager()
	TestMgr   = newTestManager()
	Test2Mgr  = newTest2Manager()
)

//go:embed xlsx/global.xlsx
//go:embed xlsx/test.xlsx
//go:embed xlsx/test2.xlsx
var f embed.FS

func LoadAll() (success bool) {
	var data []byte
	success = true

	data, _ = f.ReadFile("xlsx/global.xlsx")
	success = GlobalMgr.Load(data, "global.xlsx") && success
	data, _ = f.ReadFile("xlsx/test.xlsx")
	success = TestMgr.Load(data, "test.xlsx") && success
	data, _ = f.ReadFile("xlsx/test2.xlsx")
	success = Test2Mgr.Load(data, "test2.xlsx") && success

	return
}

func AfterLoadAll() (success bool) {
	success = true

	success = GlobalMgr.AfterLoadAll("global.xlsx") && success
	success = TestMgr.AfterLoadAll("test.xlsx") && success
	success = Test2Mgr.AfterLoadAll("test2.xlsx") && success

	return
}
