package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HuCuiGang/bank/internal/client"
	"github.com/HuCuiGang/bank/pkg/models"
	"github.com/HuCuiGang/bank/pkg/resp"
	"github.com/mitchellh/mapstructure"
)
var user models.User

func Run() error {
	for {
		Cli()      //显示页面
		var in int //功能选择
		fmt.Scanln(&in)
		switch in {
		case 1: //选择登录时
			fmt.Println("请输出ID:")
			var id string
			fmt.Scanln(&id)
			fmt.Println("请输入密码:")
			var p string
			fmt.Scanln(&p)
			result, httpCode, err := client.CliLoginBank(id, p, "127.0.0.1:8080/v1/loginBank")
			if err != nil || httpCode != 200 {
				fmt.Println("登录失败，", err)
				break
			}
			fmt.Println("登录成功！")

			var data resp.AAAResponse
			err = json.Unmarshal(result, &data)
			if err != nil {
				log.Fatal(err)
				return err
			}
			userDataMap := data.Data.(map[string]interface{})

			//将map转换成struct
			mapstructure.Decode(userDataMap,&user)

		loop:
			for {
				//登录成功后显示功能列表
				Opshow(user)
				var opin int //功能选择
				fmt.Scanln(&opin)
				switch opin {
				case 1: //存钱
					fmt.Println("请输入存入金额：")
					var money float64
					fmt.Scanln(&money)
					_, httpCode, err := client.CliSaveMoney(user.Id, money, "127.0.0.1:8080/v1/saveMoney")
					if err != nil || httpCode != 200{
						fmt.Println("存钱失败，", err)
						break
					}
					fmt.Println("存前完成")
					Renew(user.Id,user.Password) //更新信息
				case 2: //取钱
					fmt.Println("请输入取款金额：")
					var wdMoney float64
					fmt.Scanln(&wdMoney)
					_, httpCode, err := client.CliSaveMoney(user.Id, wdMoney, "127.0.0.1:8080/v1/withdrawMoney")
					if err != nil || httpCode != 200 {
						fmt.Println("取钱失败，", err)
						break
					}
					fmt.Println("取钱完成")
					Renew(user.Id,user.Password) //更新信息
				case 3: //转账
					fmt.Println("请输入转入账户ID：")
					var id string
					fmt.Scanln(&id)
					fmt.Println("请输入转入金额：")
					var tfMoney float64 //转入金额
					fmt.Scanln(&tfMoney)
					_, httpCode, err := client.CliTransfer(user.Id, id, tfMoney, "127.0.0.1:8080/v1/transfer")
					if err != nil || httpCode != 200 {
						fmt.Println("转账失败，", err)
						break
					}
					fmt.Println("转账完成！")
					Renew(user.Id,user.Password) //更新信息
				case 4: //返回上一步
					break loop
				}
			}
		case 2: //注册
			fmt.Println("请输入姓名：")
			var name string
			fmt.Scanln(&name)
			var password string
			fmt.Println("请输入密码：")
			fmt.Scanln(&password)

			result, httpCode, err := client.CliCreatUser(name, password, "127.0.0.1:8080/v1/createUser")
			if err != nil || httpCode != 200 {
				fmt.Println("注册失败，", err)
				break
			}
			var data resp.AAAResponse
			err = json.Unmarshal(result, &data)
			if err != nil {
				fmt.Println("注册失败", err)
				break
			}
			fmt.Println("注册成功！您的ID为：", data.Data)
		default:
			fmt.Println("fuckoff！输入错误，请重新输入！")
		}
	}

	return nil
}

// Cli 银行登录入口显示
//func Cli() {
//	fmt.Println("欢迎来到小刚银行~")
//	fmt.Println("登录请按1")
//	fmt.Println("注册请按2")
//}
//
//// Opshow 登录后功能显示 :
//func Opshow(u models.User) {
//	fmt.Println(u.UserName+"欢迎您！"+"您的账户id：", u.Id+"账户余额：", u.Balance)
//	fmt.Println("存钱请按1")
//	fmt.Println("取钱请按2")
//	fmt.Println("转账请按3")
//	fmt.Println("返回上一步请按4")
//}

// Renew :更新User信息
func Renew(id string ,p string){
	result, httpCode, err := client.CliLoginBank(id, p, "127.0.0.1:8080/v1/loginBank")
	if err != nil || httpCode != 200 {
		fmt.Println("刷新失败，", err)

	}

	var data resp.AAAResponse
	err = json.Unmarshal(result, &data)
	if err != nil {
		log.Fatal(err)
	}
	userDataMap := data.Data.(map[string]interface{})
	//将map转换成struct
	mapstructure.Decode(userDataMap,&user)
}