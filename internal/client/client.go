package client

import (
	"github.com/HuCuiGang/bank/pkg/req"
	"github.com/dollarkillerx/urllib"
)

func CliCreatUser(username string, password string, url string) ([]byte, int, error) {
	httpCode, bytes, err := urllib.Post(url).
		SetJsonObject(req.CreateUserReq{
		Username: username,
		Password: password,
	}).Byte()
	if err != nil {
		return nil, 0, err
	}
	return bytes, httpCode, nil
}

func CliLoginBank(id string, password string, url string) ([]byte, int, error) {
	post := urllib.Post(url)
	jsonData := req.Account{
		ID:       id,
		Password: password,
	}
	json := post.SetJsonObject(jsonData)
	httpCode, data, err := json.Byte()
	if err != nil {
		return nil, 0, err
	}
	return data, httpCode, nil
}

func CliSaveMoney(userId string, money float64, url string) ([]byte, int, error) {
	post := urllib.Put(url)
	jsonData :=req.MoneyReq{
		UserId: userId,
		Money:  money,
	}
	json := post.SetJsonObject(jsonData)
	httpCode, data, err := json.Byte()
	if err != nil {
		return nil, 0, err
	}
	return data, httpCode, nil
}

func CliWithdrawMoney(userId string, money float64, url string) ([]byte, int, error) {
	post := urllib.Put(url)
	jsonData := req.MoneyReq{
		UserId: userId,
		Money:  money,
	}
	json := post.SetJsonObject(jsonData)
	httpCode, data, err := json.Byte()
	if err != nil {
		return nil, 0, err
	}
	return data, httpCode, nil
}

func CliTransfer(outUserId string, enterUserId string, money float64, url string) ([]byte, int, error) {
	post := urllib.Put(url)
	jsonData := req.TransferReq{
		OutUserId:   outUserId,
		EnterUserId: enterUserId,
		Money:       money,
	}
	json := post.SetJsonObject(jsonData)
	httpCode, data, err := json.Byte()
	if err != nil {
		return nil, 0, err
	}
	return data, httpCode, nil
}
