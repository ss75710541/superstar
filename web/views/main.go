package main

import (
	"superstar/bootstrap"
	"superstar/web/middleware/identity"
	"superstar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "liujnye")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
