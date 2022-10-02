package connection

import (
	"errors"
	"sidney/examples/learn_fix/internal/configuration"
	"sidney/examples/learn_fix/internal/domain/repository"
	"sidney/examples/learn_fix/internal/utils/fileutils"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/datadictionary"
)

type FixConnectionParameters struct {
	GlobalConfig        *configuration.AppGlobalConfig
	AppSettings         *quickfix.Settings
	DataDictionary      *datadictionary.DataDictionary
	MessageStoreFactory quickfix.MessageStoreFactory
	LogFactory          quickfix.LogFactory

	UsersRepository repository.UserSessionRepository
}

func NewFixConnectionParameters(config *configuration.AppGlobalConfig, usersRepository repository.UserSessionRepository) (*FixConnectionParameters, error) {
	parameters := &FixConnectionParameters{
		GlobalConfig:        config,
		AppSettings:         nil,
		DataDictionary:      nil,
		MessageStoreFactory: nil,
		LogFactory:          nil,
		UsersRepository:     usersRepository,
	}

	if err := parameters.parseQuickfixSettings(); err != nil {
		return nil, err
	}

	if err := parameters.parseDataDictionary(); err != nil {
		return nil, err
	}

	if err := parameters.parseFilePersistence(); err != nil {
		return nil, err
	}

	return parameters, nil
}

func (p *FixConnectionParameters) parseQuickfixSettings() error {
	configFile, err := fileutils.OpenFile(p.GlobalConfig.GetFormattedFixConfigFileName())
	if err != nil {
		return err
	}

	defer configFile.Close()

	appSettings, err := quickfix.ParseSettings(configFile)
	if err != nil {
		return err
	}

	p.AppSettings = appSettings
	return nil
}

func (p *FixConnectionParameters) parseDataDictionary() error {

	if p.isDataDictionaryFromFlags() {
		p.updateDataDictionarySettingsFromFlags()
		return nil
	}

	dictionaryFileName := p.getDataDictionaryFileNameFromSettings()
	if len(dictionaryFileName) == 0 {
		return errors.New("data dictionary is missing")
	}

	return nil
}

func (p *FixConnectionParameters) parseFilePersistence() error {

	logFactory, err := quickfix.NewFileLogFactory(p.AppSettings)
	if err != nil {
		return err
	}

	p.LogFactory = logFactory
	p.MessageStoreFactory = quickfix.NewFileStoreFactory(p.AppSettings)

	return nil
}

func (p *FixConnectionParameters) isDataDictionaryFromFlags() bool {
	return len(p.GlobalConfig.Fix.DataDictionaryFile) > 0
}

func (p *FixConnectionParameters) updateDataDictionarySettingsFromFlags() {
	p.AppSettings.GlobalSettings().Set("UseDataDictionary", "Y")
	p.AppSettings.GlobalSettings().Set("DataDictionary", p.GlobalConfig.GetFormattedDataDictionaryFileName())
}

func (p *FixConnectionParameters) getDataDictionaryFileNameFromSettings() string {
	if p.AppSettings != nil {
		if p.getBooleanFromSettings("UseDataDictionary", true) {
			return p.getStringFromSettings("DataDictionary", "")
		}
	}

	return ""
}

func (p *FixConnectionParameters) getBooleanFromSettings(setting string, defaultValue bool) bool {

	if configValue, err := p.AppSettings.GlobalSettings().BoolSetting(setting); err == nil {
		return configValue
	}

	return defaultValue
}

func (p *FixConnectionParameters) getStringFromSettings(setting, defaultValue string) string {
	if configValue, err := p.AppSettings.GlobalSettings().Setting(setting); err == nil {
		return configValue
	}

	return defaultValue
}
