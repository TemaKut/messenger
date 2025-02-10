package mappers

import (
	"github.com/TemaKut/messenger/internal/services/auth/internal/dto"
	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
	"github.com/TemaKut/messenger/pkg/utils"
)

func UnregisteredUserFromRegisterUserRequest(req *authv1.RegisterUserRequest) *dto.UnregisteredUser {
	return &dto.UnregisteredUser{
		Username:  req.GetUsername(),
		FirstName: utils.ValueToPtr(req.GetFirstName()),
		LastName:  utils.ValueToPtr(req.GetLastName()),
		Email:     utils.ValueToPtr(req.GetEmail()),
		Phone:     utils.ValueToPtr(req.GetPhone()),
		Password:  utils.ValueToPtr(req.GetPassword()),
	}
}

// TODO Сделать нормальный маппер
func UserToProto(u *dto.User) *authv1.User {
	return &authv1.User{
		Id:        u.Id,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}
