package connection

import "github.com/quickfixgo/quickfix"

type FixInitiatorApplication interface {
	FixConnectionApplication
}

type FixInitiator struct {
	initiator   *quickfix.Initiator
	application FixInitiatorApplication
	parameters  *FixConnectionParameters
}

type FixInitiatorFactory interface {
	CreateNewInitiator(parameters *FixConnectionParameters, app FixInitiatorApplication) (*FixInitiator, error)
}

type fixInitiatorFactoryImpl struct{}

func NewFixInitiatorFactory() FixInitiatorFactory {
	return &fixInitiatorFactoryImpl{}
}

func (f *fixInitiatorFactoryImpl) CreateNewInitiator(parameters *FixConnectionParameters, app FixInitiatorApplication) (*FixInitiator, error) {
	// obs: defined previously due to OnCreate method, which is called when initiator is created
	app.SetConnectionParameters(parameters)

	initiator, err := quickfix.NewInitiator(
		app,
		parameters.MessageStoreFactory,
		parameters.AppSettings,
		parameters.LogFactory,
	)

	if err != nil {
		app.SetConnectionParameters(nil)
		return nil, err
	}

	fixInitiator := &FixInitiator{
		initiator:   initiator,
		application: app,
		parameters:  parameters,
	}

	return fixInitiator, nil
}

func (a *FixInitiator) Start() error {
	return a.initiator.Start()
}

func (a *FixInitiator) Stop() error {
	a.initiator.Stop()
	return nil
}

func (a *FixInitiator) SocketAcceptPort() (int, bool) {
	return 0, false
}
