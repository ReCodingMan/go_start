package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {

}

func (p Phone) start()  {
	fmt.Println("手机开始了")
}

func (p Phone) stop()  {
	fmt.Println("手机停止了")
}

type Camera struct {
	name string
}

func (c Camera) start()  {
	fmt.Println("相机开始了")
}

func (c Camera) stop()  {
	fmt.Println("相机停止了")
}

type Computer struct {

}

//这个方法接受usb接口类型
func (c Computer) Working(usb Usb)  {
	usb.start()
	usb.stop()
}

func main() {
	//定义一个usb接口数组，可以存放phone和camera结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{}
	usbArr[1] = Camera{"xiaomi"}
	usbArr[2] = Phone{}
	fmt.Println(usbArr)
}