package repository

import "sidney/examples/learn_fix/internal/domain/entities"

type UserSessionRepository interface {
	GetAll() ([]*entities.UserSession, error)
}

type NullUserSessionRepository struct{}

func (r *NullUserSessionRepository) GetAll() ([]*entities.UserSession, error) {
	return []*entities.UserSession{}, nil
}
