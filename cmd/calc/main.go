package main

import (
	"bufio"
	"fmt"
	"github.com/zapirus/calc/internal/service"
	"os"
	"strings"
)

func main() {
	fmt.Print("Вводи: ")
	myScn := bufio.NewScanner(os.Stdin)
	myScn.Scan()
	input := myScn.Text()

	result := strings.Split(input, " ")
	validate := service.Validate(result)
	if validate {
		fmt.Println(service.Arithmetic(result))
	}

}
