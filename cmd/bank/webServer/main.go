package main

import (
	"log"

	"github.com/HuCuiGang/bank/internal/conf"
	"github.com/HuCuiGang/bank/internal/controller"
	"github.com/HuCuiGang/bank/internal/routers"
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

	// 初始化controller
	controller, _ := controller.NewController(storage)

	//初始化router
	router, _ := routers.NewRouter(controller)
	r := router.SetupRouter()

	//启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
