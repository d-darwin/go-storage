package main

import (
	"fmt"
	"log"

	"github.com/d-darwin/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("Hello, storage!"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded", file)

	storedFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it stored", storedFile)

}
