// Code generated by protoc-gen-pbex2-go. DO NOT EDIT IT!!!
// source: common.proto

/*
It has these top-level messages:
	C2S_Ping
	S2C_Pong
	C2S_TestRepeated
	Item
	S2C_Items
*/

package msg

import "sync"
import protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"

var _ *sync.Pool
var _ = protocol.PH

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
	m.Bytes = []byte{}
	m.Info.ResetEx()
	m.V = 0

	for _, i := range m.Infos {
		Put_VersionInfo(i)
	}
	m.Infos = []*VersionInfo{}

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
		n.Bytes = []byte{}
	}

	n.Info = m.Info.Clone()
	n.V = m.V

	if len(m.Infos) > 0 {
		for _, i := range m.Infos {
			if i != nil {
				n.Infos = append(n.Infos, i.Clone())
			} else {
				n.Infos = append(n.Infos, nil)
			}
		}
	} else {
		n.Infos = []*VersionInfo{}
	}

	return n
}

func Clone_C2S_Ping_Slice(dst []*C2S_Ping, src []*C2S_Ping) []*C2S_Ping {
	for _, i := range dst {
		Put_C2S_Ping(i)
	}
	dst = []*C2S_Ping{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (C2S_Ping) V2() {
}

func (C2S_Ping) MessageID() protocol.MessageID {
	return "msg.C2S_Ping"
}

func C2S_Ping_MessageID() protocol.MessageID {
	return "msg.C2S_Ping"
}

func New_C2S_Ping() *C2S_Ping {
	m := &C2S_Ping{
		Bytes: []byte{},
		Info:  New_VersionInfo(),
		Infos: []*VersionInfo{},
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
		func() protocol.IMessage { return Get_C2S_Ping() },
		func(msg protocol.IMessage) { Put_C2S_Ping(msg) },
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
	m.Bytes = []byte{}
	m.Info.ResetEx()
	m.V = 0

	for _, i := range m.Infos {
		Put_VersionInfo(i)
	}
	m.Infos = []*VersionInfo{}

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
		n.Bytes = []byte{}
	}

	n.Info = m.Info.Clone()
	n.V = m.V

	if len(m.Infos) > 0 {
		for _, i := range m.Infos {
			if i != nil {
				n.Infos = append(n.Infos, i.Clone())
			} else {
				n.Infos = append(n.Infos, nil)
			}
		}
	} else {
		n.Infos = []*VersionInfo{}
	}

	return n
}

func Clone_S2C_Pong_Slice(dst []*S2C_Pong, src []*S2C_Pong) []*S2C_Pong {
	for _, i := range dst {
		Put_S2C_Pong(i)
	}
	dst = []*S2C_Pong{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (S2C_Pong) V2() {
}

func (S2C_Pong) MessageID() protocol.MessageID {
	return "msg.S2C_Pong"
}

func S2C_Pong_MessageID() protocol.MessageID {
	return "msg.S2C_Pong"
}

func New_S2C_Pong() *S2C_Pong {
	m := &S2C_Pong{
		Bytes: []byte{},
		Info:  New_VersionInfo(),
		Infos: []*VersionInfo{},
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
		func() protocol.IMessage { return Get_S2C_Pong() },
		func(msg protocol.IMessage) { Put_S2C_Pong(msg) },
	)
}

// message [S2C_Pong] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_TestRepeated] begin
func (m *C2S_TestRepeated) ResetEx() {
	m.Content = []string{}
	m.D = []float64{}
	m.F = []float32{}
	m.I32 = []int32{}
	m.I64 = []int64{}
	m.Ui32 = []uint32{}
	m.Ui64 = []uint64{}
	m.B = []bool{}
	m.Bytes = [][]byte{}

	for _, i := range m.Infos {
		Put_VersionInfo(i)
	}
	m.Infos = []*VersionInfo{}
	m.Vs = []Version{}

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
		n.Content = []string{}
	}

	if len(m.D) > 0 {
		n.D = make([]float64, len(m.D))
		copy(n.D, m.D)
	} else {
		n.D = []float64{}
	}

	if len(m.F) > 0 {
		n.F = make([]float32, len(m.F))
		copy(n.F, m.F)
	} else {
		n.F = []float32{}
	}

	if len(m.I32) > 0 {
		n.I32 = make([]int32, len(m.I32))
		copy(n.I32, m.I32)
	} else {
		n.I32 = []int32{}
	}

	if len(m.I64) > 0 {
		n.I64 = make([]int64, len(m.I64))
		copy(n.I64, m.I64)
	} else {
		n.I64 = []int64{}
	}

	if len(m.Ui32) > 0 {
		n.Ui32 = make([]uint32, len(m.Ui32))
		copy(n.Ui32, m.Ui32)
	} else {
		n.Ui32 = []uint32{}
	}

	if len(m.Ui64) > 0 {
		n.Ui64 = make([]uint64, len(m.Ui64))
		copy(n.Ui64, m.Ui64)
	} else {
		n.Ui64 = []uint64{}
	}

	if len(m.B) > 0 {
		n.B = make([]bool, len(m.B))
		copy(n.B, m.B)
	} else {
		n.B = []bool{}
	}

	if len(m.Bytes) > 0 {
		for _, b := range m.Bytes {
			if len(b) > 0 {
				nb := make([]byte, len(b))
				copy(nb, b)
				n.Bytes = append(n.Bytes, nb)
			} else {
				n.Bytes = append(n.Bytes, []byte{})
			}
		}
	} else {
		n.Bytes = [][]byte{}
	}

	if len(m.Infos) > 0 {
		for _, i := range m.Infos {
			if i != nil {
				n.Infos = append(n.Infos, i.Clone())
			} else {
				n.Infos = append(n.Infos, nil)
			}
		}
	} else {
		n.Infos = []*VersionInfo{}
	}

	if len(m.Vs) > 0 {
		n.Vs = make([]Version, len(m.Vs))
		copy(n.Vs, m.Vs)
	} else {
		n.Vs = []Version{}
	}

	return n
}

func Clone_C2S_TestRepeated_Slice(dst []*C2S_TestRepeated, src []*C2S_TestRepeated) []*C2S_TestRepeated {
	for _, i := range dst {
		Put_C2S_TestRepeated(i)
	}
	dst = []*C2S_TestRepeated{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (C2S_TestRepeated) V2() {
}

func (C2S_TestRepeated) MessageID() protocol.MessageID {
	return "msg.C2S_TestRepeated"
}

func C2S_TestRepeated_MessageID() protocol.MessageID {
	return "msg.C2S_TestRepeated"
}

func New_C2S_TestRepeated() *C2S_TestRepeated {
	m := &C2S_TestRepeated{
		Content: []string{},
		D:       []float64{},
		F:       []float32{},
		I32:     []int32{},
		I64:     []int64{},
		Ui32:    []uint32{},
		Ui64:    []uint64{},
		B:       []bool{},
		Bytes:   [][]byte{},
		Infos:   []*VersionInfo{},
		Vs:      []Version{},
	}
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
		func() protocol.IMessage { return Get_C2S_TestRepeated() },
		func(msg protocol.IMessage) { Put_C2S_TestRepeated(msg) },
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
	dst = []*Item{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
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
	m.Item.ResetEx()

	for _, i := range m.Items {
		Put_Item(i)
	}
	m.Items = []*Item{}

}

func (m S2C_Items) Clone() *S2C_Items {
	n, ok := g_S2C_Items_Pool.Get().(*S2C_Items)
	if !ok || n == nil {
		n = &S2C_Items{}
	}

	n.Item = m.Item.Clone()

	if len(m.Items) > 0 {
		for _, i := range m.Items {
			if i != nil {
				n.Items = append(n.Items, i.Clone())
			} else {
				n.Items = append(n.Items, nil)
			}
		}
	} else {
		n.Items = []*Item{}
	}

	return n
}

func Clone_S2C_Items_Slice(dst []*S2C_Items, src []*S2C_Items) []*S2C_Items {
	for _, i := range dst {
		Put_S2C_Items(i)
	}
	dst = []*S2C_Items{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (S2C_Items) V2() {
}

func (S2C_Items) MessageID() protocol.MessageID {
	return "msg.S2C_Items"
}

func S2C_Items_MessageID() protocol.MessageID {
	return "msg.S2C_Items"
}

func New_S2C_Items() *S2C_Items {
	m := &S2C_Items{
		Item:  New_Item(),
		Items: []*Item{},
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
		func() protocol.IMessage { return Get_S2C_Items() },
		func(msg protocol.IMessage) { Put_S2C_Items(msg) },
	)
}

// message [S2C_Items] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////