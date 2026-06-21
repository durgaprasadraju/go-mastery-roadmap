package solutions

import "errors"

var ErrNotFound = errors.New("not found")

// User is domain entity.
type User struct{ ID, Email string }

// UserRepository defines persistence port.
type UserRepository interface {
	GetByID(id string) (User, error)
}

// GetUserEmail is application use case.
func GetUserEmail(repo UserRepository, id string) (string, error) {
	u, err := repo.GetByID(id)
	if err != nil {
		return "", err
	}
	return u.Email, nil
}

// InMemoryUserRepo is infrastructure adapter.
type InMemoryUserRepo struct{ data map[string]User }

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{data: make(map[string]User)}
}

func (r *InMemoryUserRepo) Save(u User) { r.data[u.ID] = u }

func (r *InMemoryUserRepo) GetByID(id string) (User, error) {
	u, ok := r.data[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}