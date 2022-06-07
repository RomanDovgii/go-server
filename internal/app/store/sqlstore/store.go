package sqlstore

import (
	"database/sql"

	"github.com/RomanDovgii/go-restapi/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	workRepository *WorkRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Work() store.WorkRepository {
	if s.workRepository != nil {
		return s.workRepository
	}

	s.workRepository = &WorkRepository{
		store: s,
	}

	return s.workRepository
}
