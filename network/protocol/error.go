package protocol

import "fmt"

////////////////////////////////////////////////////////////////////////////////
type ErrUnsupported struct {
	MessageID interface{}
}

func (e ErrUnsupported) Error() string {
	return fmt.Sprintf("unsupported message, id=[%v]", e.MessageID)
}

////////////////////////////////////////////////////////////////////////////////
type ErrNoDispatcher struct {
	MessageID interface{}
}

func (e ErrNoDispatcher) Error() string {
	return fmt.Sprintf("no message dispatcher, id=[%v]", e.MessageID)
}

////////////////////////////////////////////////////////////////////////////////
type ErrTooLarge struct {
	MessageID interface{}
}

func (e ErrTooLarge) Error() string {
	return fmt.Sprintf("message too large, id=[%v]", e.MessageID)
}

////////////////////////////////////////////////////////////////////////////////
type ErrTooOften struct {
	MessageID interface{}
}

func (e ErrTooOften) Error() string {
	return fmt.Sprintf("message too often, id=[%v]", e.MessageID)
}

////////////////////////////////////////////////////////////////////////////////
type ErrDispatcherAlreadyRegister struct {
	MessageID interface{}
}

func (e ErrDispatcherAlreadyRegister) Error() string {
	return fmt.Sprintf("message dispatcher already register, id=[%v]", e.MessageID)
}
