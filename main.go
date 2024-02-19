package main

import (
	"fmt"
	"log"
)

func main() {
	var n float32 = 0
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %f\n", n)
}
