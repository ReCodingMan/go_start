package main
import (
	"fmt"
)

//定义一个结构体 Account
type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//存款
func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}

	if money <= 0 {
		fmt.Println("输入的金额不正确")
	}

	account.Balance += money
	fmt.Println("存款成功~")
}

//取款
func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}

	if money <= 0 || money >= account.Balance {
		fmt.Println("输入的金额不正确")
	}

	account.Balance -= money
	fmt.Println("取款成功~")
}

//查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("输入的密码不正确")
		return
	}

	fmt.Printf("账号为 %v 余额为 %v", account.AccountNo, account.Balance)
}

func main() {
	account := Account{
		AccountNo : "工商银行",
		Pwd : "66666",
		Balance : 100.00,
	}

	account.Query("66666")
}