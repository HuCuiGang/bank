package test

import (
	"fmt"
	"github.com/HuCuiGang/bank/internal/conf"
	"github.com/HuCuiGang/bank/internal/storage"
	"github.com/HuCuiGang/bank/pkg/models"
	"testing"
)

var user = models.User{
	UserName: "胡大帅",
	Id: "1",
	Password: "123456",
	Balance: 0,
}


func TestCreatUser(t *testing.T){
	//mySQLStorage,err :=
	//if err!= nil{
	//	fmt.Println(err)
	//}
	//err = mySQLStorage.CreateUser(user)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func TestQueryUser(t *testing.T){
	/// 初始化配置文件
	//err := conf.InitAppConfig()
	//if err != nil {
	//	panic(err)
	//}
	conf := conf.MySQLConfig{
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "bank",
	}
	// 初始化storage
	storage, err := storage.NewMySQLStorage(conf)
	if err != nil {
		panic(err)
	}

	queryUser, err := storage.QueryUser("9")
	if err != nil {
		return
	}

	fmt.Println(queryUser)

}


