/*package main

import (
	"fmt"
	//"github.com/eiannone/keyboard"
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
	time.Sleep(time.Second * 4)
	fmt.Println("发射发射")
	gun <- "发射"

}

func main() {

	var gun chan string
	gun = make(chan string, 1)
	go 装填(gun) //预装填
	var s string
	for {
		fmt.Scanln(&s)
		switch s {
		case "":
			x := <-gun
			switch x {
			case "装填":
				go 瞄准(gun)
			case "瞄准":
				go 发射(gun)
			case "发射":
				go 装填(gun)

			}
		case "q":
			fmt.Println("结束")
			return
		}
	}

}*/