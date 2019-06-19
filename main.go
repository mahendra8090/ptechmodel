package main

import (
	"fmt"
	"newgo/test/model"
	server "newgo/test/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// user profile

func main() {

	fmt.Print("in main")
	model.Init()
	server.StartServer()
}
