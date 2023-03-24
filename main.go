package main

import "book_Gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)

}
