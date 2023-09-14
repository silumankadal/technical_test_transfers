package main

import (
	"go-jwt/router"
	"go-jwt/database"
)

func main(){
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}