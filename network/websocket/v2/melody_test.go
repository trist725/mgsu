package websock_v2

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"testing/quick"
	"time"

	"fmt"

	"github.com/gorilla/websocket"
)

func NewTestServer() *Server {
	m := &Server{
		Melody: NewMelody(DefaultConfig),
	}
	return m
}

func NewTestServerHandler(fn SessionHandlerProduceFunc) *Server {
	m := NewTestServer()
	m.HandleProduceSessionHandler(fn)
	return m
}

func NewDialer(url string) (*websocket.Conn, error) {
	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial(strings.Replace(url, "http", "ws", 1), nil)
	return conn, err
}

type echoSessionHandler struct {
	SessionHandler
	t *testing.T
}

func (echoSessionHandler) HandleTextMessage(s *Session, msg []byte) {
	s.WriteText(msg)
}

func (echoSessionHandler) HandleBinaryMessage(s *Session, msg []byte) {
	s.WriteBinary(msg)
}

func TestEcho(t *testing.T) {
	echo := NewTestServerHandler(func(*http.Request) ISessionHandler {
		return &echoSessionHandler{}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer conn.Close()

		if err != nil {
			t.Error(err)
			return false
		}

		conn.WriteMessage(websocket.TextMessage, []byte(msg))

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if msg != string(ret) {
			t.Errorf("%s should equal %s", msg, string(ret))
			return false
		}

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

type writeClosedSessionHandler struct {
	SessionHandler
	t *testing.T
}

func (writeClosedSessionHandler) HandleConnect(s *Session) error {
	s.WriteText([]byte("hello world when connect"))
	s.Close()
	return nil
}

func (h writeClosedSessionHandler) HandleDisconnect(s *Session) {
	err := s.WriteText([]byte("hello world when disconnect"))
	if err == nil {
		h.t.Error("should be an error")
	}
}

func TestWriteClosed(t *testing.T) {
	echo := NewTestServerHandler(func(*http.Request) ISessionHandler {
		return &writeClosedSessionHandler{
			t: t,
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		if err != nil {
			t.Error(err)
			return false
		}

		_, data, err := conn.ReadMessage()
		if err != nil {
			t.Error(err)
			return false
		}

		fmt.Printf("%s\n", data)

		conn.WriteMessage(websocket.TextMessage, []byte(msg))

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestLen(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	connect := int(rand.Int31n(100))
	disconnect := rand.Float32()
	conns := make([]*websocket.Conn, connect)
	defer func() {
		for _, conn := range conns {
			if conn != nil {
				conn.Close()
			}
		}
	}()

	echo := NewTestServerHandler(func(*http.Request) ISessionHandler {
		return &echoSessionHandler{}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	disconnected := 0
	for i := 0; i < connect; i++ {
		conn, err := NewDialer(server.URL)

		if err != nil {
			t.Error(err)
		}

		if rand.Float32() < disconnect {
			conns[i] = nil
			disconnected++
			conn.Close()
			continue
		}

		conns[i] = conn
	}

	time.Sleep(time.Millisecond)

	connected := connect - disconnected

	if echo.Len() != connected {
		t.Errorf("melody len %d should equal %d", echo.Len(), connected)
	}
}

func TestEchoBinary(t *testing.T) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &echoSessionHandler{}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer conn.Close()

		if err != nil {
			t.Error(err)
			return false
		}

		conn.WriteMessage(websocket.BinaryMessage, []byte(msg))

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if msg != string(ret) {
			t.Errorf("%s should equal %s", msg, string(ret))
			return false
		}

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

type handlersSessionHandler struct {
	echoSessionHandler
	q *Session
}

func (h *handlersSessionHandler) HandleConnect(s *Session) error {
	h.q = s
	s.Close()
	return nil
}

func (h *handlersSessionHandler) HandleDisconnect(s *Session) {
	if h.q != s {
		h.t.Error("disconnecting session should be the same as connecting")
	}
}

func TestHandlers(t *testing.T) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &handlersSessionHandler{
			echoSessionHandler: echoSessionHandler{t: t},
		}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	NewDialer(server.URL)
}

type metaDataSessionHandler struct {
	SessionHandler
}

func (metaDataSessionHandler) HandleConnect(s *Session) error {
	s.Set("stamp", time.Now().UnixNano())
	return nil
}

func (metaDataSessionHandler) HandleTextMessage(s *Session, msg []byte) {
	stamp := s.MustGet("stamp").(int64)
	s.WriteText([]byte(strconv.Itoa(int(stamp))))
}

func TestMetadata(t *testing.T) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &metaDataSessionHandler{}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer conn.Close()

		if err != nil {
			t.Error(err)
			return false
		}

		conn.WriteMessage(websocket.TextMessage, []byte(msg))

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		stamp, err := strconv.Atoi(string(ret))

		if err != nil {
			t.Error(err)
			return false
		}

		diff := int(time.Now().UnixNano()) - stamp

		if diff <= 0 {
			t.Errorf("diff should be above 0 %d", diff)
			return false
		}

		return true
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

type upgraderSessionHandler struct {
	echoSessionHandler
}

func (h upgraderSessionHandler) HandleError(s *Session, err error) {
	if err == nil || err.Error() != "websocket: origin not allowed" {
		h.t.Error("there should be a origin error")
	}
}

func TestUpgrader(t *testing.T) {
	broadcast := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &upgraderSessionHandler{
			echoSessionHandler: echoSessionHandler{t: t},
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	broadcast.Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return false },
	}

	_, err := NewDialer(server.URL)

	if err == nil || err.Error() != "websocket: bad handshake" {
		t.Error("there should be a badhandshake error")
	}
}

type broadcastSessionHandler struct {
	SessionHandler
	broadcast *Server
	t         *testing.T
}

func (h broadcastSessionHandler) HandleTextMessage(s *Session, msg []byte) {
	h.broadcast.BroadcastText(msg)
}

func (h broadcastSessionHandler) HandleBinaryMessage(s *Session, msg []byte) {
	h.broadcast.BroadcastBinary(msg)
}

func TestBroadcastText(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.HandleProduceSessionHandler(func(r *http.Request) ISessionHandler {
		return &broadcastSessionHandler{
			broadcast: broadcast,
			t:         t,
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg string) bool {
		conn, _ := NewDialer(server.URL)
		defer conn.Close()

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer listeners[i].Close()
		}

		conn.WriteMessage(websocket.TextMessage, []byte(msg))

		for i := 0; i < n; i++ {
			_, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if msg != string(ret) {
				t.Errorf("%s should equal %s", msg, string(ret))
				return false
			}
		}

		return true
	}

	if !fn("test") {
		t.Errorf("should not be false")
	}
}

func TestBroadcastBinary(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.HandleProduceSessionHandler(func(r *http.Request) ISessionHandler {
		return &broadcastSessionHandler{
			broadcast: broadcast,
			t:         t,
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg []byte) bool {
		conn, _ := NewDialer(server.URL)
		defer conn.Close()

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer listeners[i].Close()
		}

		conn.WriteMessage(websocket.BinaryMessage, []byte(msg))

		for i := 0; i < n; i++ {
			messageType, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if messageType != websocket.BinaryMessage {
				t.Errorf("message type should be BinaryMessage")
				return false
			}

			if !bytes.Equal(msg, ret) {
				t.Errorf("%v should equal %v", msg, ret)
				return false
			}
		}

		return true
	}

	if !fn([]byte{2, 3, 5, 7, 11}) {
		t.Errorf("should not be false")
	}
}

type broadcastOthersSessionHandler struct {
	SessionHandler
	broadcast *Server
	t         *testing.T
}

func (h broadcastOthersSessionHandler) HandleTextMessage(s *Session, msg []byte) {
	h.broadcast.BroadcastTextOthers(msg, s)
}

func (h broadcastOthersSessionHandler) HandleBinaryMessage(s *Session, msg []byte) {
	h.broadcast.BroadcastBinaryOthers(msg, s)
}

func TestBroadcastTextOthers(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.HandleProduceSessionHandler(func(r *http.Request) ISessionHandler {
		return &broadcastOthersSessionHandler{
			broadcast: broadcast,
			t:         t,
		}
	})
	broadcast.Config.PongWait = time.Second
	broadcast.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg string) bool {
		conn, _ := NewDialer(server.URL)
		defer conn.Close()

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer listeners[i].Close()
		}

		conn.WriteMessage(websocket.TextMessage, []byte(msg))

		for i := 0; i < n; i++ {
			_, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if msg != string(ret) {
				t.Errorf("%s should equal %s", msg, string(ret))
				return false
			}
		}

		return true
	}

	if !fn("test") {
		t.Errorf("should not be false")
	}
}

func TestBroadcastBinaryOthers(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.HandleProduceSessionHandler(func(r *http.Request) ISessionHandler {
		return &broadcastOthersSessionHandler{
			broadcast: broadcast,
			t:         t,
		}
	})
	broadcast.Config.PongWait = time.Second
	broadcast.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(broadcast)
	defer server.Close()

	n := 10

	fn := func(msg []byte) bool {
		conn, _ := NewDialer(server.URL)
		defer conn.Close()

		listeners := make([]*websocket.Conn, n)
		for i := 0; i < n; i++ {
			listener, _ := NewDialer(server.URL)
			listeners[i] = listener
			defer listeners[i].Close()
		}

		conn.WriteMessage(websocket.BinaryMessage, []byte(msg))

		for i := 0; i < n; i++ {
			messageType, ret, err := listeners[i].ReadMessage()

			if err != nil {
				t.Error(err)
				return false
			}

			if messageType != websocket.BinaryMessage {
				t.Errorf("message type should be BinaryMessage")
				return false
			}

			if !bytes.Equal(msg, ret) {
				t.Errorf("%v should equal %v", msg, ret)
				return false
			}
		}

		return true
	}

	if !fn([]byte{2, 3, 5, 7, 11}) {
		t.Errorf("should not be false")
	}
}

func TestPingPong(t *testing.T) {
	noecho := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &SessionHandler{}
	})
	noecho.Config.PongWait = time.Second
	noecho.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(noecho)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	conn.SetPingHandler(func(string) error {
		return nil
	})
	defer conn.Close()

	if err != nil {
		t.Error(err)
	}

	conn.WriteMessage(websocket.TextMessage, []byte("test"))

	_, _, err = conn.ReadMessage()

	if err == nil {
		t.Error("there should be an error")
	}
}

type broadcastFilterSessionHandler struct {
	SessionHandler
	broadcast *Server
}

func (h broadcastFilterSessionHandler) HandleTextMessage(s *Session, msg []byte) {
	h.broadcast.BroadcastTextFilter(msg, func(q *Session) bool {
		return s == q
	})
}

func (h broadcastFilterSessionHandler) HandleBinaryMessage(s *Session, msg []byte) {
	h.broadcast.BroadcastBinaryFilter(msg, func(q *Session) bool {
		return s == q
	})
}

func TestBroadcastTextFilter(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.HandleProduceSessionHandler(func(r *http.Request) ISessionHandler {
		return &broadcastFilterSessionHandler{
			broadcast: broadcast,
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	fn := func(msg string) bool {
		conn, err := NewDialer(server.URL)
		defer conn.Close()

		if err != nil {
			t.Error(err)
			return false
		}

		conn.WriteMessage(websocket.TextMessage, []byte(msg))

		_, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if msg != string(ret) {
			t.Errorf("%s should equal %s", msg, string(ret))
			return false
		}

		return true
	}

	if !fn("test") {
		t.Errorf("should not be false")
	}
}

func TestBroadcastBinaryFilter(t *testing.T) {
	broadcast := NewTestServer()
	broadcast.HandleProduceSessionHandler(func(r *http.Request) ISessionHandler {
		return &broadcastFilterSessionHandler{
			broadcast: broadcast,
		}
	})
	server := httptest.NewServer(broadcast)
	defer server.Close()

	fn := func(msg []byte) bool {
		conn, err := NewDialer(server.URL)
		defer conn.Close()

		if err != nil {
			t.Error(err)
			return false
		}

		conn.WriteMessage(websocket.BinaryMessage, []byte(msg))

		messageType, ret, err := conn.ReadMessage()

		if err != nil {
			t.Error(err)
			return false
		}

		if messageType != websocket.BinaryMessage {
			t.Errorf("message type should be BinaryMessage")
			return false
		}

		if !bytes.Equal(msg, ret) {
			t.Errorf("%v should equal %v", msg, ret)
			return false
		}

		return true
	}

	if !fn([]byte{2, 3, 5, 7, 11}) {
		t.Errorf("should not be false")
	}
}

func TestStop(t *testing.T) {
	noecho := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &SessionHandler{}
	})
	server := httptest.NewServer(noecho)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	defer conn.Close()

	if err != nil {
		t.Error(err)
	}

	noecho.Close()
}

type smallMessageSessionHandler struct {
	echoSessionHandler
}

func (h smallMessageSessionHandler) HandleError(s *Session, err error) {
	if err == nil {
		h.t.Error("there should be a buffer full error here")
	}
}

func TestSmallMessageBuffer(t *testing.T) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &smallMessageSessionHandler{
			echoSessionHandler: echoSessionHandler{t: t},
		}
	})
	echo.Config.MessageBufferSize = 0
	server := httptest.NewServer(echo)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	defer conn.Close()

	if err != nil {
		t.Error(err)
	}

	conn.WriteMessage(websocket.TextMessage, []byte("12345"))
}

type pongSessionHandler struct {
	echoSessionHandler
	fired bool
}

var fired bool

func (h *pongSessionHandler) HandlePong(s *Session) {
	fired = true
}

func TestPong(t *testing.T) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &pongSessionHandler{
			echoSessionHandler: echoSessionHandler{t: t},
		}
	})
	echo.Config.PongWait = time.Second
	echo.Config.PingPeriod = time.Second * 9 / 10
	server := httptest.NewServer(echo)
	defer server.Close()

	conn, err := NewDialer(server.URL)
	defer conn.Close()

	if err != nil {
		t.Error(err)
	}

	conn.WriteMessage(websocket.PongMessage, nil)

	time.Sleep(time.Millisecond)

	if !fired {
		t.Error("should have fired pong handler")
	}
}

func BenchmarkSessionWrite(b *testing.B) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &echoSessionHandler{}
	})
	server := httptest.NewServer(echo)
	conn, _ := NewDialer(server.URL)
	defer server.Close()
	defer conn.Close()

	for n := 0; n < b.N; n++ {
		conn.WriteMessage(websocket.TextMessage, []byte("test"))
		conn.ReadMessage()
	}
}

func BenchmarkBroadcast(b *testing.B) {
	echo := NewTestServerHandler(func(r *http.Request) ISessionHandler {
		return &echoSessionHandler{}
	})
	server := httptest.NewServer(echo)
	defer server.Close()

	conns := make([]*websocket.Conn, 0)

	num := 100

	for i := 0; i < num; i++ {
		conn, _ := NewDialer(server.URL)
		conns = append(conns, conn)
	}

	for n := 0; n < b.N; n++ {
		echo.BroadcastText([]byte("test"))

		for i := 0; i < num; i++ {
			conns[i].ReadMessage()
		}
	}

	for i := 0; i < num; i++ {
		conns[i].Close()
	}
}
