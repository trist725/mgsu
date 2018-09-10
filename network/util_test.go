package network

import (
	"net/http"
	"testing"

	"time"

	"fmt"
	"log"
	"net"

	"encoding/json"

	"math/rand"

	"gitee.com/nggs/util"
)

type param struct {
	I32  int32
	I64  int64
	F32  float32
	F64  float64
	Str  string
	Time time.Time
}

var sleepTime = 1 * time.Second
var waitTime = 3 * time.Second

func TestHttpGetJson(t *testing.T) {
	addr := "localhost:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		t.Errorf("listen [%s] fail, %s", addr, err)
		return
	}

	pattern := "/test"

	mux := http.NewServeMux()
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("token=[%s]\n", r.Header.Get("Authorization"))

		recv := param{
			I32:  util.RandomInt32(0, 10000),
			I64:  util.RandomInt64(0, 10000),
			F32:  rand.Float32(),
			F64:  rand.Float64(),
			Str:  util.GenRandomString(32),
			Time: time.Now().Add(util.RandomTimeDuration(5*time.Second, 10000*time.Second)),
		}
		//log.Printf("before decode, recv=%v\n", recv)
		json.NewDecoder(r.Body).Decode(&recv)
		log.Printf("after decode, recv=%v\n", recv)

		time.Sleep(sleepTime)

		send := param{
			I32:  util.RandomInt32(0, 10000),
			I64:  util.RandomInt64(0, 10000),
			F32:  rand.Float32(),
			F64:  rand.Float64(),
			Str:  util.GenRandomString(32),
			Time: time.Now().Add(util.RandomTimeDuration(1*time.Second, 10000*time.Second)),
		}
		log.Printf("send=%v\n", send)
		json.NewEncoder(w).Encode(send)
	})

	svr := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go svr.Serve(ln)

	send := param{
		I32:  util.RandomInt32(0, 10000),
		I64:  util.RandomInt64(0, 10000),
		F32:  rand.Float32(),
		F64:  rand.Float64(),
		Str:  util.GenRandomString(32),
		Time: time.Now().Add(util.RandomTimeDuration(1*time.Second, 10000*time.Second)),
	}

	log.Printf("send=%v\n", send)

	recv := param{}

	if err := HttpGetJson(fmt.Sprintf("%s%s", addr, pattern), waitTime, "123", send, &recv); err != nil {
		t.Error(err)
		return
	}

	log.Printf("recv=%v\n", recv)

	time.Sleep(sleepTime)

	svr.Close()
}
