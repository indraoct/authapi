package user

type Service interface {
	RegisterNewUser(*User) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (svc *service) RegisterNewUser(u *User) error {
	return svc.repo.StoreUser(u)
}
