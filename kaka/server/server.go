package main

import (
	"github.com/scottkiss/kaca"
)

func main() {
	kaca.ServeWs(":8080", true)
}
