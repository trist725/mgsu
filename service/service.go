package service

import (
	"strings"

	"github.com/trist725/mgsu/util"
)

type IService interface {
	Start()
	Stop()
	Register()
}

type BaseService struct {
	IRegistry
	GreeterServiceImpl

	Index string
	Name  string
	Typ   string
	IP    string
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
	s.GreeterServiceImpl.Init()
}

func (s *BaseService) Stop() {
	s.IRegistry.Stop()
}

func (s *BaseService) Register() {
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
	b.WriteString("-")
	b.WriteString(s.Name)
	b.WriteString("-")
	b.WriteString(s.Index)
	return b.String()
}
