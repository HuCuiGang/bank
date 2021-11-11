package controller

import (
	"github.com/HuCuiGang/bank/internal/storage"
	"github.com/HuCuiGang/bank/pkg/models"
	"github.com/HuCuiGang/bank/pkg/req"
	"github.com/HuCuiGang/bank/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	storage storage.Interface
}

func NewController(storage storage.Interface) (*Controller, error) {
	return &Controller{storage: storage}, nil
}


// CheckCreateUser :注册参数校验
func CheckCreateUser(c *gin.Context){
	var cu req.CreateUserReq
	err := c.BindJSON(&cu)
	if err != nil {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "BindJSON  Error",
		})
		c.Abort()
		return 
	}
	if cu.Username == "" || cu.Password == "" {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "参数错误",
		})
		c.Abort()
		return
	}
	c.Set("data",cu)
}
// CreateUser :注册
func (cont *Controller) CreateUser(c *gin.Context) {
	//name := c.PostForm("name")
	//password := c.PostForm("password")
	//user := models.User{
	//	Name:     name,
	//	Password: password,
	//}
	// 缺少参数校验
	data, ex := c.Get("data")
	if !ex {
		c.JSON(http.StatusBadRequest, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Get Data Error",
		})
		return
	}
	reqData := data.(req.CreateUserReq) //断言 判断data类型为CreateUserReq
	user := models.User{
		UserName: reqData.Username,
		Id:       "",
		Password: reqData.Password,
		Balance:  0,
	}
	//存入数据库
	userId, err := cont.storage.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Storage To User Error",
		})
		return
	}
	//formatInt := strconv.Itoa(int(userId))
	c.JSON(http.StatusOK,resp.AAAResponse{
		ResponseCommon: resp.ResponseCommon{
			Code: resp.SUCCESS,
		},
		Data:           userId,
	})
}
// CheckLoginBank :登录参数校验
func CheckLoginBank(c *gin.Context) {
	var lg req.Account
	err := c.BindJSON(&lg)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "BindJSON Error",
		})
		c.Abort()
		return
	}
	if lg.ID == "" || lg.Password == "" {
		c.JSON(http.StatusBadRequest, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "参数错误",
		})
		c.Abort()
		return
	}
	c.Set("data", lg)
}
// LoginBank :登录
func (cont *Controller) LoginBank(c *gin.Context) {
	//var user models.User
	//c.BindJSON(&user)
	//userId := c.PostForm("userId")
	//password := c.PostForm("password")
	data, ex := c.Get("data")
	if !ex {
		c.JSON(http.StatusBadRequest, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Get Data Error",
		})
		return
	}
	reqData := data.(req.Account) //断言 判断data类型为Account

	user, err := cont.storage.LoginBank(reqData.ID, reqData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Storage To LoginBank Error ",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Marshal User Error",
		})
		return
	}
	c.JSON(http.StatusOK,resp.AAAResponse{
		ResponseCommon: resp.ResponseCommon{
			Code: resp.SUCCESS,
		},
		Data:user,
	})
}
// CheckSaveMoney :存钱参数校验
func CheckSaveMoney(c *gin.Context){
	var sm req.MoneyReq
	err := c.BindJSON(&sm)
	if err != nil {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "BindJSON SaveMoneyReq Error",
		})
		c.Abort()
		return
	}
	if sm.Money == 0 || sm.UserId == "" {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "参数错误",
		})
		c.Abort()
		return
	}
	c.Set("data",sm)
}
// SaveMoney :存钱
func (cont *Controller) SaveMoney(c *gin.Context) {
	//userId := c.PostForm("userId")
	//money := c.PostForm("money")
	data, ex := c.Get("data")
	if !ex {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Get Data Error",
		})
		return
	}
	moneyReq := data.(req.MoneyReq)

	//moneyFloat, err := strconv.ParseFloat(moneyReq., 64)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//}
	err := cont.storage.SaveMoney(moneyReq.Money, moneyReq.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Storage SaveMoney Error",
		})
		return
	}
	c.JSON(http.StatusOK,resp.AAAResponse{
		ResponseCommon: resp.ResponseCommon{
			Code: resp.SUCCESS,
		},
		Data:           "ok",
	})
}
// CheckWithdrawMoney :取前参数校验
func CheckWithdrawMoney(c *gin.Context){
	var wm req.MoneyReq
	err := c.BindJSON(&wm)
	if err != nil {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "BindJSON MoneyReq Error",
		})
		c.Abort()
		return
	}
	if wm.Money == 0 || wm.UserId == "" {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "参数错误",
		})
		c.Abort()
		return
	}
	c.Set("data",wm)
}
// WithdrawMoney  :取钱
func (cont *Controller) WithdrawMoney(c *gin.Context) {
	//userId := c.PostForm("userId")
	//money := c.PostForm("money")
	data, ex := c.Get("data")
	if !ex {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Get Data Error",
		})
		c.Abort()
		return
	}
	moneyReq := data.(req.MoneyReq)
	err := cont.storage.WithdrawMoney(moneyReq.Money, moneyReq.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "storage WithdrawMoney Error",
		})
		return
	}
	c.JSON(http.StatusOK,resp.AAAResponse{
		ResponseCommon: resp.ResponseCommon{
			Code: resp.SUCCESS,
		},
		Data:           "ok",
	})
}
// CheckTransfer :校验转账参数
func CheckTransfer(c *gin.Context){
	var ts req.TransferReq
	err := c.BindJSON(&ts)
	if err != nil {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "BindJSON TransferReq Error",
		})
		return
	}
	if ts.OutUserId == "" || ts.EnterUserId == "" || ts.Money == 0 {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "参数错误",
		})
		return
	}
	c.Set("data",ts)
}
// Transfer :转账
func (cont *Controller) Transfer(c *gin.Context) {
	data, ex := c.Get("data")
	if !ex {
		c.JSON(http.StatusBadRequest,resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Get Data Error",
		})
		return
	}
	transferReq := data.(req.TransferReq)
	err := cont.storage.Transfer(transferReq.Money, transferReq.OutUserId, transferReq.EnterUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.ResponseCommon{
			Code:    resp.ERROR,
			Message: "Storage Transfer Error",
		})
		return
	}
	c.JSON(http.StatusOK,resp.AAAResponse{
		ResponseCommon: resp.ResponseCommon{
			Code: resp.SUCCESS,
		},
		Data:           "ok",
	})
}
