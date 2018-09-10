package websock_v1

type envelope struct {
	t      int
	msg    []byte
	filter filterFunc
}
