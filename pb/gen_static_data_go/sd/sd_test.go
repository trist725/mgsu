package sd

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	if !LoadAll("..") {
		t.Error("加载静态数据失败")
		return
	}

	if !AfterLoadAll("..") {
		t.Error("加载全部静态数据后失败")
		return
	}

	rawSd := TestMgr.dataMap[1]
	fmt.Printf("rawSd=%v\n", rawSd)
	fmt.Printf("rawSd.Reward2darr len=[%d] address=[%p]\n", len(rawSd.Reward2darr), &rawSd.Reward2darr)
	fmt.Printf("rawSd.Reward2darr[0][0] address=[%p]\n", &rawSd.Reward2darr[0][0])

	sd := TestMgr.Get(1)
	if sd == nil {
		t.Error("not found id=1 sd")
		return
	}

	fmt.Printf("sd=%v\n", sd)
	fmt.Printf("sd.Reward2darr len=[%d] address=[%p]\n", len(sd.Reward2darr), &sd.Reward2darr)
	fmt.Printf("sd.Reward2darr[0][0] address=[%p]\n", &sd.Reward2darr[0][0])
}
