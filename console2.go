// Nicolas Ramirez
// fmt permite usar utilidades de consola como println, similar a como lo hacen librerias de c++

package main

import (
	"fmt"
	"bufio"
	"os"
	//"strconv"
	)

func main() {
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	input := scanner.Text();
	fmt.Printf("ingresaste: \n", input);
}
