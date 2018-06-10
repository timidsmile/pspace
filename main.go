package main

import (
	"github.com/timidsmile/pspace/initial"
	"github.com/timidsmile/pspace/router"
)

func main() {

	if err := initial.InitDb(); err != nil {
		panic(err)
	}

	if err := initial.InitRedis(); err != nil {
		panic(err)
	}

	router := router.LoadRouters()

	router.Run() // listen and serve on 0.0.0.0:8080
}
