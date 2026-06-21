package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/93-migrations/exercises/solutions"
)

type memRepo struct{ users map[string]solutions.User }

func (m memRepo) GetByID(id string) (solutions.User, error) {
	return m.users[id], nil
}

func TestGetUserEmail(t *testing.T) {
	repo := memRepo{users: map[string]solutions.User{"1": {ID: "1", Email: "a@b.com"}}}
	email, err := solutions.GetUserEmail(repo, "1")
	if err != nil || email != "a@b.com" {
		t.Fatalf("email=%q err=%v", email, err)
	}
}
