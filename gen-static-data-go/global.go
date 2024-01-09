package main

// 全局表格元数据
type globalMeta struct {
	ID    int64
	Name  string
	Value string
	Desc  string
}

func newGlobalMeta() *globalMeta {
	m := &globalMeta{}
	return m
}
