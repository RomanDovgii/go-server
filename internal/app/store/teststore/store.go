package teststore

import (
	"database/sql"

	"github.com/RomanDovgii/go-restapi/internal/app/model"
	"github.com/RomanDovgii/go-restapi/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
