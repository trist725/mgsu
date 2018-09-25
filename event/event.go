package event

import (
	"fmt"
)

type ID = uint16

type IEvent interface {
	Event()
	EventID() ID
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Handler func(iEvent IEvent, args ...interface{})

type HandlerManager struct {
	handlers map[ID][]Handler
}

func NewHandlerManager() *HandlerManager {
	return &HandlerManager{
		handlers: make(map[ID][]Handler),
	}
}

func (m *HandlerManager) Register(id ID, handler Handler) {
	if handlers, ok := m.handlers[id]; !ok {
		m.handlers[id] = []Handler{handler}
	} else {
		m.handlers[id] = append(handlers, handler)
	}
}

func (m *HandlerManager) Process(iEvent IEvent, args ...interface{}) error {
	if handlers, ok := m.handlers[iEvent.EventID()]; !ok {
		return fmt.Errorf("no handler to process event, id=[%d], [%#v]", iEvent.EventID(), iEvent)
	} else {
		for _, handler := range handlers {
			handler(iEvent, args...)
		}
	}
	return nil
}
