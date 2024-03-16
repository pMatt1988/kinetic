package main

import (
	"fmt"

	"github.com/pmatt1988/kinetic/internal/db"
	"github.com/pmatt1988/kinetic/internal/router"
)

func main() {
	fmt.Println("Starting Kinetic...")
	db.Init()
	router.Init()
}
