package cmd

import (
	"github.com/HuCuiGang/bank/internal/conf"
	"github.com/HuCuiGang/bank/internal/server"
	"github.com/HuCuiGang/bank/internal/storage"
)

func main1() {
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
