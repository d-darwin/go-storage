package main

import (
	"fmt"

	"github.com/d-darwin/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello, storage!", st)
}
