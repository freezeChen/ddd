package RpcService

import (
	userDomain "ddd/user/domain"
	"ddd/user/domain/repository"
)

type UserRpcService struct {
	userRepo *repository.UserRepository
}

func (svc UserRpcService) GetUserById(uid int64) userDomain.User {

}
