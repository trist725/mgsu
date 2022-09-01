package service

type IRPCServerImpl interface {
	Serve()
}

type IRPCClientImpl interface {
	Dial()
}
