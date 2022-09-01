package service

import (
	"sync"
	"testing"
	"time"

	"github.com/trist725/mgsu/service/rpc"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var s = NewBaseService("testtype", "1", "testname", NewEtcdRegistry(&clientv3.Config{
	Endpoints:   []string{"localhost:2379"},
	DialTimeout: 3 * time.Second,
}, 3*time.Second), &rpc.GreeterServiceImpl{Addr: "[::]:7777"},
	&rpc.GreeterClientImpl{Addr: "localhost:7777", Timeout: 2 * time.Second})

func TestBaseService_Start(t *testing.T) {
	t.Log(s.ID())
	t.Log(s.GetIP())
	t.Log(Conf.GRPCPort)
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
