package service

import (
	"strings"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/trist725/mgsu/log"
	"github.com/trist725/mgsu/util"
)

type IService interface {
	Start()
	Stop()
	Register()
	Sync()
}

type BaseService struct {
	IRegistry
	GreeterServiceImpl

	Index string
	Name  string
	Typ   string
	IP    string

	Cfgs sync.Map // serviceID-key-value
}

func NewBaseService(typ, index, name string, registry IRegistry) *BaseService {
	return &BaseService{
		IRegistry: registry,
		Name:      name,
		Typ:       typ,
		Index:     index,
		IP:        util.GetOutboundIP().String(),
	}
}

func (s *BaseService) Start() {
	s.IRegistry.Init()
	s.Register()
	s.GreeterServiceImpl.Init()
}

func (s *BaseService) Stop() {
	s.IRegistry.Stop()
}

// Register 必须在s.IRegistry.Init()后调用
func (s *BaseService) Register() {
	s.IRegistry.Register(s.WrapPrefix(), map[string]string{"ip": s.IP, "port": Conf.GRPCPort})
	s.Cfgs.Store("ip", s.IP)
	s.Cfgs.Store("port", Conf.GRPCPort)
}

// Sync 必须在s.IRegistry.Init()后调用
func (s *BaseService) Sync() {
	resp, err := s.IRegistry.(*EtcdRegistry).Get(Conf.BasePrefix, clientv3.WithPrefix())
	if err != nil {
		log.Error(err.Error())
		return
	}

	for _, kv := range resp.Kvs {
		subs := strings.Split(string(kv.Key), "/")
		if len(subs) > 0 {
			var (
				subMap sync.Map
				tmp    any
			)
			tmp, _ = s.Cfgs.Load(string(kv.Key)[:len(subs)-1])
			if tmp != nil {
				subMap = tmp.(sync.Map)
			}
			subMap.Store(subs[len(subs)-1], string(kv.Value))
			s.Cfgs.Store(string(kv.Key)[len(Conf.BasePrefix):len(string(kv.Key))-len(subs[len(subs)-1])], subMap)
		}
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
	return Conf.BasePrefix + s.ID()
}
