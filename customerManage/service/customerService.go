package service

import (
	"GolangStudy/customerManage/model"
	"fmt"
)

// CustomerService 结构体，用于实现增删改查操作
type CustomerService struct {
	Customers []model.Customer
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customer := model.NewCustomer(1, "张三", "男", 20, "13882586512", "123@qq.com")
	customerService.Customers = append(customerService.Customers, customer)
	return customerService
}

// FindById 方法根据ID查找客户在切片中的下标
func (c *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.Customers); i++ {
		if c.Customers[i].Id == id {
			index = i
		}
	}
	return index
}

// AddCustomer 方法，功能1，添加客户
func (c *CustomerService) AddCustomer(customer model.Customer) bool {
	c.Customers = append(c.Customers, customer)
	return true
}

// UpdateCustomer 方法，功能2，修改客户
func (c *CustomerService) UpdateCustomer(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	fmt.Println("原信息:", c.Customers[index])
	key := 0
	flag := true
	for flag {
		fmt.Println("请输入要修改的选项")
		fmt.Println("1 姓名")
		fmt.Println("2 性别")
		fmt.Println("3 年龄")
		fmt.Println("4 手机号")
		fmt.Println("5 邮箱")
		fmt.Println("6 退出")
		fmt.Scanln(&key)
		switch key {
		case 1:
			var name string
			fmt.Println("请输入修改后的姓名")
			fmt.Scanln(&name)
			c.Customers[index].Name = name
		case 2:
			var gender string
			fmt.Println("请输入修改后的性别")
			fmt.Scanln(&gender)
			c.Customers[index].Gender = gender
		case 3:
			var age int
			fmt.Println("请输入修改后的年龄")
			fmt.Scanln(&age)
			c.Customers[index].Age = age
		case 4:
			var phone string
			fmt.Println("请输入修改后的手机号")
			fmt.Scanln(&phone)
			c.Customers[index].Phone = phone
		case 5:
			var email string
			fmt.Println("请输入修改后的邮箱")
			fmt.Scanln(&email)
			c.Customers[index].Email = email
		case 6:
			flag = false
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
	return true
}

// DeleteCustomer 方法，功能3，删除客户
func (c *CustomerService) DeleteCustomer(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.Customers = append(c.Customers[0:index], c.Customers[index+1:]...)
	return true
}

// CustomerList 方法，功能4，客户列表
func (c *CustomerService) CustomerList() []model.Customer {
	return c.Customers
}
