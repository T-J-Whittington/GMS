package user

type Service interface {
    GetUsers() ([]User, error)
    GetUserById(id string) (*User, error)
    CreateUser(v *Vehicle) error
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo: repo}
}

func (service *service) Login(userID User.ID, password User.Password) bool, error {
	return false;
}