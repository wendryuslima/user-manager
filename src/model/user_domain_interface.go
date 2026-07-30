package model

type UserDomainInterface interface {
	GetEmail() string
	GetID() string
	GetPassword() string
	GetName() string
	GetAge() int8
	SetID(string)
	EncryptPassword()
}

func NewUserDomain(
	email, id, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		id:       id,
		password: password,
		name:     name,
		age:      age,
	}
}
