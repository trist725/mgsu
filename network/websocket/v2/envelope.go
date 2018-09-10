package websock_v2

type envelope struct {
	t      int
	msg    []byte
	filter filterFunc
}
