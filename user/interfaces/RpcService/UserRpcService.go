package RpcService

import (
	userDomain "learn/ddd/user/domain"
	"learn/ddd/user/infrastructure/repository"
)

type UserRpcService struct {
	userRepo *repository.UserRepository
}

func (svc UserRpcService) GetUserById(uid int64) userDomain.User {

}
