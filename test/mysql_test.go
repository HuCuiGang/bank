package test

import (
	"github.com/HuCuiGang/bank/pkg/models"
	"testing"
)

var user = models.User{
	Name: "胡大帅",
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

