package main

import (
	"github.com/pmatt1988/kinetic/lib/db"
	"github.com/pmatt1988/kinetic/lib/router"
)

func main() {
	db.Init()
	router.Init()
}
