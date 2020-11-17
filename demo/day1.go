package main1

import "fmt"

// 多行申明
type User struct {
	Username string
	Email    string
}

func main1() {
	// 多行初始化
	u := User{
		Username: "王二麻子",
		Email:    "astaxie@gmail.com",
	}

	fmt.Println(u.Username)
}
