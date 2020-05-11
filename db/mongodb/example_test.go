package mongodb_test

import (
	"testing"

	"github.com/globalsign/mgo"

	"github.com/trist725/mgsu/db/mongodb"
)

func TestExample(t *testing.T) {
	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		t.Error(err)
		return
	}
	defer c.Close()

	// session
	s := c.Ref()
	defer c.UnRef(s)
	err = s.DB("test").C("counters").RemoveId("test")
	if err != nil && err != mgo.ErrNotFound {
		t.Error(err)
		return
	}

	// auto increment
	err = c.EnsureCounter("test", "counters", "test")
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 3; i++ {
		id, err := c.NextSeq("test", "counters", "test")
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(id)
	}

	// index
	err = c.EnsureUniqueIndex("test", "counters", []string{"key1"})
	if err != nil {
		t.Error(err)
		return
	}

	// Output:
	// 1
	// 2
	// 3
}
