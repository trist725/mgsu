package v3

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	p "github.com/trist725/mgsu/network/protocol"
)

type MessageID = uint16
type MessageSeq = uint32

type IMessage interface {
	proto.Message
	V3()
	MessageID() MessageID
	ResponseMessageID() MessageID
	Size() int
	MarshalTo([]byte) (int, error)
	Unmarshal([]byte) error
	ResetEx()
	JsonString() string
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type messageProducer func() IMessage
type messageRecycler func(IMessage)

type messageFactory struct {
	id       MessageID
	producer messageProducer
	recycler messageRecycler
}

func newMessageFactory(id MessageID, producer messageProducer, recycler messageRecycler) *messageFactory {
	return &messageFactory{
		id:       id,
		producer: producer,
		recycler: recycler,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type IMessageFactoryManager interface {
	Register(iMsg IMessage, producer messageProducer, recycler messageRecycler)
	Produce(id MessageID) (IMessage, error)
	Recycle(iMsg IMessage) error
	ReflectType(id MessageID) reflect.Type
}

type messageFactoryManager struct {
	factories map[MessageID]*messageFactory
	id2Type   map[MessageID]reflect.Type
}

func NewMessageFactoryManager() IMessageFactoryManager {
	m := &messageFactoryManager{
		factories: make(map[MessageID]*messageFactory),
		id2Type:   make(map[MessageID]reflect.Type),
	}
	return m
}

func (m *messageFactoryManager) Register(iMsg IMessage, producer messageProducer, recycler messageRecycler) {
	rt := reflect.TypeOf(iMsg)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	name := rt.Name()

	if producer == nil {
		log.Panicf("register message[%s] factory fail, producer is nil, id=[%v]", name, iMsg.MessageID())
	}

	if recycler == nil {
		log.Panicf("register message[%s] factory fail, recycler is nil, id=[%v]", name, iMsg.MessageID())
	}

	if f, ok := m.factories[iMsg.MessageID()]; ok {
		log.Panicf("duplicate message[%s] factory, id=[%v], factory=[%+v]", name, iMsg.MessageID(), f)
	}

	m.factories[iMsg.MessageID()] = newMessageFactory(iMsg.MessageID(), producer, recycler)

	m.id2Type[iMsg.MessageID()] = rt
}

func (m *messageFactoryManager) Produce(id MessageID) (IMessage, error) {
	if factory := m.factories[id]; factory != nil {
		return factory.producer(), nil
	}
	return nil, fmt.Errorf("unsupported message, id=[%v]", id)
}

func (m *messageFactoryManager) Recycle(iMsg IMessage) error {
	if factory := m.factories[iMsg.MessageID()]; factory != nil {
		factory.recycler(iMsg)
		iMsg = nil
		return nil
	}
	return fmt.Errorf("unsupported message, id=[%v]", iMsg.MessageID())
}

func (m messageFactoryManager) ReflectType(id MessageID) reflect.Type {
	return m.id2Type[id]
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MessageHandler func(iRecv IMessage, args ...interface{})

type IMessageDispatcherManager interface {
	MustRegister(id MessageID, handler MessageHandler)
	Register(id MessageID, handler MessageHandler) (err error)
	MustRegisterWithLimit(id MessageID, handler MessageHandler, minDispatchInterval time.Duration)
	RegisterWithLimit(id MessageID, handler MessageHandler, minDispatchInterval time.Duration) (err error)
	Dispatch(iMsg IMessage, args ...interface{}) error
	SetLimit(id MessageID, minInterval time.Duration) error
}

const DefaultMinDispatchInterval = 100 * time.Millisecond

type messageDispatcher struct {
	Handler             MessageHandler
	MinDispatchInterval time.Duration
	LastDispatchTime    time.Time
}

type messageDispatcherManager struct {
	Dispatchers map[MessageID]*messageDispatcher
}

func NewMessageDispatcherManager() IMessageDispatcherManager {
	return &messageDispatcherManager{
		Dispatchers: map[MessageID]*messageDispatcher{},
	}
}

func (mgr *messageDispatcherManager) MustRegister(id MessageID, handler MessageHandler) {
	err := mgr.Register(id, handler)
	if err != nil {
		panic(err)
	}
}

func (mgr *messageDispatcherManager) Register(id MessageID, handler MessageHandler) (err error) {
	_, ok := mgr.Dispatchers[id]
	if ok {
		return &p.ErrDispatcherAlreadyRegister{MessageID: id}
	}
	mgr.Dispatchers[id] = &messageDispatcher{
		Handler:             handler,
		MinDispatchInterval: DefaultMinDispatchInterval,
	}
	return
}

func (mgr *messageDispatcherManager) MustRegisterWithLimit(id MessageID, handler MessageHandler, minDispatchInterval time.Duration) {
	err := mgr.RegisterWithLimit(id, handler, minDispatchInterval)
	if err != nil {
		panic(err)
	}
}

func (mgr *messageDispatcherManager) RegisterWithLimit(id MessageID, handler MessageHandler, minDispatchInterval time.Duration) (err error) {
	_, ok := mgr.Dispatchers[id]
	if ok {
		return &p.ErrDispatcherAlreadyRegister{MessageID: id}
	}
	mgr.Dispatchers[id] = &messageDispatcher{
		Handler:             handler,
		MinDispatchInterval: minDispatchInterval,
	}
	return
}

func (mgr *messageDispatcherManager) Dispatch(iMsg IMessage, args ...interface{}) error {
	dispatcher, ok := mgr.Dispatchers[iMsg.MessageID()]
	if !ok {
		return &p.ErrNoDispatcher{
			MessageID: iMsg.MessageID(),
		}
	}

	now := time.Now()
	if now.Sub(dispatcher.LastDispatchTime) < dispatcher.MinDispatchInterval {
		return &p.ErrTooOften{
			MessageID: iMsg.MessageID(),
		}
	}

	dispatcher.LastDispatchTime = now

	dispatcher.Handler(iMsg, args...)

	return nil
}

func (mgr *messageDispatcherManager) SetLimit(id MessageID, minInterval time.Duration) error {
	dispatcher, ok := mgr.Dispatchers[id]
	if !ok {
		return &p.ErrNoDispatcher{
			MessageID: id,
		}
	}
	dispatcher.MinDispatchInterval = minInterval
	return nil
}
