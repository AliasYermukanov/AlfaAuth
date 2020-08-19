package authorization

import (
	"github.com/AliasYermukanov/AlfaAuth/src/dbo"
	"github.com/AliasYermukanov/AlfaAuth/src/domain"
	"github.com/AliasYermukanov/AlfaAuth/src/errors"
	satori "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
}

// Service is the interface that provides methods.
func NewService() Service {
	return &service{}
}

type Service interface {
	createUser(req *createUserRequest) (resp *createUserResponse, err error)
	authenticateUser(req *createUserRequest) (resp *authenticateUserResponse, err error)
	updateUser()
}
func (s *service) updateUser()  {
	//Any functionality that will be needed
}

func (s *service) authenticateUser(req *createUserRequest) (resp *authenticateUserResponse, err error) {
	userData, err := dbo.FindUserByPhoneNumber(req.MobilePhone)
	if err != nil {
		return nil,err
	}
	err = bcrypt.CompareHashAndPassword(userData.EncryptedPassword,[]byte(req.Password))

	if err != nil {
		Logger.Debugln("")
		errors.AccessDenied.DeveloperMessage = "wrong password"
		return nil,errors.AccessDenied
	}else {
		return &authenticateUserResponse{Uid: userData.Uid},nil
	}
}

func (s *service) createUser(req *createUserRequest) (resp *createUserResponse, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		Logger.Debugln("error encrypting password", err.Error())
		return nil, err
	}

	userData := domain.User{
		Uid:               satori.NewV1().String(),
		MobilePhone:       req.MobilePhone,
		Password:          "",
		EncryptedPassword: encryptedPassword,
		Scope:             nil,
	}

	err = dbo.SaveToElastic("users", userData.Uid, userData)
	if err != nil {
		Logger.Debugln("Elastic save error", err.Error())
		return nil, err
	}

	return &createUserResponse{Message: "success"}, nil
}
