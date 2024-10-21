package main

import "fmt"

const (
	age        int = 30 // privada para este paquete
	MagicNumer int = 33 // posibilidad de usar fuera del paquete
)

func main() {
	fmt.Println(age, MagicNumer)
}
