package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	TblCounters = "counters" // 用来生成递增序列的表
	TblUser     = "user"     // 角色表
)

// 定义库表的递增序列
var seqs = []string{
	TblUser,
}

// 定义库表的唯一索引
var uniqueIndexes = map[string][][]string{
	TblUser: {
		[]string{"ID"},
	},
}

// 定义库表的索引
var indexes = map[string][][]string{
	TblUser: {
		[]string{"VName"},
		[]string{"Token"},
	},
}

func NewObjectID() string {
	return primitive.NewObjectID().Hex()
}
