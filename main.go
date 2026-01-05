package main

import (
	"fmt"
	"platform-game/model"
)

func main() {
	user := model.User{ID: 1, Username: "Jack", Active: true}
	fmt.Println("Usuário criado:")
	fmt.Println(user)
}
