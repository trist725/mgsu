package model

import (
	"testing"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	url = "mongodb://admin:admin111@192.168.101.100:27017/admin"
)

func testUser(t *testing.T) {
	if err := Init(url, 1, "game-test"); err != nil {
		t.Error(err)
		return
	}
	defer Release()

	session := GetSession()
	defer PutSession(session)

	newUser, err := CreateUser(1, 1, "test", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if err := newUser.Insert(session); err != nil {
		t.Error(err)
		return
	}

	_, err = FindOne_User(session, bson.M{"_id": newUser.ID})
	if err != nil {
		t.Error(err)
	}

	_, err = FindOne_User(session, bson.M{"Name": "test"})
	if err != nil {
		t.Error(err)
	}

	some, err := FindSome_User(session, bson.M{"Name": "test"})
	if err != nil || len(some) != 1 {
		t.Error(err)
	}

	err = newUser.RemoveByID(session)
	if err != nil {
		t.Error(err)
	}
}

func TestClone(t *testing.T) {
	var dst = []*User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	}

	fmt.Println("before, dst=")
	for _, i := range dst {
		fmt.Println(i)
	}

	var src = []*User{
		{ID: 2},
		{ID: 3},
		{ID: 4},
		{ID: 5},
		{ID: 6},
	}

	fmt.Println("src=")
	for _, i := range src {
		fmt.Println(i)
	}

	fmt.Printf("dst=%#v, src=%#v\n", dst, src)

	dst = Clone_User_Slice(dst, src)

	fmt.Println("after, dst=")
	for _, i := range dst {
		fmt.Println(i)
	}

	fmt.Println("src=")
	for _, i := range src {
		fmt.Println(i)
	}

	fmt.Printf("dst=%#v, src=%#v\n", dst, src)
}
