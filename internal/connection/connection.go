package connection

import "github.com/quickfixgo/quickfix"

type FixConnection interface {
	Start() error
	Stop() error
	SocketAcceptPort() (int, bool)
}

type FixConnectionApplication interface {
	quickfix.Application
	SetConnectionParameters(parameters *FixConnectionParameters)
}
