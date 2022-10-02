package dbjsonlocal

import (
	"encoding/json"
	"io/ioutil"
	"sidney/examples/learn_fix/internal/domain/entities"
	"sidney/examples/learn_fix/internal/domain/repository"
	"sidney/examples/learn_fix/internal/utils/fileutils"
)

type jsonUserSessionRepository struct {
	data []*entities.UserSession
}

func NewUserSessionRepositoryFromJsonFile(jsonFile string) (repository.UserSessionRepository, error) {
	r := &jsonUserSessionRepository{
		data: make([]*entities.UserSession, 0),
	}

	if err := r.parseJson(jsonFile); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *jsonUserSessionRepository) parseJson(jsonFile string) error {
	file, err := fileutils.OpenFile(jsonFile)
	if err != nil {
		return err
	}

	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, &r.data); err != nil {
		return err
	}

	return nil
}

func (r *jsonUserSessionRepository) GetAll() ([]*entities.UserSession, error) {
	return r.data, nil
}
