package main

import (
	"github.com/HuCuiGang/bank/internal/conf"
	"github.com/HuCuiGang/bank/internal/server"
	"github.com/HuCuiGang/bank/internal/storage"
)

func main() {
	// 初始化配置文件
	err := conf.InitAppConfig()
	if err != nil {
		panic(err)
	}

	// 初始化storage
	storage, err := storage.NewMySQLStorage(conf.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}


	// 初始化控制层
	server := server.NewServer(storage)
	if err := server.Run(); err != nil {
		panic(err)
	}

}

//func main() {
//	var uid int = 0 //初始化用户ID
//	m := &storage.Mem{Users: make(map[string]models.User)}
//	for {
//		Cli()      //显示页面
//		var in int //功能选择
//		fmt.Scanln(&in)
//		switch in {
//		case 1:
//			fmt.Println("请输入ID")
//			var id string
//			fmt.Scanln(&id)
//			fmt.Println("请输入密码：")
//			var p string
//			fmt.Scanln(&p)
//			user , err := m.LoginBank(id,p)
//			if err != nil{
//				fmt.Println(err)
//				break
//			}
//			fmt.Println("登录成功！")
//		loop:
//			for {
//				//功能列表
//				Opshow(m.Users[user.Id])
//				var opin int //功能选择
//				fmt.Scanln(&opin)
//				switch opin {
//				case 1: //存钱
//					fmt.Println("请输入存入金额：")
//					var money float64
//					fmt.Scanln(&money)
//					err := m.SaveMoney(money,user.Id)
//					if err != nil {
//						fmt.Println(err)
//						break
//					}
//					fmt.Println("存钱完成！")
//
//				case 2: //取钱
//					fmt.Println("请输入取款金额：")
//					var wdMony float64
//					fmt.Scanln(&wdMony)
//					err := m.WithdrawMoney(wdMony,user.Id)
//					if err != nil {
//						fmt.Println(err)
//						break
//					}
//					fmt.Println("取钱完成")
//
//				case 3: //转账
//					fmt.Println("请输入转入账户ID：")
//					var id string
//					fmt.Scanln(&id)
//					fmt.Println("请输入转入金额：")
//					var tfMoney float64 //转入金额
//					fmt.Scanln(&tfMoney)
//					err := m.Transfer(tfMoney,user.Id,id)
//					if err != nil {
//						fmt.Println(err)
//						break
//					}
//					fmt.Println("转账完成！")
//
//				case 4: //返回上一步
//					break loop
//				default:
//					fmt.Println("输入错误，请重新输入！")
//				}
//			}
//		case 2: //注册
//			fmt.Println("请输入姓名：")
//			var n string //姓名
//			fmt.Scanln(&n)
//			uid++        //ID
//			var p string //密码
//			fmt.Println("请输入密码：")
//			fmt.Scanln(&p)
//			u := models.User{
//				Name:     n,
//				Id:       strconv.Itoa(uid),
//				Password: p,
//				Money:    0,
//			}
//			err := m.CreateUser(u)
//			if err != nil{
//				fmt.Println(err)
//				break
//			}
//			fmt.Println("注册成功！ID为：", u.Id)
//		default:
//			fmt.Println("输入错误，请重新输入！")
//		}
//	}
//}
//
////银行登录入口显示
//func Cli() {
//	fmt.Println("欢迎来到小刚银行~")
//	fmt.Println("登录请按1")
//	fmt.Println("注册请按2")
//}
//
////登录后功能显示
//func Opshow(u models.User) {
//	fmt.Println(u.Name+"欢迎您！"+"您的账户id：", u.Id+"账户余额：", u.Money)
//	fmt.Println("存钱请按1")
//	fmt.Println("取钱请按2")
//	fmt.Println("转账请按3")
//	fmt.Println("返回上一步请按4")
//}
