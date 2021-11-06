package storage

import (
	"errors"
	"fmt"

	"github.com/HuCuiGang/bank/internal/conf"
	"github.com/HuCuiGang/bank/pkg/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLStorage struct {
	db *sqlx.DB
}

func NewMySQLStorage(conf conf.MySQLConfig) (*MySQLStorage, error) {
	//db,err :=sqlx.Open("mysql","root:123456@tcp(127.0.0.1)/bank")
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName))
	if err != nil {
		return nil, err
	}

	return &MySQLStorage{db: db}, nil
}

func (m *MySQLStorage) CreateUser(user models.User) (int64,error) {
	result, err := m.db.Exec("insert into user(name,password) value(?,?) ;", user.Name, user.Password)
	if err != nil {
		return 0,err
	}
	insertId, _ :=result.LastInsertId() //返回最后一行ID
	return insertId,nil
}

func (m *MySQLStorage) LoginBank(userId string, password string) (models.User,error) {
	var users []models.User

	err := m.db.Select(&users, "select id ,name ,balance from user where id=? and password=? ;", userId, password)
	if err != nil {
		return models.User{},err
	}
	if len(users)<1{
		return models.User{},errors.New("登录失败")
	}
	return users[0],nil
}

func (m *MySQLStorage) SaveMoney(money float64, userID string) error {
	var users []models.User

	err := m.db.Select(&users, "select balance from user where id=? ;",userID)
	if err != nil {
		return  err
	}

	fmt.Println(users)
	balance := users[0].Balance + money

	result, err := m.db.Exec("update user set balance=? where id=? ;", balance, userID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected() //受影响的行数
	if rowsAffected == 0 {
		return errors.New("存钱失败")
	}
	return nil
}

func (m *MySQLStorage) WithdrawMoney(money float64, userId string) error {
	var user []models.User

	err := m.db.Select(&user, "select balance from user where id = ? ;", userId)
	if err != nil {
		return err
	}
	if user[0].Balance-money < 0 {
		return errors.New("余额不足")
	}

	balance := user[0].Balance - money

	m.db.Exec("update user set balance=? where id = ? ;", balance,userId)

	return nil
}

func (m *MySQLStorage) Transfer(money float64, outUserId string, inUserId string) error {
	var outuser []models.User
	var inUser []models.User

	err := m.db.Select(&outuser, "select balance from user where id =? ;", outUserId)
	if err != nil {
		return err
	}
	if outuser[0].Balance-money < 0 {
		return errors.New("余额不足")
	}

	outuserbalance := outuser[0].Balance - money

	_, err = m.db.Exec("update user set balance=? where id = ? ;", outuserbalance,outUserId)
	if err != nil {
		return err
	}

	err = m.db.Select(&inUser, "select balance from user where id = ? ;", inUserId)
	if err != nil {
		return err
	}

	inUserbalance := inUser[0].Balance + money

	m.db.Exec("update user set balance= ? where id= ? ;", inUserbalance, inUserId)

	return nil
}

func (m *MySQLStorage)QueryUser(userId string) (models.User,error){
	var users []models.User
	err := m.db.Select(&users, "select * from user where id=? ;", userId)
	if err != nil {
		return models.User{}, err
	}
	return users[0],nil
}
