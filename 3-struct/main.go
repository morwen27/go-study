package main

import (
	"fmt"
	"work-with-bins/file"
	"work-with-bins/storage"
)

func main() {
	storage := storage.Init(file.Init("binlist.json"))
	fmt.Println(storage)
}
