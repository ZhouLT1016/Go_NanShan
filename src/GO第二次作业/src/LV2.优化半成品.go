package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

func 装填(gun chan string) {
	time.Sleep(time.Second)
	gun <- "装填"
	fmt.Println("装填完毕")
}
func 瞄准(gun chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("瞄准成功")
	gun <- "瞄准"

}
func 发射(gun chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println("发射发射")
	gun <- "发射"
}
func listenForKeyPress(keyCh chan<- rune) {
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// 如果是按键事件，将按键字符发送给管道
		if char == 'q' {
			keyCh <- char
		}
	}
}
func main() {
	var gun chan string
	gun = make(chan string, 1)
	go 装填(gun) //预装填
	// 开始监听键盘事件
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer func() {
		// 程序退出前关闭键盘监听
		keyboard.Close()
	}()

	// 创建一个用于接收键盘事件的管道
	keyCh := make(chan rune)

	go listenForKeyPress(keyCh) // 启动按键监听

	// 执行开炮程序，直到监听到键盘按键q
	for {
		x := <-gun
		switch x {
		case "装填":
			go 瞄准(gun)
		case "瞄准":
			go 发射(gun)
		case "发射":
			go 装填(gun)

			select {
			case s := <-keyCh:
				if s == 'q' {
					fmt.Print("停止打炮")
					return
				}
			default:
				// 没有监听到键盘按键事件，继续执行循环
				continue
			}
		}
	}
}
