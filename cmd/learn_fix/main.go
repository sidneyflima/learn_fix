package main

import (
	"errors"
	"log"
	"sidney/examples/learn_fix/internal/application"
	"sidney/examples/learn_fix/internal/cli"
	"sidney/examples/learn_fix/internal/configuration"
	"sidney/examples/learn_fix/internal/connection"
	"sidney/examples/learn_fix/internal/database/dbjsonlocal"
	"sidney/examples/learn_fix/internal/domain/repository"
)

func GetUserSessionRepositoryFromConfig(config *configuration.AppGlobalConfig) (repository.UserSessionRepository, error) {
	if config.IsJsonLocalDatabase() {
		return dbjsonlocal.NewUserSessionRepositoryFromJsonFile(config.GetFormattedUsersJsonFileName())
	}

	return &repository.NullUserSessionRepository{}, nil
}

func CreateConnectionFromType(parameters *connection.FixConnectionParameters) (connection.FixConnection, error) {
	if parameters.GlobalConfig.IsConnectionAcceptor() {
		return connection.NewFixAcceptorFactory().CreateNewAcceptor(parameters, application.NewFixAcceptorApplication())
	} else if parameters.GlobalConfig.IsConnectionInitiator() {
		return connection.NewFixInitiatorFactory().CreateNewInitiator(parameters, application.NewFixInitiatorApplication())
	} else {
		return nil, errors.New("connection type is not valid")
	}
}

func waitConsoleApp() {
	select {}
}

func main() {
	appConfig, err := cli.ConfigureFromCommandLineFlags()
	if err != nil {
		log.Fatalln("Could not configure application", err)
	}

	usersRepository, err := GetUserSessionRepositoryFromConfig(appConfig)
	if err != nil {
		log.Fatalln("Could not create users repository", err)
	}

	parameters, err := connection.NewFixConnectionParameters(appConfig, usersRepository)
	if err != nil {
		log.Fatalln("Could not create connection parameters", err)
	}

	fixConnection, err := CreateConnectionFromType(parameters)
	if err != nil {
		log.Fatalln("Create connection error:", err)
	}

	if err := fixConnection.Start(); err != nil {
		log.Fatalln("Could not start fix connection", err)
	}

	port, ok := fixConnection.SocketAcceptPort()
	if ok {
		log.Println("Connection type", parameters.GlobalConfig.ConnectionType, "started and running at port", port, "...")
	} else {
		log.Println("Connection type", parameters.GlobalConfig.ConnectionType, "started and running...")
	}

	waitConsoleApp()
}
