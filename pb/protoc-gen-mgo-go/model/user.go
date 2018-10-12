package model

import (
	"time"
)

func CreateUser(accountID int64, serverID int32, name string, sex int32) (m *User, err error) {
	nextSeq, err := NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	m = Get_User()
	m.ID = int64(nextSeq)*10000 + int64(serverID)
	m.AccountID = accountID
	m.ServerID = serverID
	m.Name = name
	m.Sex = sex
	m.CreateTime = time.Now().Unix()
	return
}
