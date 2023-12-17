package main

import (
	"Go-Todo/app/models"
	"fmt"
)


func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	//  fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// u.CreateUser()
	// user, _ := models.GetUser(3)
	// user.CreateTodo("test")

	// t, _ :=models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("test2")

	// todos, _ := models.GetTodos()
	// fmt.Println(todos)
	// for _ , v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(3)
	fmt.Println(t)
	t.Content="update test" 
	(&t).UpdateTodo()
	t, _ = models.GetTodo(3)
	fmt.Println(t)
} 