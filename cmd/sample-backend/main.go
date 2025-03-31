package main

import (
	"fmt"

	"github.com/sample/temp/api"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("Hey, learn go\n")
	server := api.BuildServer(":8001")
	server.ListenAndServe()
}
