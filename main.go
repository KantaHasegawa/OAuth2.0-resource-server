package main

import "github.com/KantaHasegawa/OAuth2.0-resource-server/router"

func main() {
	r := router.NewRouter()
	r.Run(":8082")
}
