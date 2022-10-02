package configuration

import (
	"encoding/json"
	"io/ioutil"
	"sidney/examples/learn_fix/internal/utils/fileutils"
)

const (
	ConnectionType_Initiator = "initiator"
	ConnectionType_Acceptor  = "acceptor"

	DatabaseKind_Json   = "json"
	DatabaseKind_SQLite = "sqlite"
)

type AppGlobalConfig struct {
	ConnectionType string          `json:"connectionType,omitempty"`
	WorkDir        string          `json:"workDir,omitempty"`
	Fix            *FixConfig      `json:"fix,omitempty"`
	Database       *DatabaseConfig `json:"database,omitempty"`
}

type FixConfig struct {
	ConfigFile         string `json:"configFile,omitempty"`
	DataDictionaryFile string `json:"dataDictionaryFile,omitempty"`
}

type DatabaseConfig struct {
	UseLocal           bool                `json:"useLocal,omitempty"`
	Kind               string              `json:"kind,omitempty"`
	LocalDatabaseFiles *LocalDatabaseFiles `json:"localDatabaseFiles,omitempty"`
}

type LocalDatabaseFiles struct {
	Users string `json:"users,omitempty"`
}

func ParseConfigFromJson(jsonConfigFile string) (*AppGlobalConfig, error) {
	file, err := fileutils.OpenFile(jsonConfigFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	result := &AppGlobalConfig{}
	if err = json.Unmarshal(jsonBytes, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AppGlobalConfig) IsConnectionAcceptor() bool {
	return a.ConnectionType == ConnectionType_Acceptor
}

func (a *AppGlobalConfig) IsConnectionInitiator() bool {
	return a.ConnectionType == ConnectionType_Initiator
}

func (a *AppGlobalConfig) GetFormattedFixConfigFileName() string {
	return fileutils.GetFormattedFullFileName(a.WorkDir, a.Fix.ConfigFile)
}

func (a *AppGlobalConfig) GetFormattedDataDictionaryFileName() string {
	return fileutils.GetFormattedFullFileName(a.WorkDir, a.Fix.DataDictionaryFile)
}

func (a *AppGlobalConfig) GetFormattedUsersJsonFileName() string {
	if a.Database == nil || !a.Database.UseLocal || a.Database.LocalDatabaseFiles == nil {
		return ""
	}
	return fileutils.GetFormattedFullFileName(a.WorkDir, a.Database.LocalDatabaseFiles.Users)
}

func (a *AppGlobalConfig) IsJsonLocalDatabase() bool {
	if a.Database == nil {
		return false
	}
	return a.Database.UseLocal && a.Database.Kind == DatabaseKind_Json
}

func (a *AppGlobalConfig) IsSQLiteLocalDatabase() bool {
	return a.Database.UseLocal && a.Database.Kind == DatabaseKind_SQLite
}
