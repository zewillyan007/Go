package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Escreva usa cidade: ")
	city, _ := reader.ReadString('\n')
	fmt.Print("Você vive em " + city)
}
