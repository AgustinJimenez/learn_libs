package main

import (
	"./controllers"
)

func main() {
	user_controller := controllers.UserController{}
	user_controller.Create()
}
