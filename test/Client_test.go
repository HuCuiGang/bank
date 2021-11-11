package test

import (
	"encoding/json"
	"fmt"
	"github.com/HuCuiGang/bank/pkg/resp"
	"log"
	"testing"

	"github.com/HuCuiGang/bank/internal/client"
)

func TestCliLoginBank(t *testing.T)  {
	result, httpCode, err := client.CliLoginBank("33","666","127.0.0.1:8080/v1/loginBank")
	if err != nil {
		log.Fatal(err)
		return
	}
	var data resp.AAAResponse
	err = json.Unmarshal(result, &data)
	fmt.Println(string(result))
	log.Printf("httpCode:%d \n",httpCode)
	fmt.Println(data.Data)
}

func TestCliCreatUser(t *testing.T)  {
	result, httpCode, err := client.CliCreatUser("二哥","666","127.0.0.1:8080/v1/createUser")
	if err != nil {
		log.Fatal(err)
		return
	}
	var data resp.AAAResponse
	err = json.Unmarshal(result, &data)
	fmt.Println(string(result))
	log.Printf("httpCode:%d \n",httpCode)
	fmt.Println(data.Data)
}

func TestCliSaveMoney(t *testing.T){
	result, httpCode, err := client.CliSaveMoney("30","500","127.0.0.1:8080/v1/saveMoney")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("httpCode:%d \n",httpCode)
	fmt.Println(result)
}

func TestCliWithdrawMoney(t *testing.T)  {
	result, httpCode, err := client.CliWithdrawMoney("30","200","127.0.0.1:8080/v1/withdrawMoney")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("httpCode:%d \n",httpCode)
	fmt.Println(result)
}

func TestCliTransfer(t *testing.T)  {
	result, httpCode, err := client.CliTransfer("30","29","100","127.0.0.1:8080/v1/transfer")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("httpCode:%d \n",httpCode)
	fmt.Println(result)
}
