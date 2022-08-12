package model

import (
	"context"
	"testing"

	"github.com/qiniu/qmgo"

	"github.com/globalsign/mgo/bson"
)

const (
	uri = "mongodb://127.0.0.1:27017/?connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-256"
)

func TestUser(t *testing.T) {
	if err := SC.Init(context.Background(), uri, &qmgo.Config{Uri: uri, Database: "db_test", Coll: "user"}); err != nil {
		t.Error(err)
		return
	}
	defer SC.Release()

	newUser, err := SC.CreateUser(1, 1, "test", 1)
	if err != nil {
		t.Error(err)
		return
	}

	newUser.IMap = map[int32]int32{
		1: 1,
		2: 2,
		3: 3,
	}

	newUser.TestMap = map[int32]*Test{
		1: {
			I32: 1,
		},
		2: {
			I32: 2,
		},
	}

	if _, err := newUser.Insert(); err != nil {
		t.Error(err)
		return
	}

	u, err := SC.FindOne_User(bson.M{"_id": newUser.ID})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v\n", u)

	_, err = SC.FindOne_User(bson.M{"name": "test"})
	if err != nil {
		t.Error(err)
	}

	some, err := SC.FindSome_User(bson.M{"name": "test"})
	if err != nil || len(some) != 1 {
		t.Error(err)
	}

	err = newUser.RemoveByID()
	if err != nil {
		t.Error(err)
	}
}

func TestClone_User_Slice(t *testing.T) {
	var dst = []*User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	}

	t.Logf("before, dst=\n")
	for _, i := range dst {
		t.Log(i)
	}

	var src = []*User{
		{ID: 2},
		{ID: 3},
		{ID: 4},
		{ID: 5},
		{ID: 6},
	}

	t.Logf("src=\n")
	for _, i := range src {
		t.Log(i)
	}

	t.Logf("dst=%#v\nsrc=%#v\n", dst, src)

	dst = Clone_User_Slice(dst, src)

	t.Logf("after clone, dst=\n")
	for _, i := range dst {
		t.Log(i)
	}

	t.Logf("src=\n")
	for _, i := range src {
		t.Log(i)
	}

	t.Logf("dst=%#v\nsrc=%#v\n", dst, src)
}
