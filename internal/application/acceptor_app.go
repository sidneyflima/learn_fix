package application

import (
	"log"
	"sidney/examples/learn_fix/internal/connection"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
)

type fixAcceptorApplicationImpl struct {
	parameters *connection.FixConnectionParameters
}

func NewFixAcceptorApplication() connection.FixAcceptorApplication {
	return &fixAcceptorApplicationImpl{
		parameters: nil,
	}
}

func (a *fixAcceptorApplicationImpl) SetConnectionParameters(parameters *connection.FixConnectionParameters) {
	a.parameters = parameters
}

//Notification of a session begin created.
func (a *fixAcceptorApplicationImpl) OnCreate(sessionID quickfix.SessionID) {
	log.Println("Acceptor - OnCreate", sessionID.TargetCompID)
}

//Notification of a session successfully logging on.
func (a *fixAcceptorApplicationImpl) OnLogon(sessionID quickfix.SessionID) {
	log.Println("Acceptor - OnLogon", sessionID.TargetCompID)
}

//Notification of a session logging off or disconnecting.
func (a *fixAcceptorApplicationImpl) OnLogout(sessionID quickfix.SessionID) {
	log.Println("Acceptor - OnLogout", sessionID.TargetCompID)
}

//Notification of admin message being sent to target.
func (a *fixAcceptorApplicationImpl) ToAdmin(message *quickfix.Message, sessionID quickfix.SessionID) {
	msgtype, _ := message.MsgType()
	log.Println("Acceptor - ToAdmin", sessionID.TargetCompID, "MsgType", msgtype)
}

//Notification of app message being sent to target.
func (a *fixAcceptorApplicationImpl) ToApp(message *quickfix.Message, sessionID quickfix.SessionID) error {
	msgtype, _ := message.MsgType()
	log.Println("Acceptor - ToApp", sessionID.TargetCompID, "MsgType", msgtype)
	return nil
}

//Notification of admin message being received from target.
func (a *fixAcceptorApplicationImpl) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	msgType, _ := msg.MsgType()
	log.Println("Acceptor - FromAdmin", sessionID.TargetCompID, "MsgType", msgType)

	// assume for a while that all login attempts will be successfull
	if msgType == enum.MsgType_LOGON {
		log.Println("Acceptor - FromAdmin", "logon")
	}

	return nil
}

//Notification of app message being received from target.
func (a *fixAcceptorApplicationImpl) FromApp(message *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	msgtype, _ := message.MsgType()
	log.Println("Acceptor - FromApp", sessionID.TargetCompID, "MsgType", msgtype)
	return nil
}
