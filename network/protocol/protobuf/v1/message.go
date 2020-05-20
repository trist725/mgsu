package protocol_v1

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	p "github.com/trist725/mgsu/network/protocol"
)

type MessageID = uint16

type IMessage interface {
	proto.Message
	V1()
	MessageID() MessageID
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
type MessageHandler func(IMessage, ...interface{})

type IMessageDispatcher interface {
	Register(id MessageID, handler MessageHandler)
	RegisterWithLimit(id MessageID, handler MessageHandler, minDispatchInterval time.Duration)
	Dispatch(iMsg IMessage, args ...interface{}) error
	SetLimit(id MessageID, minInterval time.Duration) error
}

const DefaultMinDispatchInterval = 100 * time.Millisecond

type OneMessageDispatcher struct {
	Handlers            []MessageHandler
	MinDispatchInterval time.Duration
	LastDispatchTime    time.Time
}

type MessageDispatcher struct {
	Dispatchers map[MessageID]*OneMessageDispatcher
}

func NewMessageDispatcher() IMessageDispatcher {
	return &MessageDispatcher{
		Dispatchers: map[MessageID]*OneMessageDispatcher{},
	}
}

func (md *MessageDispatcher) Register(id MessageID, handler MessageHandler) {
	if dispatcher, ok := md.Dispatchers[id]; !ok {
		md.Dispatchers[id] = &OneMessageDispatcher{
			Handlers:            []MessageHandler{handler},
			MinDispatchInterval: DefaultMinDispatchInterval,
		}
	} else {
		md.Dispatchers[id].Handlers = append(dispatcher.Handlers, handler)
	}
}

func (md *MessageDispatcher) RegisterWithLimit(id MessageID, handler MessageHandler, minDispatchInterval time.Duration) {
	if dispatcher, ok := md.Dispatchers[id]; !ok {
		md.Dispatchers[id] = &OneMessageDispatcher{
			Handlers:            []MessageHandler{handler},
			MinDispatchInterval: minDispatchInterval,
		}
	} else {
		md.Dispatchers[id].Handlers = append(dispatcher.Handlers, handler)
	}
}

func (md *MessageDispatcher) Dispatch(iMsg IMessage, args ...interface{}) error {
	dispatcher, ok := md.Dispatchers[iMsg.MessageID()]
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

	for _, handler := range dispatcher.Handlers {
		handler(iMsg, args...)
	}

	return nil
}

func (md *MessageDispatcher) SetLimit(id MessageID, minInterval time.Duration) error {
	dispatcher, ok := md.Dispatchers[id]
	if !ok {
		return &p.ErrNoDispatcher{
			MessageID: id,
		}
	}
	dispatcher.MinDispatchInterval = minInterval
	return nil
}
