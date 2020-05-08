// Code generated by protoc-gen-pbex-go. DO NOT EDIT IT!!!
// source: login.proto

package msg

import (
	json "encoding/json"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_Register_E_Error] begin

var S2C_Register_E_Error_Slice = []int32{
	0,
	1,
	2,
	3,
	4,
}

func S2C_Register_E_Error_Len() int {
	return len(S2C_Register_E_Error_Slice)
}

func Check_S2C_Register_E_Error_I(value int32) bool {
	if _, ok := S2C_Register_E_Error_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_Register_E_Error(value S2C_Register_E_Error) bool {
	return Check_S2C_Register_E_Error_I(int32(value))
}

func Each_S2C_Register_E_Error(f func(S2C_Register_E_Error) bool) {
	for _, value := range S2C_Register_E_Error_Slice {
		if !f(S2C_Register_E_Error(value)) {
			break
		}
	}
}

func Each_S2C_Register_E_Error_I(f func(int32) bool) {
	for _, value := range S2C_Register_E_Error_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_Register_E_Error] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_Register_E_State] begin

var S2C_Register_E_State_Slice = []int32{
	0,
	1,
	2,
	3,
}

func S2C_Register_E_State_Len() int {
	return len(S2C_Register_E_State_Slice)
}

func Check_S2C_Register_E_State_I(value int32) bool {
	if _, ok := S2C_Register_E_State_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_Register_E_State(value S2C_Register_E_State) bool {
	return Check_S2C_Register_E_State_I(int32(value))
}

func Each_S2C_Register_E_State(f func(S2C_Register_E_State) bool) {
	for _, value := range S2C_Register_E_State_Slice {
		if !f(S2C_Register_E_State(value)) {
			break
		}
	}
}

func Each_S2C_Register_E_State_I(f func(int32) bool) {
	for _, value := range S2C_Register_E_State_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_Register_E_State] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_Login_E_Error] begin

var S2C_Login_E_Error_Slice = []int32{
	0,
	1,
	2,
	3,
	4,
	5,
}

func S2C_Login_E_Error_Len() int {
	return len(S2C_Login_E_Error_Slice)
}

func Check_S2C_Login_E_Error_I(value int32) bool {
	if _, ok := S2C_Login_E_Error_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_Login_E_Error(value S2C_Login_E_Error) bool {
	return Check_S2C_Login_E_Error_I(int32(value))
}

func Each_S2C_Login_E_Error(f func(S2C_Login_E_Error) bool) {
	for _, value := range S2C_Login_E_Error_Slice {
		if !f(S2C_Login_E_Error(value)) {
			break
		}
	}
}

func Each_S2C_Login_E_Error_I(f func(int32) bool) {
	for _, value := range S2C_Login_E_Error_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_Login_E_Error] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Close] begin
func (m *S2C_Close) ResetEx() {

	m.Err = 0

}

func (m S2C_Close) Clone() *S2C_Close {
	n, ok := g_S2C_Close_Pool.Get().(*S2C_Close)
	if !ok || n == nil {
		n = &S2C_Close{}
	}

	n.Err = m.Err

	return n
}

func Clone_S2C_Close_Slice(dst []*S2C_Close, src []*S2C_Close) []*S2C_Close {
	for _, i := range dst {
		Put_S2C_Close(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_Close, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_Close{}
		dst = nil
	}
	return dst
}

func (m S2C_Close) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_Close:" + string(ba)
}

func New_S2C_Close() *S2C_Close {
	m := &S2C_Close{}
	return m
}

var g_S2C_Close_Pool = sync.Pool{}

func Get_S2C_Close() *S2C_Close {
	m, ok := g_S2C_Close_Pool.Get().(*S2C_Close)
	if !ok {
		m = New_S2C_Close()
	} else {
		if m == nil {
			m = New_S2C_Close()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Close(i interface{}) {
	if m, ok := i.(*S2C_Close); ok && m != nil {
		g_S2C_Close_Pool.Put(i)
	}
}

// message [S2C_Close] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_Register] begin
func (m *C2S_Register) ResetEx() {

	m.AccountName = ""

	m.Passwrod = ""

}

func (m C2S_Register) Clone() *C2S_Register {
	n, ok := g_C2S_Register_Pool.Get().(*C2S_Register)
	if !ok || n == nil {
		n = &C2S_Register{}
	}

	n.AccountName = m.AccountName

	n.Passwrod = m.Passwrod

	return n
}

func Clone_C2S_Register_Slice(dst []*C2S_Register, src []*C2S_Register) []*C2S_Register {
	for _, i := range dst {
		Put_C2S_Register(i)
	}
	if len(src) > 0 {
		dst = make([]*C2S_Register, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*C2S_Register{}
		dst = nil
	}
	return dst
}

func (m C2S_Register) JsonString() string {
	ba, _ := json.Marshal(m)
	return "C2S_Register:" + string(ba)
}

func New_C2S_Register() *C2S_Register {
	m := &C2S_Register{}
	return m
}

var g_C2S_Register_Pool = sync.Pool{}

func Get_C2S_Register() *C2S_Register {
	m, ok := g_C2S_Register_Pool.Get().(*C2S_Register)
	if !ok {
		m = New_C2S_Register()
	} else {
		if m == nil {
			m = New_C2S_Register()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_Register(i interface{}) {
	if m, ok := i.(*C2S_Register); ok && m != nil {
		g_C2S_Register_Pool.Put(i)
	}
}

// message [C2S_Register] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Register] begin
func (m *S2C_Register) ResetEx() {

	m.Err = 0

	m.AccountName = ""

	m.Password = ""

	m.State = 0

}

func (m S2C_Register) Clone() *S2C_Register {
	n, ok := g_S2C_Register_Pool.Get().(*S2C_Register)
	if !ok || n == nil {
		n = &S2C_Register{}
	}

	n.Err = m.Err

	n.AccountName = m.AccountName

	n.Password = m.Password

	n.State = m.State

	return n
}

func Clone_S2C_Register_Slice(dst []*S2C_Register, src []*S2C_Register) []*S2C_Register {
	for _, i := range dst {
		Put_S2C_Register(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_Register, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_Register{}
		dst = nil
	}
	return dst
}

func (m S2C_Register) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_Register:" + string(ba)
}

func New_S2C_Register() *S2C_Register {
	m := &S2C_Register{}
	return m
}

var g_S2C_Register_Pool = sync.Pool{}

func Get_S2C_Register() *S2C_Register {
	m, ok := g_S2C_Register_Pool.Get().(*S2C_Register)
	if !ok {
		m = New_S2C_Register()
	} else {
		if m == nil {
			m = New_S2C_Register()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Register(i interface{}) {
	if m, ok := i.(*S2C_Register); ok && m != nil {
		g_S2C_Register_Pool.Put(i)
	}
}

// message [S2C_Register] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_Login] begin
func (m *C2S_Login) ResetEx() {

	m.Account = ""

	m.Password = ""

}

func (m C2S_Login) Clone() *C2S_Login {
	n, ok := g_C2S_Login_Pool.Get().(*C2S_Login)
	if !ok || n == nil {
		n = &C2S_Login{}
	}

	n.Account = m.Account

	n.Password = m.Password

	return n
}

func Clone_C2S_Login_Slice(dst []*C2S_Login, src []*C2S_Login) []*C2S_Login {
	for _, i := range dst {
		Put_C2S_Login(i)
	}
	if len(src) > 0 {
		dst = make([]*C2S_Login, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*C2S_Login{}
		dst = nil
	}
	return dst
}

func (m C2S_Login) JsonString() string {
	ba, _ := json.Marshal(m)
	return "C2S_Login:" + string(ba)
}

func New_C2S_Login() *C2S_Login {
	m := &C2S_Login{}
	return m
}

var g_C2S_Login_Pool = sync.Pool{}

func Get_C2S_Login() *C2S_Login {
	m, ok := g_C2S_Login_Pool.Get().(*C2S_Login)
	if !ok {
		m = New_C2S_Login()
	} else {
		if m == nil {
			m = New_C2S_Login()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_Login(i interface{}) {
	if m, ok := i.(*C2S_Login); ok && m != nil {
		g_C2S_Login_Pool.Put(i)
	}
}

// message [C2S_Login] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Login] begin
func (m *S2C_Login) ResetEx() {

	m.Err = 0

	m.Key = ""

	for _, i := range m.Server {
		Put_S2C_Login_ServerInfo(i)
	}

	//m.Server = []*S2C_Login_ServerInfo{}
	m.Server = nil

	m.LastLoginServerID = 0

}

func (m S2C_Login) Clone() *S2C_Login {
	n, ok := g_S2C_Login_Pool.Get().(*S2C_Login)
	if !ok || n == nil {
		n = &S2C_Login{}
	}

	n.Err = m.Err

	n.Key = m.Key

	if len(m.Server) > 0 {
		n.Server = make([]*S2C_Login_ServerInfo, len(m.Server))
		for i, e := range m.Server {

			if e != nil {
				n.Server[i] = e.Clone()
			}

		}
	} else {
		//n.Server = []*S2C_Login_ServerInfo{}
		n.Server = nil
	}

	n.LastLoginServerID = m.LastLoginServerID

	return n
}

func Clone_S2C_Login_Slice(dst []*S2C_Login, src []*S2C_Login) []*S2C_Login {
	for _, i := range dst {
		Put_S2C_Login(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_Login, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_Login{}
		dst = nil
	}
	return dst
}

func (m S2C_Login) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_Login:" + string(ba)
}

func New_S2C_Login() *S2C_Login {
	m := &S2C_Login{}
	return m
}

var g_S2C_Login_Pool = sync.Pool{}

func Get_S2C_Login() *S2C_Login {
	m, ok := g_S2C_Login_Pool.Get().(*S2C_Login)
	if !ok {
		m = New_S2C_Login()
	} else {
		if m == nil {
			m = New_S2C_Login()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Login(i interface{}) {
	if m, ok := i.(*S2C_Login); ok && m != nil {
		g_S2C_Login_Pool.Put(i)
	}
}

// message [S2C_Login] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Login_ServerInfo] begin
func (m *S2C_Login_ServerInfo) ResetEx() {

	m.ID = 0

	m.Addr = ""

	m.Name = ""

}

func (m S2C_Login_ServerInfo) Clone() *S2C_Login_ServerInfo {
	n, ok := g_S2C_Login_ServerInfo_Pool.Get().(*S2C_Login_ServerInfo)
	if !ok || n == nil {
		n = &S2C_Login_ServerInfo{}
	}

	n.ID = m.ID

	n.Addr = m.Addr

	n.Name = m.Name

	return n
}

func Clone_S2C_Login_ServerInfo_Slice(dst []*S2C_Login_ServerInfo, src []*S2C_Login_ServerInfo) []*S2C_Login_ServerInfo {
	for _, i := range dst {
		Put_S2C_Login_ServerInfo(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_Login_ServerInfo, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_Login_ServerInfo{}
		dst = nil
	}
	return dst
}

func (m S2C_Login_ServerInfo) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_Login_ServerInfo:" + string(ba)
}

func New_S2C_Login_ServerInfo() *S2C_Login_ServerInfo {
	m := &S2C_Login_ServerInfo{}
	return m
}

var g_S2C_Login_ServerInfo_Pool = sync.Pool{}

func Get_S2C_Login_ServerInfo() *S2C_Login_ServerInfo {
	m, ok := g_S2C_Login_ServerInfo_Pool.Get().(*S2C_Login_ServerInfo)
	if !ok {
		m = New_S2C_Login_ServerInfo()
	} else {
		if m == nil {
			m = New_S2C_Login_ServerInfo()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Login_ServerInfo(i interface{}) {
	if m, ok := i.(*S2C_Login_ServerInfo); ok && m != nil {
		g_S2C_Login_ServerInfo_Pool.Put(i)
	}
}

// message [S2C_Login_ServerInfo] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
