package sd

import (
	"testing"
)

func TestClone(t *testing.T) {
	if !LoadAll() {
		t.Error("加载静态数据失败")
		return
	}

	if !AfterLoadAll() {
		t.Error("加载全部静态数据后失败")
		return
	}

	rawSd := TestMgr.dataMap[1]
	t.Logf("rawSd=%v\n", rawSd)
	t.Logf("rawSd.Reward2darr len=[%d] address=[%p]\n", len(rawSd.Reward2darr), &rawSd.Reward2darr)
	t.Logf("rawSd.Reward2darr[0][0] address=[%p]\n", &rawSd.Reward2darr[0][0])

	sd := TestMgr.Get(1)
	if sd == nil {
		t.Error("not found id=1 sd")
		return
	}

	t.Logf("sd=%v\n", sd)
	t.Logf("sd.Reward2darr len=[%d] address=[%p]\n", len(sd.Reward2darr), &sd.Reward2darr)
	t.Logf("sd.Reward2darr[0][0] address=[%p]\n", &sd.Reward2darr[0][0])
}
