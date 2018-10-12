package main

import (
	"testing"

	"gitee.com/nggs/tools/gen_static_data_go/sd"
)

func TestSourceCodeGenerate(t *testing.T) {
	sdcg, err := newStaticDataCodeGenerator(gCfg)
	if err != nil {
		t.Error(err)
		return
	}
	if err := sdcg.generate(); err != nil {
		t.Error(err)
		return
	}
}

func TestLoad(t *testing.T) {
	if !sd.LoadAll(".") {
		t.Error("加载静态数据失败")
		return
	}

	if !sd.AfterLoadAll(".") {
		t.Error("加载全部静态数据后失败")
		return
	}

	t.Logf("size=[%d]\n", sd.TestMgr.Size())

	sd.TestMgr.Each(func(sd *sd.Test) bool {
		t.Log(sd.String())
		return true
	})

	sd.Test2Mgr.Each(func(sd *sd.Test2) bool {
		t.Log(sd.String())
		return true
	})
}
