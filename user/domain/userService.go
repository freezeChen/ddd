package domain

type UserService interface {
	GetUserById(uid int64) *User
}

type userService struct {
}

func (u userService) GetUserById(uid int64) *User {
	return &User{
		Id:      uid,
		Name:    "晓东",
		Balance: 200,
	}
}
