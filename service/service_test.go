package service

import (
	"log"
	"sync"
	"testing"
	"time"

	mlog "github.com/trist725/myleaf/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var s = NewBaseService("testtype", "1", "testname", NewEtcdRegistry(&clientv3.Config{
	Endpoints:   []string{"118.195.177.161:12379"},
	DialTimeout: 3 * time.Second,
}, 3*time.Second, 10), &GreeterServiceImpl{Addr: "[::]:7777"},
	&GreeterClientImpl{BaseClient: NewBaseClient("localhost:7777"), Timeout: 2 * time.Second})

func TestBaseService_Start(t *testing.T) {
	s.Init()
	t.Log(s.ID())
	t.Log(s.GetIP())
	t.Log(s.GetAddr())
	go s.Start()
	s.IRPCClientImpl.Dial()
}

func TestBaseService_Register(t *testing.T) {
	s.IRegistry.Init()
	s.Register()
	resp, err := s.IRegistry.(*EtcdRegistry).Get("/serv", clientv3.WithPrefix())
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Kvs)
}

func TestBaseService_Sync(t *testing.T) {
	s.IRegistry.Init()
	s.Sync()
	s.Cfgs.Range(func(key, value any) bool {
		t.Log("prefix:", key)
		if value != nil {
			subMap := value.(sync.Map)
			subMap.Range(func(key, value any) bool {
				t.Log("key:", key)
				t.Log("value:", value)
				return true
			})
		}
		return true
	})
}

func TestBaseService_Watch(t *testing.T) {
	s.IRegistry.Init()
	s.Sync()
	s.Watch()

	time.Sleep(10 * time.Second)
}

func TestBaseClient(t *testing.T) {
	logger, err := mlog.New("debug", "./log", log.LstdFlags)
	if err != nil {
		panic(err)
	}
	mlog.Export(logger)
	defer logger.Close()
	s.Init()
	s.Dial(grpc.WithTransportCredentials(insecure.NewCredentials()))
	go s.IRPCClientImpl.(*GreeterClientImpl).Do()
	defer s.Close()
	s.Start()
}
func TestBaseService_GetCfgByTyp(t *testing.T) {
	s.Init()
	t.Log(s.ID())
	t.Log(s.GetIP())
	t.Log(s.GetAddr())
	go s.Start()

	res := s.GetCfgByTyp("amax")
	for _, v := range res {
		for k, v2 := range v {
			t.Logf("key:%s, value:%s", k, v2)
		}
	}

}
