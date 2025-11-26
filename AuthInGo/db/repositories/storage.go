package db

type Storage struct { //facilitates dependency injection for all repositories
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}
