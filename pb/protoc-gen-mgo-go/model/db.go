package model

import (
	"fmt"

	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

const (
	TblCounters = "counters" // 用来生成递增序列的表
	TblUser     = "user"     // 角色表
)

var (
	dbName      string
	dialContext *mongodb.DialContext
)

func Init(url string, sessionNum int, name string) (err error) {
	dialContext, err = mongodb.Dial(url, sessionNum)
	if err != nil {
		err = fmt.Errorf("connect to %s fail, %s", url, err)
		return
	}

	dbName = name

	err = dialContext.EnsureCounter(dbName, TblCounters, TblUser)
	if err != nil {
		err = fmt.Errorf("ensure counters error, %s", err)
		return
	}

	err = dialContext.EnsureUniqueIndex(dbName, TblUser, []string{"Name"})
	if err != nil {
		err = fmt.Errorf("ensure table user unique index error, %s", err)
		return
	}

	err = dialContext.EnsureIndex(dbName, TblUser, []string{"AccountID", "ServerID"})
	if err != nil {
		err = fmt.Errorf("ensure table user index error, %s", err)
		return
	}

	return
}

func Release() {
	if dialContext != nil {
		dialContext.Close()
		dialContext = nil
	}
}

func DialContext() *mongodb.DialContext {
	return dialContext
}

func GetSession() *mongodb.Session {
	return dialContext.Ref()
}

func PutSession(session *mongodb.Session) {
	dialContext.UnRef(session)
}

func NextSeq(id string) (int, error) {
	return dialContext.NextSeq(dbName, TblCounters, id)
}

func NewObjectId() string {
	return bson.NewObjectId().Hex()
}
