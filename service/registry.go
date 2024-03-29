package service

import (
	context "context"
	"time"

	"github.com/trist725/myleaf/log"

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
	Register(string, interface{})
}

type EtcdRegistry struct {
	ReqTimeout time.Duration
	LeaseTTL   int64 // time-to-live in seconds
	Typ        string
	Endpoints  []string

	Cfg *clientv3.Config
	Cli *clientv3.Client
}

func NewEtcdRegistry(cfg *clientv3.Config, timeout time.Duration, ttl int64) *EtcdRegistry {
	return &EtcdRegistry{
		Cfg:        cfg,
		ReqTimeout: timeout,
		LeaseTTL:   ttl,
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

func (e *EtcdRegistry) Register(prefix string, kvs interface{}) {
	resp, err := e.Cli.Grant(context.TODO(), e.LeaseTTL)
	if err != nil {
		panic(err)
	}

	for k, v := range kvs.(map[string]string) {
		if _, err := e.Put(prefix+k, v, clientv3.WithLease(resp.ID)); err != nil {
			panic(err)
		}
	}
	kaCh, err := e.Cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			ka, ok := <-kaCh
			if !ok {
				log.Debug("etcd KeepAlive failed, chan closed.")
				break
			}
			log.Debug("LeaseID:%d ttl:%d keepalive.", ka.ID, ka.TTL)
		}
	}()
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

func (e *EtcdRegistry) WatchCallBack(watchChan clientv3.WatchChan, fn etcd_v3.WatchCallback) {
	go etcd_v3.WatchLoop(watchChan, fn)
}
