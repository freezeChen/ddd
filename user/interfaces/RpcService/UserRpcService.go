package RpcService

import (
	"ddd/user/domain/repository"
)

type UserRpcService struct {
	userRepo *repository.UserRepository
}



func (svc UserRpcService) GetUserById(uid int64)  {

}
