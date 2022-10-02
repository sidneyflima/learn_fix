package connection

import (
	"sidney/examples/learn_fix/internal/domain/repository"
	"sync"

	"github.com/quickfixgo/quickfix"
)

// obs: inspired on workaround used to allow dynamic sessions on acceptor
// source: https://github.com/quickfixgo/quickfix/issues/484

type dynamicSessionMessageFactory struct {
	store           quickfix.MessageStoreFactory
	settings        *quickfix.Settings
	usersRepository repository.UserSessionRepository
	mu              sync.Mutex
}

func newDynamicSessionMessageFactory(parameters *FixConnectionParameters) quickfix.MessageStoreFactory {
	return &dynamicSessionMessageFactory{
		store:           parameters.MessageStoreFactory,
		settings:        parameters.AppSettings,
		usersRepository: parameters.UsersRepository,
	}
}

func (dfs *dynamicSessionMessageFactory) Create(sessionID quickfix.SessionID) (msgStore quickfix.MessageStore, err error) {
	usersSession, err := dfs.usersRepository.GetAll()
	if err != nil {
		return dfs.store.Create(sessionID)
	}

	dfs.mu.Lock()
	for _, user := range usersSession {
		session := quickfix.NewSessionSettings()
		session.Set("BeginString", sessionID.BeginString)
		session.Set("TargetCompID", user.SenderCompID) // use user SenderCompID as TargetCompID
		session.Set("TargetSubID", sessionID.TargetSubID)
		session.Set("TargetLocationID", sessionID.TargetLocationID)
		session.Set("SenderCompID", sessionID.SenderCompID)
		session.Set("SenderSubID", sessionID.SenderSubID)
		session.Set("SenderLocationID", sessionID.SenderLocationID)
		session.Set("SessionQualifier", sessionID.Qualifier)

		// obs: commented for testing only
		// if _, ok := dfs.settings.SessionSettings()[sessionID]; !ok {
		dfs.settings.AddSession(session)
		//}
	}
	dfs.mu.Unlock()

	return dfs.store.Create(sessionID)
}
