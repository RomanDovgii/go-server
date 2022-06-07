package store

type Store interface {
	User() UserRepository
	Work() WorkRepository
}
