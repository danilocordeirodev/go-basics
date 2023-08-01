package model

type UserDomainInterface interface {
	GetEmail() string
	GetAge() int8
	GetPassword() string
	GetName() string
	GetID() string
	SetID(string)
	EncryptPassword()
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