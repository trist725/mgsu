package service

import (
	context "context"
	"time"

	etcd_v3 "github.com/trist725/mgsu/etcd/client/v3"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	Etcd      string = "etcd"
	Consul           = "consul"
	Zookeeper        = "zookeeper"
)

type IRegistry interface {
	Init()
	Stop()
	GetType() string
}

type EtcdRegistry struct {
	ReqTimeout time.Duration

	Typ       string
	Endpoints []string

	Cfg *clientv3.Config
	Cli *clientv3.Client
}

func NewEtcdRegistry(cfg *clientv3.Config) *EtcdRegistry {
	return &EtcdRegistry{
		Cfg: cfg,
	}
}

func (e *EtcdRegistry) Init() {
	e.Typ = Etcd

	if client, err := clientv3.New(*e.Cfg); err != nil {
		panic(err)
		return
	} else {
		e.Cli = client
	}
}

func (e *EtcdRegistry) Stop() {
	if e.Cfg == nil {
		return
	}
	e.Cli.Close()
}

func (e *EtcdRegistry) GetType() string {
	return e.Typ
}

func (e *EtcdRegistry) Put(key, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.ReqTimeout)
	resp, err := e.Cli.Put(ctx, key, value, opts...)
	cancel()
	return resp, err
}

func (e *EtcdRegistry) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.ReqTimeout)
	resp, err := e.Cli.Get(ctx, key, opts...)
	cancel()
	return resp, err
}

func (e *EtcdRegistry) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.ReqTimeout)
	resp, err := e.Cli.Delete(ctx, key, opts...)
	cancel()
	return resp, err
}

func (e *EtcdRegistry) Watch(key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	wch := e.Cli.Watch(context.Background(), key, opts...)
	return wch
}

func (e *EtcdRegistry) WatchCallBack(watchChan clientv3.WatchChan) {
	go etcd_v3.WatchLoop(watchChan, func(ev *clientv3.Event) {

	})
}
