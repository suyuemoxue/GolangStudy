package main

import (
	"GoStudy/customerManage/model"
	"GoStudy/customerManage/service"
	"fmt"
)

type CustomerView struct {
	key             string
	flag            bool
	customerService *service.CustomerService
}

// addCustomer 方法功能1，添加客户
func (cv *CustomerView) addCustomer() {
	fmt.Println("--------------添加客户--------------")
	fmt.Println("编号:")
	id := 0
	fmt.Scanln(&id)
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机号:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(id, name, gender, age, phone, email)
	if cv.customerService.AddCustomer(customer) {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}
	fmt.Println()
}

// 功能2，修改客户
func (cv *CustomerView) updateCustomer() {
	fmt.Println("请输入要修改的用户的ID")
	id := -1
	fmt.Scanln(&id)
	if cv.customerService.UpdateCustomer(id) {
		fmt.Println("修改成功")
	}
}

// deleteCustomer 方法，功能3，删除客户
func (cv *CustomerView) deleteCustomer() {
	fmt.Println("请输入要删除用户的ID")
	id := -1
	fmt.Scanln(&id)
	if cv.customerService.DeleteCustomer(id) {
		fmt.Println("删除成功")
	}
}

// CustomerList 方法，功能4，客户列表
func (cv *CustomerView) customerList() {
	customer := cv.customerService.CustomerList()
	fmt.Println("------------------------------------客户列表------------------------------------")
	fmt.Println("ID\t用户名\t性别\t年龄\t电话\t\t邮件")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println()
}

// mainMenu 方法，主菜单
func (cv *CustomerView) mainMenu() {
	for cv.flag {
		fmt.Println("---------------客户信息管理系统----------------")
		fmt.Println("                1 添加客户")
		fmt.Println("                2 修改客户")
		fmt.Println("                3 删除客户")
		fmt.Println("                4 客户列表")
		fmt.Println("                5 退   出")
		fmt.Println("           请根据需求选择对应序号")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.addCustomer()
		case "2":
			cv.updateCustomer()
		case "3":
			cv.deleteCustomer()
		case "4":
			cv.customerList()
		case "5":
			fmt.Println()
			cv.flag = false
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}

func main() {
	customer := CustomerView{
		key:             "",
		flag:            true,
		customerService: service.NewCustomerService(),
	}
	customer.mainMenu()
}
