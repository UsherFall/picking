package main

import "picking/router"

func main() {
	r := router.InitUserRouter()
	r.Run(":8000")
}
