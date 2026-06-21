package solutions

// User is domain entity.
type User struct{ ID, Email string }

// UserRepository defines persistence port.
type UserRepository interface {
	GetByID(id string) (User, error)
}

// GetUserEmail is application use case.
func GetUserEmail(repo UserRepository, id string) (string, error) {
	u, err := repo.GetByID(id)
	if err != nil { return "", err }
	return u.Email, nil
}