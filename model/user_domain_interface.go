package model

import "github.com/danilocordeirodev/go-basics/config/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetAge() int8
	GetPassword() string
	GetName() string
	GetID() string
	SetID(string)
	EncryptPassword()

	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email: email, 
		password: password,
		name: name, 
		age: age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name, 
		age: age,
	}
}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email: email, 
		password: password,
	}
}