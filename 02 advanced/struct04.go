package main

import "fmt"

// Account 结构体 模拟银行账户
type Account struct {
	AccountNo string  // 账户编号
	Password  string  // 账户密码
	Balance   float64 //账户余额
}

// Query 方法 查询余额
func (account *Account) Query(password string) {
	if password != account.Password { // 判断密码是否正确
		fmt.Println("密码错误")
		return
	}
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~")
	fmt.Println("当前账户:", account.AccountNo)
	fmt.Println("当前余额:", account.Balance)
	fmt.Println("~~~~~~~~~~~~~~~~~")
}

// Deposit 方法 存款
func (account *Account) Deposit(money float64, password string) {
	if password != account.Password { // 判断密码是否正确
		fmt.Println("密码错误")
		return
	}
	if money <= 0 { // 判断金额是否正确
		fmt.Println("金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
	//fmt.Println("当前余额为:", account.Balance)
	account.Query(password) // 调用Query方法查询余额
}

// WithDraw 方法 取款
func (account *Account) WithDraw(money float64, password string) {
	if password != account.Password { // 判断密码是否正确
		fmt.Println("密码错误")
		return
	}
	if money <= 0 || money > account.Balance { // 判断余额是否充足
		fmt.Println("余额不足，无法取款")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
	//fmt.Println("当前余额为:", account.Balance)
	account.Query(password) // 调用Query方法查询余额
}

func main() {
	account01 := &Account{
		AccountNo: "js20240101",
		Password:  "sxj2000",
		Balance:   100.00,
	}
	//account01.Query("sxj2000")
	//account01.Deposit(10000, "sxj2000")
	//account01.WithDraw(3000, "sxj2000")

	var number int

	var password string
	var money float64

	var flag = true
	for flag {
		fmt.Println("--------------------------")
		fmt.Println("请选择需要的服务")
		fmt.Println("1  查询余额")
		fmt.Println("2  存款")
		fmt.Println("3  取款")
		fmt.Println("4  退出")
		fmt.Println("请根据需要的服务选择编号")
		fmt.Println("--------------------------")
		fmt.Scanln(&number)

		switch number {
		case 1:
			fmt.Println("请输入密码")
			fmt.Scanln(&password)
			account01.Query(password)
		case 2:
			fmt.Println("请输入密码")
			fmt.Scanln(&password)
			fmt.Println("请输入存款金额")
			fmt.Scanln(&money)
			account01.Deposit(money, password)
		case 3:
			fmt.Println("请输入密码")
			fmt.Scanln(&password)
			fmt.Println("请输入取款金额")
			fmt.Scanln(&money)
			account01.WithDraw(money, password)
		case 4:
			flag = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
