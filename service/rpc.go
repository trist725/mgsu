package service

type IRPCServerImpl interface {
	Serve()
	GetAddr() string
}

type IRPCClientImpl interface {
	Dial()
}
