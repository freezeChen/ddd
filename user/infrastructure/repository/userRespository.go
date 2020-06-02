package repository

import userDomain "learn/ddd/user/domain"

type UserRepository struct {
}

func (repo UserRepository) GetUserById(uid string) userDomain.User {

}
