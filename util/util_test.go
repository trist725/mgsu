package util

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func testMustMkdirIfNotExist(t *testing.T) {
	MustMkdirIfNotExist("./log/robot")
}

func TestGenRandomByteArray(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(GenRandomByteArray(16))
	}
}

func BenchmarkGenRandomByteArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomByteArray(RandomInt(8, 32))
	}
}

func TestGenRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(GenRandomString(64))
	}
}

func BenchmarkGenRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomString(RandomInt(8, 32))
	}
}

func TestIsDirOrFileExist(t *testing.T) {
	if err := IsDirOrFileExist("../network/fb"); err != nil {
		t.Error(err)
	}
}

func TestTimeFunc(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Truncate(1 * time.Minute))
}

func TestMD5(t *testing.T) {
	fixKey := "8aL3gmNw9bd77hRRc7sRgWSsPccxQGecybgyHFt7yfOj8LcEVcar4u2M75BebWpb"

	i1 := 1 // RandomInt(1, 10000)
	i2 := 2 // RandomInt(1, 10000)
	// key := GenRandomString(64)
	log.Printf("md5 from MD5Sum=[%x]\n", MD5Sum([]byte(fmt.Sprintf("%d%d%s", i1, i2, fixKey))))
	log.Printf("md5 from MD5Sumf=[%x]\n", MD5Sumf("%d%d%s", i1, i2, fixKey))
}

func TestGetTheDayBeginTime(t *testing.T) {
	now := time.Now()
	fmt.Println(GetTheDayBeginTime(now, 0))
}

func TestGetDurationToNextHalfHour(t *testing.T) {
	fmt.Println(time.Now())
	fmt.Println(GetNextHalfHour())
	fmt.Println(GetDurationToNextHalfHour())
}

func TestParseTimeString(t *testing.T) {
	ti, err := ParseTimeString("2017-07-28 8:40:01")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ti)

	ti, err = ParseTimeStringByFormat("2006-01-02 15:04", "2017-07-28 8:40")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ti)
}

func TestParseTimeStringInLocation(t *testing.T) {
	ti, err := ParseTimeStringInLocation("2017-07-28 8:40:01")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ti)

	ti, err = ParseTimeStringInLocationByFormat("2006-01-02 15:04", "2017-07-28 8:40")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ti)
}

func TestGetFuncName(t *testing.T) {
	t.Log(GetFuncName(1))
	t.Log(GetFuncName(2))
	t.Log(GetFuncName(3))
	t.Log(GetFuncName(4))
	t.Log(GetFuncName(0))
	t.Log(GetFuncName(-1))
	t.Log(GetFuncName(-2))
	t.Log(GetFuncName(-3))
	t.Log(GetFuncName(-100))
}
