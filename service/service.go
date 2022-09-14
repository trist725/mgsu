package service

import (
	"net"
	"strings"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/trist725/mgsu/log"
	"github.com/trist725/mgsu/util"
)

type IService interface {
	Init()
	Start()
	Stop()
	Register()
	Sync()
}

type BaseService struct {
	IRegistry
	IRPCServerImpl
	IRPCClientImpl

	Index      string
	Name       string
	Typ        string
	IP         string
	BasePrefix string

	Cfgs sync.Map // serviceID-key-value
}

func NewBaseService(typ, index, name string, registry IRegistry, server IRPCServerImpl, client IRPCClientImpl) *BaseService {
	ip, _ := util.GetWANIP()
	return &BaseService{
		BasePrefix:     "/service/",
		IRegistry:      registry,
		Name:           name,
		Typ:            typ,
		Index:          index,
		IP:             ip,
		IRPCServerImpl: server,
		IRPCClientImpl: client,
	}
}

func (s *BaseService) Init() {
	if s.IRegistry == nil {
		log.Warn("nil Registry")
		return
	}
	s.IRegistry.Init()
	s.Register()
	s.Sync()
	s.Watch()
}

func (s *BaseService) Start() {
	if s.IRPCServerImpl == nil {
		log.Warn("nil RPC srever")
		return
	}
	s.IRPCServerImpl.Serve()
}

func (s *BaseService) Stop() {
	s.IRegistry.Stop()
}

// Register 必须在s.IRegistry.Init()后调用
func (s *BaseService) Register() {
	if s.IRPCServerImpl == nil {
		log.Warn("Register(): nil RPC srever")
		return
	}
	_, port, err := net.SplitHostPort(s.IRPCServerImpl.GetAddr())
	if err != nil {
		panic(err)
	}
	s.IRegistry.Register(s.WrapPrefix(), map[string]string{"ip": s.IP, "port": port})
	m := sync.Map{}
	m.Store("ip", s.IP)
	m.Store("port", port)
	s.Cfgs.Store(s.ID(), m)
}

// Sync 必须在s.IRegistry.Init()后调用
func (s *BaseService) Sync() {
	switch s.IRegistry.(type) {
	case *EtcdRegistry:
		resp, err := s.IRegistry.(*EtcdRegistry).Get(s.BasePrefix, clientv3.WithPrefix())
		if err != nil {
			log.Error(err.Error())
			return
		}

		for _, v := range resp.Kvs {
			s.etcdSync(v, clientv3.EventTypePut)
		}
	}

}

func (s *BaseService) etcdSync(kv *mvccpb.KeyValue, evt mvccpb.Event_EventType) {
	subs := strings.Split(string(kv.Key), "/")
	if len(subs) > 0 {
		var (
			subMap    sync.Map
			tmp       any
			serviceID = string(kv.Key)[len(s.BasePrefix) : len(string(kv.Key))-len(subs[len(subs)-1])]
			key       = subs[len(subs)-1]
		)

		tmp, _ = s.Cfgs.Load(serviceID)
		if tmp != nil {
			subMap = tmp.(sync.Map)
		}
		switch evt {
		case clientv3.EventTypePut:
			subMap.Store(key, string(kv.Value))
			log.Info("put %s:%s:%s", serviceID, key, string(kv.Value))
		case clientv3.EventTypeDelete:
			subMap.Delete(key)
			log.Info("delete %s:%s:%s", serviceID, key, string(kv.Value))
		}
		s.Cfgs.Store(serviceID, subMap)
	}
}

func (s *BaseService) GetType() string {
	return s.Typ
}

func (s *BaseService) GetName() string {
	return s.Name
}

func (s *BaseService) GetIP() string {
	return s.IP
}

func (s *BaseService) ID() string {
	b := strings.Builder{}
	b.WriteString(s.Typ)
	b.WriteString("/")
	b.WriteString(s.Name)
	b.WriteString(s.Index)
	b.WriteString("/")
	return b.String()
}

func (s *BaseService) WrapPrefix() string {
	return s.BasePrefix + s.ID()
}

func (s *BaseService) Watch() {
	switch s.IRegistry.(type) {
	case *EtcdRegistry:
		e := s.IRegistry.(*EtcdRegistry)
		e.WatchCallBack(e.Watch(s.BasePrefix, clientv3.WithPrefix()), func(ev *clientv3.Event) {
			switch ev.Type {
			case clientv3.EventTypePut:
				s.etcdSync(ev.Kv, clientv3.EventTypePut)
			case clientv3.EventTypeDelete:
				s.etcdSync(ev.Kv, clientv3.EventTypeDelete)
			}
		})
	}
}

func (s *BaseService) SetBasePrefix(prefix string) {
	s.BasePrefix = prefix
}
