package main

import (
	"frontendmod/env"
	"frontendmod/routes"
	"frontendmod/util"
	"log"
	"net/http"
)

func main() {
	util.ExecTemplates()

	muxRouter := routes.MyRouter()

	http.Handle("/", muxRouter)

	log.Println("Client running on " + env.FRONT_END_HOST_URL)
	log.Fatal(http.ListenAndServe(":"+env.PORT, nil))
}
