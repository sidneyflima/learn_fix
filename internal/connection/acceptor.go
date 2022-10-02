package connection

import "github.com/quickfixgo/quickfix"

type FixAcceptorApplication interface {
	FixConnectionApplication
}

type FixAcceptor struct {
	acceptor    *quickfix.Acceptor
	application FixAcceptorApplication
	parameters  *FixConnectionParameters
}

type FixAcceptorFactory interface {
	CreateNewAcceptor(parameters *FixConnectionParameters, app FixAcceptorApplication) (*FixAcceptor, error)
}

type fixAcceptorFactoryImpl struct{}

func NewFixAcceptorFactory() FixAcceptorFactory {
	return &fixAcceptorFactoryImpl{}
}

func (f *fixAcceptorFactoryImpl) CreateNewAcceptor(parameters *FixConnectionParameters, app FixAcceptorApplication) (*FixAcceptor, error) {
	// obs: defined previously due to OnCreate method, which is called when acceptor is created
	app.SetConnectionParameters(parameters)

	acceptor, err := quickfix.NewAcceptor(
		app,
		newDynamicSessionMessageFactory(parameters),
		parameters.AppSettings,
		parameters.LogFactory,
	)

	if err != nil {
		app.SetConnectionParameters(nil)
		return nil, err
	}

	acceptor.BuildInitiators = false

	fixAcceptor := &FixAcceptor{
		acceptor:    acceptor,
		application: app,
		parameters:  parameters,
	}

	return fixAcceptor, nil
}

func (a *FixAcceptor) Start() error {
	return a.acceptor.Start()
}

func (a *FixAcceptor) Stop() error {
	a.acceptor.Stop()
	return nil
}

func (a *FixAcceptor) SocketAcceptPort() (int, bool) {
	sessionSettings := a.parameters.AppSettings.GlobalSettings()
	if !sessionSettings.HasSetting("SocketAcceptPort") {
		return 0, false
	}

	if port, err := sessionSettings.IntSetting("SocketAcceptPort"); err == nil {
		return port, true
	}

	return 0, false
}
