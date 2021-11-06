package storage

import (
	"errors"

	"github.com/HuCuiGang/bank/pkg/models"
)

type Mem struct {
	Users map[string]models.User
}


func (m *Mem) NewMem() Mem {
	return Mem{
		make(map[string]models.User),
	}
}

func (m *Mem) CreateUser(u models.User) error {
	m.Users[u.Id] = u
	return nil
}

func (m *Mem) LoginBank(userId string , password string) (user models.User,err error)   {
	user , ex := m.Users[userId]
	if !ex || password !=user.Password {
		return user , errors.New("username does not exist or password is wrong")
	}
	return user ,nil
}

func (m *Mem) SaveMoney(money float64 , userId string ) error {
	user,ex := m.Users[userId]
	if !ex{
		return errors.New("user does not exist")
	}
	user.Balance += money
	m.Users[userId] = user

	return nil
}

func (m *Mem) WithdrawMoney(money float64,userId string ) error {
	user, ex := m.Users[userId]
	if !ex || user.Balance - money<0  {
		return errors.New("sorry, your balance is insufficient")
	}
	user.Balance -= money
	m.Users[userId]=user
	return nil
}

func (m *Mem) Transfer(money float64 ,outUserId string ,inUserId string) error   {
	outUser , ex := m.Users[outUserId]
	if !ex || outUser.Balance - money < 0  {
		return errors.New("transfer failed")
	}
	inUser , ex := m.Users[inUserId]
	if !ex {
		return errors.New("transfer failed")
	}
	outUser.Balance -= money
	m.Users[outUserId] = outUser
	inUser.Balance += money
	m.Users[inUserId] = inUser
	return nil

}