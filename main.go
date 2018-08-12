package main

import (
	"flag"

	"github.com/joaopedrosgs/OpenLoU/app"
)

func main() {
	addrPtr := flag.String("addr", ":8080", "Address")

	flag.Parse()
	app.Run(*addrPtr)
}
