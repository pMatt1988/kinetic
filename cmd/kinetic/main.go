package main

import (
	"github.com/pmatt1988/kinetic/internal/db"
	"github.com/pmatt1988/kinetic/internal/router"
)

func main() {
	db.Init()
	router.Init()
}
