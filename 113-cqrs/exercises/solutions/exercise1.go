package solutions

// CreateUserCommand is a write model command.
type CreateUserCommand struct{ ID, Email string }

// UserView is read model projection.
type UserView struct{ ID, Email string }

// UserStore holds projections.
type UserStore struct{ views map[string]UserView }

func NewUserStore() *UserStore { return &UserStore{views: make(map[string]UserView)} }

func (s *UserStore) Handle(cmd CreateUserCommand) {
	s.views[cmd.ID] = UserView{ID: cmd.ID, Email: cmd.Email}
}

func (s *UserStore) Get(id string) (UserView, bool) {
	v, ok := s.views[id]
	return v, ok
}