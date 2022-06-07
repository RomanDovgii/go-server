package store

import "github.com/RomanDovgii/go-restapi/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type WorkRepository interface {
	Create(*model.Work) error
	Find(int) (*model.Work, error)
	FindByName(string) (*model.Work, error)
	Delete(int, int) error
	FindAll(int, int) ([]model.Work, error)
	FindAllByName(string, int, int) ([]model.Work, error)
}
