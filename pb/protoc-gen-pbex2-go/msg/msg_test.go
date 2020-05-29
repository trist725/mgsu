package msg

import (
	"bytes"
	"testing"

	"fmt"

	"time"

	"github.com/trist725/mgsu/util"
)

func TestProtocol(t *testing.T) {
	m := GetC2S_Ping()
	m.B = true
	m.Content = "hello protocol"
	m.Bytes = util.GenRandomByteArray(16)

	//fmt.Printf("%p=%#v\n", m, m)

	d, err := Protocol.Encode(m)
	if err != nil {
		t.Error(err)
		return
	}

	//fmt.Printf("%#v\n", d)

	i, err := Protocol.Decode(d)
	if err != nil {
		t.Error(err)
		return
	}

	m2, ok := i.(*C2S_Ping)
	if !ok {
		t.Error("convert to *C_Ping fail")
		return
	}

	//fmt.Printf("%p=%#v\n", m2, m2)

	if m2.B != m.B {
		t.Errorf("m2.B != m.B ")
		return
	}
	if m2.Content != m.Content {
		t.Error("m2.Content != m.Content")
		return
	}
	if !bytes.Equal(m2.Bytes, m.Bytes) {
		t.Error("m2.Bytes != m.Bytes")
		return
	}

	//fmt.Printf("m=[%#v]\nm2=[%#v]\n", m, m2)
}

func TestS2C_Items(t *testing.T) {
	now := time.Now().Unix()

	m := GetS2C_Items()
	m.Item.Uid = "S2C_Items.Item"
	m.Item.CreateTime = now

	for i := 0; i < 5; i++ {
		m.Items = append(m.Items, &Item{
			Uid:        fmt.Sprintf("S2C_Items.Items[%d]", i),
			CreateTime: now + int64(i),
		})
	}

	d, err := Protocol.Encode(m)
	if err != nil {
		t.Error(err)
		return
	}

	i, err := Protocol.Decode(d)
	if err != nil {
		t.Error(err)
		return
	}

	m2, ok := i.(*S2C_Items)
	if !ok {
		t.Error("convert to *S2C_Items fail")
		return
	}

	if m.Item.Uid != m2.Item.Uid {
		t.Errorf("m.Item.Uid[%s] != m2.Item.Uid[%s]", m.Item.Uid, m2.Item.Uid)
		return
	}

	//fmt.Printf("m.Item=[%#v] m2.Item=[%#v]\n", m.Item, m2.Item)

	for i := 0; i < len(m.Items); i++ {
		if m.Items[i].Uid != m2.Items[i].Uid {
			t.Errorf("m.Items[%d].Uid[%s] != m2.Items[%d].Uid[%s]", i, m.Items[i].Uid, i, m2.Items[i].Uid)
			return
		}
		//fmt.Printf("m.Items[%d]=[%#v] m2.Items[%d]=[%#v]\n", i, m.Items[i], i, m2.Items[i])
	}
}
