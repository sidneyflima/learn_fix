package application

import (
	"log"
	"sidney/examples/learn_fix/internal/connection"

	"github.com/quickfixgo/quickfix"
)

type fixInitiatorApplicationImpl struct {
	parameters *connection.FixConnectionParameters
}

func NewFixInitiatorApplication() connection.FixInitiatorApplication {
	return &fixInitiatorApplicationImpl{
		parameters: nil,
	}
}

func (a *fixInitiatorApplicationImpl) SetConnectionParameters(parameters *connection.FixConnectionParameters) {
	a.parameters = parameters
}

//Notification of a session begin created.
func (a *fixInitiatorApplicationImpl) OnCreate(sessionID quickfix.SessionID) {
	log.Println("Initiator - OnCreate", sessionID.TargetCompID)
}

//Notification of a session successfully logging on.
func (a *fixInitiatorApplicationImpl) OnLogon(sessionID quickfix.SessionID) {
	log.Println("Initiator - OnLogon", sessionID.TargetCompID)
}

//Notification of a session logging off or disconnecting.
func (a *fixInitiatorApplicationImpl) OnLogout(sessionID quickfix.SessionID) {
	log.Println("Initiator - OnLogout", sessionID.TargetCompID)
}

//Notification of admin message being sent to target.
func (a *fixInitiatorApplicationImpl) ToAdmin(message *quickfix.Message, sessionID quickfix.SessionID) {
	msgtype, _ := message.MsgType()
	log.Println("Initiator - ToAdmin", sessionID.TargetCompID, "MsgType", msgtype)
}

//Notification of app message being sent to target.
func (a *fixInitiatorApplicationImpl) ToApp(message *quickfix.Message, sessionID quickfix.SessionID) error {
	msgtype, _ := message.MsgType()
	log.Println("Initiator - ToApp", sessionID.TargetCompID, "MsgType", msgtype)
	return nil
}

//Notification of admin message being received from target.
func (a *fixInitiatorApplicationImpl) FromAdmin(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	msgtype, _ := message.MsgType()
	log.Println("Initiator - FromAdmin", sessionID.TargetCompID, "MsgType", msgtype)
	return nil
}

//Notification of app message being received from target.
func (a *fixInitiatorApplicationImpl) FromApp(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	msgtype, _ := message.MsgType()
	log.Println("Initiator - FromApp", sessionID.TargetCompID, "MsgType", msgtype)
	return nil
}
