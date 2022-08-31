package service

import (
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
	resp, err := s.IRegistry.(*EtcdRegistry).Get(s.WrapPrefix(), clientv3.WithPrefix())
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Kvs)
}
