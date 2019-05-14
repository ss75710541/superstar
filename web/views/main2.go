package main

import (
	"superstar/bootstrap"
	"superstar/web/middleware/identity"
	"superstar/web/routes"
)

func main(){
	app := bootstrap.New("Superstar database", "liujinye")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}
