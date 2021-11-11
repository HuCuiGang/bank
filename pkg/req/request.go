package req

// Account :登录参数结构
type Account struct {
	ID string `json:"id"`
	Password string `json:"password"`
}

// CreateUserReq :注册参数结构
type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MoneyReq :存取钱参数结构
type MoneyReq struct {
	UserId string	`json:"user_id"`
	Money float64 `json:"money"`
}

type TransferReq struct {
	OutUserId string `json:"out_user_id"`
	EnterUserId string `json:"enter_user_id"`
	Money float64 `json:"money"`
}