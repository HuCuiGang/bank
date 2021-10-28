package storage

import "github.com/HuCuiGang/bank/pkg/models"

type Interface interface {
	CreateUser(user models.User) error //创建用户
	LoginBank(userId string ,password string) error //登录
	SaveMoney(money float64 ,userId string) error	//存钱
	WithdrawMoney(money float64 ,userId string) error //取钱
	Transfer(money float64,outUserId string ,inUserId string) error //转账
}
