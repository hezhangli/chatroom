package main

import "fmt"

// 定义用户名和密码
var userId int
var password string

func main() {

	//输入的数字
	var num int
	bol := true

	for bol {

		fmt.Println("冒牌qq欢迎你")
		fmt.Println("\t\t\t1.登录聊天室")
		fmt.Println("\t\t\t2.注册用户")
		fmt.Println("\t\t\t3.退出系统")
		fmt.Println("请输入你的选择（1-3）")

		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			fmt.Println("登录聊天室")
			bol = false
		case 2:
			fmt.Println("注册用户")
			bol = false
		case 3:
			fmt.Println("退出系统")
			bol = false
		default:
			fmt.Println("输入错误请重新输入")
		}
	}

	//判断其登录后的步骤
	if num == 1 {
		//输入用户名和密码
		fmt.Println("请输入userId：")
		fmt.Scanf("%d\n", &userId)

		fmt.Println("请输入password：")
		fmt.Scanf("%s\n", &password)

		err := login(userId, password)
		if err != nil {
			fmt.Println("用户名或密码错误， 登录失败")
		}

		fmt.Println("登录成功 ")
	}
}
