package main

import (
	"fmt"
	"go-to-do-app/app/controllers"
	"go-to-do-app/app/models"
)

func main() {
	fmt.Println(models.Db) //initを呼び出す

	controllers.StartMainServer()
}
