package main

import (
	"fmt"
	"g2e-orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `g2e-orm:"primary key"`
	Age  int
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/gee?charset=utf8mb4&parseTime=True&loc=Local"
	engine, _ := g2e_orm.NewEngine("mysql", dsn)
	defer engine.Close()
	s := engine.NewSession()
	s = s.Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		fmt.Println("Failed to create table User")
	}
}
