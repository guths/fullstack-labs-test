package main

import (
	"battle-of-monsters/app/config"
	"battle-of-monsters/app/db"
)

func main() {
	config.Load()
	db.Connect()
}
