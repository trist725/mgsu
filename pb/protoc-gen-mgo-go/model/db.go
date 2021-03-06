package model

import (
	"fmt"

	"github.com/globalsign/mgo/bson"

	"github.com/trist725/mgsu/db/mongodb"
)

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
var SC = NewSimpleClient()

type SimpleClient struct {
	url        string
	sessionNum int
	dbName     string

	dialContext *mongodb.DialContext
}

func NewSimpleClient() (sc *SimpleClient) {
	sc = &SimpleClient{}
	return
}

func (sc *SimpleClient) Init(url string, sessionNum int, dbName string) (err error) {
	sc.url = url
	sc.sessionNum = sessionNum
	sc.dbName = dbName

	sc.dialContext, err = mongodb.Dial(sc.url, sc.sessionNum)
	if err != nil {
		err = fmt.Errorf("connect to %s fail, %s", sc.url, err)
		return
	}

	for _, seq := range seqs {
		err = sc.dialContext.EnsureCounter(sc.dbName, TblCounters, seq)
		if err != nil {
			err = fmt.Errorf("ensure counters [%s] error, %s", seq, err)
			return
		}
	}

	for tbl, indexes := range uniqueIndexes {
		for _, index := range indexes {
			err = sc.dialContext.EnsureUniqueIndex(sc.dbName, tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] unique index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	for tbl, is := range indexes {
		for _, index := range is {
			err = sc.dialContext.EnsureIndex(sc.dbName, tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	return
}

func (sc *SimpleClient) Release() {
	if sc.dialContext != nil {
		sc.dialContext.Close()
		sc.dialContext = nil
	}
}

func (sc *SimpleClient) DialContext() *mongodb.DialContext {
	return sc.dialContext
}

func (sc SimpleClient) DBName() string {
	return sc.dbName
}

func (sc *SimpleClient) GetSession() *mongodb.Session {
	return sc.dialContext.Ref()
}

func (sc *SimpleClient) PutSession(session *mongodb.Session) {
	sc.dialContext.UnRef(session)
}

func (sc *SimpleClient) NextSeq(id string) (int, error) {
	return sc.dialContext.NextSeq(sc.dbName, TblCounters, id)
}

func NewObjectID() string {
	return bson.NewObjectId().Hex()
}
