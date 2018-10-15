// Code generated by protoc-gen-enum-go. DO NOT EDIT IT!!!
// source: version.proto

/*
It has these top-level messages:
	VersionInfo
*/

package msg

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [Version] begin

/// 协议版本枚举
type Version int32

const (
	Version_ Version = 0
	/// 协议版本
	Num Version = 1
)

var Version_name = map[int32]string{
	0: "Version_",
	1: "Num",
}

var Version_value = map[string]int32{
	"Version_": 0,
	"Num":      1,
}

var Version_Slice = []int32{
	0,
	1,
}

func (x Version) String() string {
	if name, ok := Version_name[int32(x)]; ok {
		return name
	}
	return ""
}

func Version_Len() int {
	return len(Version_Slice)
}

func Check_Version_I(value int32) bool {
	if _, ok := Version_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_Version(value Version) bool {
	return Check_Version_I(int32(value))
}

func Each_Version(f func(Version) bool) {
	for _, value := range Version_Slice {
		if !f(Version(value)) {
			break
		}
	}
}

func Each_Version_I(f func(int32) bool) {
	for _, value := range Version_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [Version] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
