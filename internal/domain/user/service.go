package user

type UserService struct {
	User       User
	repository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}
func (u *UserService) UserByUid(uid string) {

}

//checkAuth

//checkLogin
