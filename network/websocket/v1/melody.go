package websock_v1

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

// Close codes defined in RFC 6455, section 11.7.
// Duplicate of codes from gorilla/websocket for convenience.
const (
	CloseNormalClosure           = 1000
	CloseGoingAway               = 1001
	CloseProtocolError           = 1002
	CloseUnsupportedData         = 1003
	CloseNoStatusReceived        = 1005
	CloseAbnormalClosure         = 1006
	CloseInvalidFramePayloadData = 1007
	ClosePolicyViolation         = 1008
	CloseMessageTooBig           = 1009
	CloseMandatoryExtension      = 1010
	CloseInternalServerErr       = 1011
	CloseServiceRestart          = 1012
	CloseTryAgainLater           = 1013
	CloseTLSHandshake            = 1015
)

// Duplicate of codes from gorilla/websocket for convenience.
var validReceivedCloseCodes = map[int]bool{
	// see http://www.iana.org/assignments/websocket/websocket.xhtml#close-code-number

	CloseNormalClosure:           true,
	CloseGoingAway:               true,
	CloseProtocolError:           true,
	CloseUnsupportedData:         true,
	CloseNoStatusReceived:        false,
	CloseAbnormalClosure:         false,
	CloseInvalidFramePayloadData: true,
	ClosePolicyViolation:         true,
	CloseMessageTooBig:           true,
	CloseMandatoryExtension:      true,
	CloseInternalServerErr:       true,
	CloseServiceRestart:          true,
	CloseTryAgainLater:           true,
	CloseTLSHandshake:            false,
}

type handleMessageFunc func(*Session, []byte)
type handleErrorFunc func(*Session, error)
type handleCloseFunc func(*Session, int, string) error
type handleSessionFunc func(*Session)
type filterFunc func(*Session) bool

// Melody implements a websocket manager.
type Melody struct {
	Config                        Config
	Upgrader                      *websocket.Upgrader
	textMessageHandler            handleMessageFunc
	binaryMessageHandler          handleMessageFunc
	afterSendTextMessageHandler   handleMessageFunc
	afterSendBinaryMessageHandler handleMessageFunc
	errorHandler                  handleErrorFunc
	closeHandler                  handleCloseFunc
	connectHandler                handleSessionFunc
	disconnectHandler             handleSessionFunc
	pongHandler                   handleSessionFunc
	sessionMgr                    *sessionManager
}

// NewMelody creates a new melody instance with default Upgrader and Config.
func NewMelody(cfg Config) *Melody {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	m := &Melody{
		Config:                        cfg,
		Upgrader:                      upgrader,
		textMessageHandler:            func(*Session, []byte) {},
		binaryMessageHandler:          func(*Session, []byte) {},
		afterSendTextMessageHandler:   func(*Session, []byte) {},
		afterSendBinaryMessageHandler: func(*Session, []byte) {},
		errorHandler:                  func(*Session, error) {},
		closeHandler:                  nil,
		connectHandler:                func(*Session) {},
		disconnectHandler:             func(*Session) {},
		pongHandler:                   func(*Session) {},
		sessionMgr:                    newSessionManager(),
	}

	return m
}

// HandleConnect fires fn when a session connects.
func (m *Melody) HandleConnect(fn func(*Session)) {
	m.connectHandler = fn
}

// HandleDisconnect fires fn when a session disconnects.
func (m *Melody) HandleDisconnect(fn func(*Session)) {
	m.disconnectHandler = fn
}

// HandlePong fires fn when a pong is received from a session.
func (m *Melody) HandlePong(fn func(*Session)) {
	m.pongHandler = fn
}

// HandleTextMessage fires fn when a text message comes in.
func (m *Melody) HandleTextMessage(fn func(*Session, []byte)) {
	m.textMessageHandler = fn
}

// HandleBinaryMessage fires fn when a binary message comes in.
func (m *Melody) HandleBinaryMessage(fn func(*Session, []byte)) {
	m.binaryMessageHandler = fn
}

// HandleAfterSendTextMessage fires fn when a text message is successfully sent.
func (m *Melody) HandleAfterSendTextMessage(fn func(*Session, []byte)) {
	m.afterSendTextMessageHandler = fn
}

// HandleAfterSendBinaryMessage fires fn when a binary message is successfully sent.
func (m *Melody) HandleAfterSendBinaryMessage(fn func(*Session, []byte)) {
	m.afterSendBinaryMessageHandler = fn
}

// HandleError fires fn when a session has an error.
func (m *Melody) HandleError(fn func(*Session, error)) {
	m.errorHandler = fn
}

// HandleClose sets the handler for close messages received from the session.
// The code argument to h is the received close code or CloseNoStatusReceived
// if the close message is empty. The default close handler sends a close frame
// back to the session.
//
// The application must read the connection to process close messages as
// described in the section on Control Frames above.
//
// The connection read methods return a CloseError when a close frame is
// received. Most applications should handle close messages as part of their
// normal error handling. Applications should only set a close handler when the
// application must perform some action before sending a close frame back to
// the session.
func (m *Melody) HandleClose(fn func(*Session, int, string) error) {
	if fn != nil {
		m.closeHandler = fn
	}
}

// HandleRequest upgrades http requests to websocket connections and dispatches them to be handled by the melody instance.
func (m *Melody) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return m.HandleRequestWithStates(w, r, nil)
}

// HandleRequestWithStates does the same as HandleRequest but populates session.States with states.
func (m *Melody) HandleRequestWithStates(w http.ResponseWriter, r *http.Request, states SessionStates) error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is closed")
	}

	conn, err := m.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	session := NewSession(r, conn, m, states)

	m.sessionMgr.register(session)

	m.connectHandler(session)

	go session.writePump()

	session.readPump()

	if !m.sessionMgr.closed() {
		m.sessionMgr.unregister(session)
	}

	session.close()

	m.disconnectHandler(session)

	return nil
}

// BroadcastText broadcasts a text message to all sessions.
func (m *Melody) BroadcastText(msg []byte) error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is closed")
	}

	message := &envelope{t: websocket.TextMessage, msg: msg}
	m.sessionMgr.broadcast(message)

	return nil
}

// BroadcastTextFilter broadcasts a text message to all sessions that fn returns true for.
func (m *Melody) BroadcastTextFilter(msg []byte, fn func(*Session) bool) error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is closed")
	}

	message := &envelope{t: websocket.TextMessage, msg: msg, filter: fn}
	m.sessionMgr.broadcast(message)

	return nil
}

// BroadcastTextOthers broadcasts a text message to all sessions except session s.
func (m *Melody) BroadcastTextOthers(msg []byte, s *Session) error {
	return m.BroadcastTextFilter(msg, func(q *Session) bool {
		return s != q
	})
}

// BroadcastTextMultiple broadcasts a text message to multiple sessions given in the sessions slice.
func (m *Melody) BroadcastTextMultiple(msg []byte, sessions []*Session) error {
	for _, s := range sessions {
		if writeErr := s.WriteText(msg); writeErr != nil {
			return writeErr
		}
	}
	return nil
}

// BroadcastBinary broadcasts a binary message to all sessions.
func (m *Melody) BroadcastBinary(msg []byte) error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is closed")
	}

	message := &envelope{t: websocket.BinaryMessage, msg: msg}
	m.sessionMgr.broadcast(message)

	return nil
}

// BroadcastBinaryFilter broadcasts a binary message to all sessions that fn returns true for.
func (m *Melody) BroadcastBinaryFilter(msg []byte, fn func(*Session) bool) error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is closed")
	}

	message := &envelope{t: websocket.BinaryMessage, msg: msg, filter: fn}
	m.sessionMgr.broadcast(message)

	return nil
}

// BroadcastBinaryOthers broadcasts a binary message to all sessions except session s.
func (m *Melody) BroadcastBinaryOthers(msg []byte, s *Session) error {
	return m.BroadcastBinaryFilter(msg, func(q *Session) bool {
		return s != q
	})
}

// BroadcastBinaryMultiple broadcasts a binary message to multiple sessions given in the sessions slice.
func (m *Melody) BroadcastBinaryMultiple(msg []byte, sessions []*Session) error {
	for _, s := range sessions {
		if writeErr := s.WriteBinary(msg); writeErr != nil {
			return writeErr
		}
	}
	return nil
}

// Close closes the melody instance and all connected sessions.
func (m *Melody) Close() error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is already closed")
	}

	m.sessionMgr.close(&envelope{t: websocket.CloseMessage, msg: []byte{}})

	return nil
}

// CloseWithMsg closes the melody instance with the given close payload and all connected sessions.
// Use the FormatCloseMessage function to format a proper close message payload.
func (m *Melody) CloseWithMsg(msg []byte) error {
	if m.sessionMgr.closed() {
		return errors.New("melody instance is already closed")
	}

	m.sessionMgr.close(&envelope{t: websocket.CloseMessage, msg: msg})

	return nil
}

// Len return the number of connected sessions.
func (m *Melody) Len() int {
	return m.sessionMgr.len()
}

// IsClosed returns the status of the melody instance.
func (m *Melody) IsClosed() bool {
	return m.sessionMgr.closed()
}

// FormatCloseMessage formats closeCode and text as a WebSocket close message.
func FormatCloseMessage(closeCode int, text string) []byte {
	return websocket.FormatCloseMessage(closeCode, text)
}
