package server

import (
	"fmt"

	"github.com/HuCuiGang/bank/internal/storage"
	"github.com/HuCuiGang/bank/pkg/models"
)

type Server struct {
	storage storage.Interface
}

func NewServer(storage storage.Interface) *Server {
	return &Server{storage: storage}
}

func (s *Server) Run() error {
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
			user, err := s.storage.LoginBank(id, p)
			if err != nil {
				fmt.Println("登录失败，", err)
				break
			}
			fmt.Println("登录成功！")
		loop:
			for {
				//查询用户信息
				user,err = s.storage.QueryUser(user.Id)
				if err!=nil {
					fmt.Println(err)
					break
				}
				//登录成功后显示功能列表
				Opshow(user)
				var opin int //功能选择
				fmt.Scanln(&opin)
				switch opin {
				case 1: //存钱
					fmt.Println("请输入存入金额：")
					var money float64
					fmt.Scanln(&money)
					err := s.storage.SaveMoney(money, user.Id)
					if err != nil {
						fmt.Println("存钱失败，", err)
						break
					}
					fmt.Println("存前完成")
				case 2: //取钱
					fmt.Println("请输入取款金额：")
					var wdMony float64
					fmt.Scanln(&wdMony)
					err := s.storage.WithdrawMoney(wdMony, user.Id)
					if err != nil {
						fmt.Println("取前失败，", err)
						break
					}
					fmt.Println("取钱完成")
				case 3: //转账
					fmt.Println("请输入转入账户ID：")
					var id string
					fmt.Scanln(&id)
					fmt.Println("请输入转入金额：")
					var tfMoney float64 //转入金额
					fmt.Scanln(&tfMoney)
					err := s.storage.Transfer(tfMoney, user.Id, id)
					if err != nil {
						fmt.Println("转账失败，", err)
						break
					}
					fmt.Println("转账完成！")
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
			user := models.User{
				Name:     name,
				Id:       "",
				Password: password,
			}

			createUserId, err := s.storage.CreateUser(user)
			if err != nil {
				fmt.Println(err)
				break
			}
			if err != nil {
				fmt.Println("注册失败，",err)
				break
			}
			fmt.Println("注册成功！ID为：",createUserId)
		default:
			fmt.Println("fuckoff！输入错误，请重新输入！")
		}
	}

	return nil
}

//银行登录入口显示
func Cli() {
	fmt.Println("欢迎来到小刚银行~")
	fmt.Println("登录请按1")
	fmt.Println("注册请按2")
}

//登录后功能显示
func Opshow(u models.User) {
	fmt.Println(u.Name+"欢迎您！"+"您的账户id：", u.Id+"账户余额：", u.Balance)
	fmt.Println("存钱请按1")
	fmt.Println("取钱请按2")
	fmt.Println("转账请按3")
	fmt.Println("返回上一步请按4")
}
