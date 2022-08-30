package service

import (
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestBaseService_Start(t *testing.T) {
	s := NewBaseService("testtype", "1", "testname", NewEtcdRegistry(&clientv3.Config{
		Endpoints:   []string{"118.195.177.161:12379"},
		DialTimeout: 3 * time.Second,
	}))

	t.Log(s.ID())
	t.Log(s.GetIP())
	t.Log(Conf.GRPCPort)
	s.Start()
}
