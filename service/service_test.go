package service

import (
	"sync"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var s = NewBaseService("testtype", "1", "testname", NewEtcdRegistry(&clientv3.Config{
	Endpoints:   []string{"118.195.177.161:12379"},
	DialTimeout: 3 * time.Second,
}, 3*time.Second))

func TestBaseService_Start(t *testing.T) {
	t.Log(s.ID())
	t.Log(s.GetIP())
	t.Log(Conf.GRPCPort)
	s.Start()
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
