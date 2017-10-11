package main

import (
	"main/route"
)

func main() {
	r := route.CreateRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}