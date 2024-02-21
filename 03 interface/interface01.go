package main

//接口练习
import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

type KeyBoard struct {
	Name string
}

type Computer struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "手机已接入")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "手机已断开连接")
}

func (p Phone) call() {
	fmt.Println(p.Name, "打电话")
}

func (k KeyBoard) start() {
	fmt.Println(k.Name, "键盘已接入")
}

func (k KeyBoard) stop() {
	fmt.Println(k.Name, "键盘已断开连接")
}

func (c Computer) working(usb Usb) {
	usb.start()
	if phone, ok := usb.(Phone); ok {
		phone.call()
	}
	usb.stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = KeyBoard{"罗技"}
	usbArr[2] = Phone{"华为"}

	var computer Computer
	for _, usb := range usbArr {
		computer.working(usb)
	}
	//computer.working(&phone)
	//computer.working(&keyBoard)
}
