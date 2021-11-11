package routers

import (
	"net/http"

	"github.com/HuCuiGang/bank/internal/controller"
	"github.com/gin-gonic/gin"
)


type Router struct {
	controller *controller.Controller
}

func NewRouter(controller *controller.Controller) (*Router, error) {
	return &Router{controller: controller},nil
}

func (rt *Router) SetupRouter() *gin.Engine{
	r := gin.Default()

	r.Static("/statics","statics")

	r.LoadHTMLGlob("../../../templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK,"register.html",nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/createUser",controller.CheckCreateUser,rt.controller.CreateUser)
		v1Group.POST("/loginBank",controller.CheckLoginBank,rt.controller.LoginBank)
		v1Group.PUT("/saveMoney",controller.CheckSaveMoney,rt.controller.SaveMoney)
		v1Group.PUT("/withdrawMoney",controller.CheckWithdrawMoney,rt.controller.WithdrawMoney)
		v1Group.PUT("/transfer",controller.CheckTransfer,rt.controller.Transfer)
	}


	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"message":"你要去哪里呀~",
		})
	})
    return r
}
