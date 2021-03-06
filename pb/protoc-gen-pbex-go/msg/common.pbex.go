// Code generated by protoc-gen-pbex-go. DO NOT EDIT IT!!!
// source: common.proto

package msg

import (
	json "encoding/json"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/trist725/mgsu/network/protocol/protobuf/v1"
	math "math"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_Ping] begin
func (m *C2S_Ping) ResetEx() {

	m.Content = ""

	m.D = 0.0

	m.F = 0.0

	m.I32 = 0

	m.I64 = 0

	m.Ui32 = 0

	m.Ui64 = 0

	m.B = false

	m.Bytes = nil

	if m.Info != nil {
		m.Info.ResetEx()
	} else {
		m.Info = Get_VersionInfo()
	}

	m.V = 0

	for _, i := range m.Infos {
		Put_VersionInfo(i)
	}

	//m.Infos = []*VersionInfo{}
	m.Infos = nil

	//m.I32Map = map[int32]int32{}
	m.I32Map = nil

	for _, i := range m.InfosMap {
		Put_VersionInfo(i)
	}

	//m.InfosMap = map[int32]*VersionInfo{}
	m.InfosMap = nil

	//m.VersionMap = map[int32]Version{}
	m.VersionMap = nil

}

func (m C2S_Ping) Clone() *C2S_Ping {
	n, ok := g_C2S_Ping_Pool.Get().(*C2S_Ping)
	if !ok || n == nil {
		n = &C2S_Ping{}
	}

	n.Content = m.Content

	n.D = m.D

	n.F = m.F

	n.I32 = m.I32

	n.I64 = m.I64

	n.Ui32 = m.Ui32

	n.Ui64 = m.Ui64

	n.B = m.B

	if len(m.Bytes) > 0 {
		n.Bytes = make([]byte, len(m.Bytes))
		copy(n.Bytes, m.Bytes)
	} else {
		//n.Bytes = []byte{}
		n.Bytes = nil
	}

	if m.Info != nil {
		n.Info = m.Info.Clone()
	}

	n.V = m.V

	if len(m.Infos) > 0 {
		n.Infos = make([]*VersionInfo, len(m.Infos))
		for i, e := range m.Infos {

			if e != nil {
				n.Infos[i] = e.Clone()
			}

		}
	} else {
		//n.Infos = []*VersionInfo{}
		n.Infos = nil
	}

	if len(m.I32Map) > 0 {
		n.I32Map = make(map[int32]int32, len(m.I32Map))
		for i, e := range m.I32Map {

			n.I32Map[i] = e

		}
	} else {
		//n.I32Map = map[int32]int32{}
		n.I32Map = nil
	}

	if len(m.InfosMap) > 0 {
		n.InfosMap = make(map[int32]*VersionInfo, len(m.InfosMap))
		for i, e := range m.InfosMap {

			if e != nil {
				n.InfosMap[i] = e.Clone()
			}

		}
	} else {
		//n.InfosMap = map[int32]*VersionInfo{}
		n.InfosMap = nil
	}

	if len(m.VersionMap) > 0 {
		n.VersionMap = make(map[int32]Version, len(m.VersionMap))
		for i, e := range m.VersionMap {

			n.VersionMap[i] = e

		}
	} else {
		//n.VersionMap = map[int32]Version{}
		n.VersionMap = nil
	}

	return n
}

func Clone_C2S_Ping_Slice(dst []*C2S_Ping, src []*C2S_Ping) []*C2S_Ping {
	for _, i := range dst {
		Put_C2S_Ping(i)
	}
	if len(src) > 0 {
		dst = make([]*C2S_Ping, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*C2S_Ping{}
		dst = nil
	}
	return dst
}

func (m C2S_Ping) JsonString() string {
	ba, _ := json.Marshal(m)
	return "C2S_Ping:" + string(ba)
}

func (C2S_Ping) V1() {
}

func (C2S_Ping) MessageID() v1.MessageID {
	return 1
}

func C2S_Ping_MessageID() v1.MessageID {
	return 1
}

func (C2S_Ping) ResponseMessageID() v1.MessageID {

	return S2C_Pong_MessageID()

}

func C2S_Ping_ResponseMessageID() v1.MessageID {

	return S2C_Pong_MessageID()

}

func New_C2S_Ping() *C2S_Ping {
	m := &C2S_Ping{

		Info: New_VersionInfo(),
	}
	return m
}

var g_C2S_Ping_Pool = sync.Pool{}

func Get_C2S_Ping() *C2S_Ping {
	m, ok := g_C2S_Ping_Pool.Get().(*C2S_Ping)
	if !ok {
		m = New_C2S_Ping()
	} else {
		if m == nil {
			m = New_C2S_Ping()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_Ping(i interface{}) {
	if m, ok := i.(*C2S_Ping); ok && m != nil {
		g_C2S_Ping_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&C2S_Ping{},
		func() v1.IMessage { return Get_C2S_Ping() },
		func(msg v1.IMessage) { Put_C2S_Ping(msg) },
	)
}

// message [C2S_Ping] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Pong] begin
func (m *S2C_Pong) ResetEx() {

	m.Content = ""

	m.D = 0.0

	m.F = 0.0

	m.I32 = 0

	m.I64 = 0

	m.Ui32 = 0

	m.Ui64 = 0

	m.B = false

	m.Bytes = nil

	if m.Info != nil {
		m.Info.ResetEx()
	} else {
		m.Info = Get_VersionInfo()
	}

	m.V = 0

	for _, i := range m.Infos {
		Put_VersionInfo(i)
	}

	//m.Infos = []*VersionInfo{}
	m.Infos = nil

	//m.I32Map = map[int32]int32{}
	m.I32Map = nil

	for _, i := range m.InfosMap {
		Put_VersionInfo(i)
	}

	//m.InfosMap = map[int32]*VersionInfo{}
	m.InfosMap = nil

	//m.VersionMap = map[int32]Version{}
	m.VersionMap = nil

}

func (m S2C_Pong) Clone() *S2C_Pong {
	n, ok := g_S2C_Pong_Pool.Get().(*S2C_Pong)
	if !ok || n == nil {
		n = &S2C_Pong{}
	}

	n.Content = m.Content

	n.D = m.D

	n.F = m.F

	n.I32 = m.I32

	n.I64 = m.I64

	n.Ui32 = m.Ui32

	n.Ui64 = m.Ui64

	n.B = m.B

	if len(m.Bytes) > 0 {
		n.Bytes = make([]byte, len(m.Bytes))
		copy(n.Bytes, m.Bytes)
	} else {
		//n.Bytes = []byte{}
		n.Bytes = nil
	}

	if m.Info != nil {
		n.Info = m.Info.Clone()
	}

	n.V = m.V

	if len(m.Infos) > 0 {
		n.Infos = make([]*VersionInfo, len(m.Infos))
		for i, e := range m.Infos {

			if e != nil {
				n.Infos[i] = e.Clone()
			}

		}
	} else {
		//n.Infos = []*VersionInfo{}
		n.Infos = nil
	}

	if len(m.I32Map) > 0 {
		n.I32Map = make(map[int32]int32, len(m.I32Map))
		for i, e := range m.I32Map {

			n.I32Map[i] = e

		}
	} else {
		//n.I32Map = map[int32]int32{}
		n.I32Map = nil
	}

	if len(m.InfosMap) > 0 {
		n.InfosMap = make(map[int32]*VersionInfo, len(m.InfosMap))
		for i, e := range m.InfosMap {

			if e != nil {
				n.InfosMap[i] = e.Clone()
			}

		}
	} else {
		//n.InfosMap = map[int32]*VersionInfo{}
		n.InfosMap = nil
	}

	if len(m.VersionMap) > 0 {
		n.VersionMap = make(map[int32]Version, len(m.VersionMap))
		for i, e := range m.VersionMap {

			n.VersionMap[i] = e

		}
	} else {
		//n.VersionMap = map[int32]Version{}
		n.VersionMap = nil
	}

	return n
}

func Clone_S2C_Pong_Slice(dst []*S2C_Pong, src []*S2C_Pong) []*S2C_Pong {
	for _, i := range dst {
		Put_S2C_Pong(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_Pong, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_Pong{}
		dst = nil
	}
	return dst
}

func (m S2C_Pong) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_Pong:" + string(ba)
}

func (S2C_Pong) V1() {
}

func (S2C_Pong) MessageID() v1.MessageID {
	return 2
}

func S2C_Pong_MessageID() v1.MessageID {
	return 2
}

func (S2C_Pong) ResponseMessageID() v1.MessageID {

	return 0

}

func S2C_Pong_ResponseMessageID() v1.MessageID {

	return 0

}

func New_S2C_Pong() *S2C_Pong {
	m := &S2C_Pong{

		Info: New_VersionInfo(),
	}
	return m
}

var g_S2C_Pong_Pool = sync.Pool{}

func Get_S2C_Pong() *S2C_Pong {
	m, ok := g_S2C_Pong_Pool.Get().(*S2C_Pong)
	if !ok {
		m = New_S2C_Pong()
	} else {
		if m == nil {
			m = New_S2C_Pong()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Pong(i interface{}) {
	if m, ok := i.(*S2C_Pong); ok && m != nil {
		g_S2C_Pong_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&S2C_Pong{},
		func() v1.IMessage { return Get_S2C_Pong() },
		func(msg v1.IMessage) { Put_S2C_Pong(msg) },
	)
}

// message [S2C_Pong] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_TestRepeated] begin
func (m *C2S_TestRepeated) ResetEx() {

	//m.Content = []string{}
	m.Content = nil

	//m.D = []float64{}
	m.D = nil

	//m.F = []float32{}
	m.F = nil

	//m.I32 = []int32{}
	m.I32 = nil

	//m.I64 = []int64{}
	m.I64 = nil

	//m.Ui32 = []uint32{}
	m.Ui32 = nil

	//m.Ui64 = []uint64{}
	m.Ui64 = nil

	//m.B = []bool{}
	m.B = nil

	//m.Bytes = [][]byte{}
	m.Bytes = nil

	for _, i := range m.Infos {
		Put_VersionInfo(i)
	}

	//m.Infos = []*VersionInfo{}
	m.Infos = nil

	//m.Vs = []Version{}
	m.Vs = nil

}

func (m C2S_TestRepeated) Clone() *C2S_TestRepeated {
	n, ok := g_C2S_TestRepeated_Pool.Get().(*C2S_TestRepeated)
	if !ok || n == nil {
		n = &C2S_TestRepeated{}
	}

	if len(m.Content) > 0 {
		n.Content = make([]string, len(m.Content))
		copy(n.Content, m.Content)
	} else {
		//n.Content = []string{}
		n.Content = nil
	}

	if len(m.D) > 0 {
		n.D = make([]float64, len(m.D))
		copy(n.D, m.D)
	} else {
		//n.D = []float64{}
		n.D = nil
	}

	if len(m.F) > 0 {
		n.F = make([]float32, len(m.F))
		copy(n.F, m.F)
	} else {
		//n.F = []float32{}
		n.F = nil
	}

	if len(m.I32) > 0 {
		n.I32 = make([]int32, len(m.I32))
		copy(n.I32, m.I32)
	} else {
		//n.I32 = []int32{}
		n.I32 = nil
	}

	if len(m.I64) > 0 {
		n.I64 = make([]int64, len(m.I64))
		copy(n.I64, m.I64)
	} else {
		//n.I64 = []int64{}
		n.I64 = nil
	}

	if len(m.Ui32) > 0 {
		n.Ui32 = make([]uint32, len(m.Ui32))
		copy(n.Ui32, m.Ui32)
	} else {
		//n.Ui32 = []uint32{}
		n.Ui32 = nil
	}

	if len(m.Ui64) > 0 {
		n.Ui64 = make([]uint64, len(m.Ui64))
		copy(n.Ui64, m.Ui64)
	} else {
		//n.Ui64 = []uint64{}
		n.Ui64 = nil
	}

	if len(m.B) > 0 {
		n.B = make([]bool, len(m.B))
		copy(n.B, m.B)
	} else {
		//n.B = []bool{}
		n.B = nil
	}

	if len(m.Bytes) > 0 {
		for _, b := range m.Bytes {
			if len(b) > 0 {
				nb := make([]byte, len(b))
				copy(nb, b)
				n.Bytes = append(n.Bytes, nb)
			} else {
				//n.Bytes = append(n.Bytes, []byte{})
				n.Bytes = append(n.Bytes, nil)
			}
		}
	} else {
		//n.Bytes = [][]byte{}
		n.Bytes = nil
	}

	if len(m.Infos) > 0 {
		n.Infos = make([]*VersionInfo, len(m.Infos))
		for i, e := range m.Infos {

			if e != nil {
				n.Infos[i] = e.Clone()
			}

		}
	} else {
		//n.Infos = []*VersionInfo{}
		n.Infos = nil
	}

	if len(m.Vs) > 0 {
		n.Vs = make([]Version, len(m.Vs))
		copy(n.Vs, m.Vs)
	} else {
		//n.Vs = []Version{}
		n.Vs = nil
	}

	return n
}

func Clone_C2S_TestRepeated_Slice(dst []*C2S_TestRepeated, src []*C2S_TestRepeated) []*C2S_TestRepeated {
	for _, i := range dst {
		Put_C2S_TestRepeated(i)
	}
	if len(src) > 0 {
		dst = make([]*C2S_TestRepeated, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*C2S_TestRepeated{}
		dst = nil
	}
	return dst
}

func (m C2S_TestRepeated) JsonString() string {
	ba, _ := json.Marshal(m)
	return "C2S_TestRepeated:" + string(ba)
}

func (C2S_TestRepeated) V1() {
}

func (C2S_TestRepeated) MessageID() v1.MessageID {
	return 3
}

func C2S_TestRepeated_MessageID() v1.MessageID {
	return 3
}

func (C2S_TestRepeated) ResponseMessageID() v1.MessageID {

	return 0

}

func C2S_TestRepeated_ResponseMessageID() v1.MessageID {

	return 0

}

func New_C2S_TestRepeated() *C2S_TestRepeated {
	m := &C2S_TestRepeated{}
	return m
}

var g_C2S_TestRepeated_Pool = sync.Pool{}

func Get_C2S_TestRepeated() *C2S_TestRepeated {
	m, ok := g_C2S_TestRepeated_Pool.Get().(*C2S_TestRepeated)
	if !ok {
		m = New_C2S_TestRepeated()
	} else {
		if m == nil {
			m = New_C2S_TestRepeated()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_TestRepeated(i interface{}) {
	if m, ok := i.(*C2S_TestRepeated); ok && m != nil {
		g_C2S_TestRepeated_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&C2S_TestRepeated{},
		func() v1.IMessage { return Get_C2S_TestRepeated() },
		func(msg v1.IMessage) { Put_C2S_TestRepeated(msg) },
	)
}

// message [C2S_TestRepeated] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [Item] begin
func (m *Item) ResetEx() {

	m.Uid = ""

	m.Stack = 0

	m.CreateTime = 0

}

func (m Item) Clone() *Item {
	n, ok := g_Item_Pool.Get().(*Item)
	if !ok || n == nil {
		n = &Item{}
	}

	n.Uid = m.Uid

	n.Stack = m.Stack

	n.CreateTime = m.CreateTime

	return n
}

func Clone_Item_Slice(dst []*Item, src []*Item) []*Item {
	for _, i := range dst {
		Put_Item(i)
	}
	if len(src) > 0 {
		dst = make([]*Item, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*Item{}
		dst = nil
	}
	return dst
}

func (m Item) JsonString() string {
	ba, _ := json.Marshal(m)
	return "Item:" + string(ba)
}

func New_Item() *Item {
	m := &Item{}
	return m
}

var g_Item_Pool = sync.Pool{}

func Get_Item() *Item {
	m, ok := g_Item_Pool.Get().(*Item)
	if !ok {
		m = New_Item()
	} else {
		if m == nil {
			m = New_Item()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Item(i interface{}) {
	if m, ok := i.(*Item); ok && m != nil {
		g_Item_Pool.Put(i)
	}
}

// message [Item] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Items] begin
func (m *S2C_Items) ResetEx() {

	if m.Item != nil {
		m.Item.ResetEx()
	} else {
		m.Item = Get_Item()
	}

	for _, i := range m.Items {
		Put_Item(i)
	}

	//m.Items = []*Item{}
	m.Items = nil

}

func (m S2C_Items) Clone() *S2C_Items {
	n, ok := g_S2C_Items_Pool.Get().(*S2C_Items)
	if !ok || n == nil {
		n = &S2C_Items{}
	}

	if m.Item != nil {
		n.Item = m.Item.Clone()
	}

	if len(m.Items) > 0 {
		n.Items = make([]*Item, len(m.Items))
		for i, e := range m.Items {

			if e != nil {
				n.Items[i] = e.Clone()
			}

		}
	} else {
		//n.Items = []*Item{}
		n.Items = nil
	}

	return n
}

func Clone_S2C_Items_Slice(dst []*S2C_Items, src []*S2C_Items) []*S2C_Items {
	for _, i := range dst {
		Put_S2C_Items(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_Items, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_Items{}
		dst = nil
	}
	return dst
}

func (m S2C_Items) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_Items:" + string(ba)
}

func (S2C_Items) V1() {
}

func (S2C_Items) MessageID() v1.MessageID {
	return 4
}

func S2C_Items_MessageID() v1.MessageID {
	return 4
}

func (S2C_Items) ResponseMessageID() v1.MessageID {

	return 0

}

func S2C_Items_ResponseMessageID() v1.MessageID {

	return 0

}

func New_S2C_Items() *S2C_Items {
	m := &S2C_Items{

		Item: New_Item(),
	}
	return m
}

var g_S2C_Items_Pool = sync.Pool{}

func Get_S2C_Items() *S2C_Items {
	m, ok := g_S2C_Items_Pool.Get().(*S2C_Items)
	if !ok {
		m = New_S2C_Items()
	} else {
		if m == nil {
			m = New_S2C_Items()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Items(i interface{}) {
	if m, ok := i.(*S2C_Items); ok && m != nil {
		g_S2C_Items_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&S2C_Items{},
		func() v1.IMessage { return Get_S2C_Items() },
		func(msg v1.IMessage) { Put_S2C_Items(msg) },
	)
}

// message [S2C_Items] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
