// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: user.proto

/*
It has these top-level messages:
	User
*/

package model

import "fmt"
import "encoding/json"
import "sync"
import "github.com/name5566/leaf/db/mongodb"
import "gopkg.in/mgo.v2"

var _ = fmt.Sprintf
var _ = json.Marshal
var _ *sync.Pool
var _ *mongodb.DialContext
var _ *mgo.DBRef

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [User] begin

/// 用户数据 @collection
type User struct {
	/// 用户id @bson=_id
	ID int64 `bson:"_id"`
	/// 帐号id @bson=accountid
	AccountID int64 `bson:"accountid"`
	/// 服务器ID
	ServerID int32 `bson:"ServerID"`
	/// 名字
	Name string `bson:"Name"`
	/// 性别
	Sex int32 `bson:"Sex"`
	/// 创建时刻
	CreateTime int64 `bson:"CreateTime"`
	/// 测试数组
	Arr []int32 `bson:"Arr"`
}

func New_User() *User {
	m := &User{
		Arr: []int32{},
	}
	return m
}

func (m User) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"User\":%s}", string(ba))
}

func (m *User) Reset() {
	m.ID = 0
	m.AccountID = 0
	m.ServerID = 0
	m.Name = ""
	m.Sex = 0
	m.CreateTime = 0
	m.Arr = []int32{}

}

func (m User) Clone() *User {
	n, ok := g_User_Pool.Get().(*User)
	if !ok || n == nil {
		n = &User{}
	}

	n.ID = m.ID
	n.AccountID = m.AccountID
	n.ServerID = m.ServerID
	n.Name = m.Name
	n.Sex = m.Sex
	n.CreateTime = m.CreateTime

	if len(m.Arr) > 0 {
		n.Arr = make([]int32, len(m.Arr))
		copy(n.Arr, m.Arr)
	} else {
		n.Arr = []int32{}
	}

	return n
}

func Clone_User_Slice(dst []*User, src []*User) []*User {
	for _, i := range dst {
		Put_User(i)
	}
	dst = []*User{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func FindOne_User(session *mongodb.Session, query interface{}) (one *User, err error) {
	one = Get_User()
	err = session.DB(dbName).C(TblUser).Find(query).One(one)
	if err != nil {
		Put_User(one)
		return nil, err
	}
	return
}

func FindSome_User(session *mongodb.Session, query interface{}) (some []*User, err error) {
	some = []*User{}
	err = session.DB(dbName).C(TblUser).Find(query).All(&some)
	if err != nil {
		return nil, err
	}
	return
}

func UpdateSome_User(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblUser).UpdateAll(selector, update)
	return
}

func Upsert_User(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblUser).Upsert(selector, update)
	return
}

func UpsertID_User(session *mongodb.Session, id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblUser).UpsertId(id, update)
	return
}

func (m User) Insert(session *mongodb.Session) error {
	return session.DB(dbName).C(TblUser).Insert(m)
}

func (m User) Update(session *mongodb.Session, selector interface{}, update interface{}) error {
	return session.DB(dbName).C(TblUser).Update(selector, update)
}

func (m User) UpdateByID(session *mongodb.Session) error {
	return session.DB(dbName).C(TblUser).UpdateId(m.ID, m)
}

func (m User) RemoveByID(session *mongodb.Session) error {
	return session.DB(dbName).C(TblUser).RemoveId(m.ID)
}

var g_User_Pool = sync.Pool{}

func Get_User() *User {
	m, ok := g_User_Pool.Get().(*User)
	if !ok {
		m = New_User()
	} else {
		if m == nil {
			m = New_User()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_User(i interface{}) {
	if m, ok := i.(*User); ok && m != nil {
		g_User_Pool.Put(i)
	}
}

// collection [User] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
