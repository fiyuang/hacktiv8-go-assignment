package main

import (
	"secondAssignment/routers"
)

var PORT = ":3333"

func main() {
	routers.StartServer().Run(PORT)
}
