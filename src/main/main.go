package main

import (
	"main/database"
	"main/route"
)

func main() {
	database.Init()
	r := route.CreateRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}

// main.go -> main.a
// util.go -> util.a
// router.go -> router.a
