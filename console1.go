// Nicolas Ramirez
// fmt permite usar utilidades de consola como println, similar a como lo hacen librerias de c++
// fmt need two % for a ture percentage use
// Sprintf does the instructions, but doesnt print them
// \n breaks current line
// \t añade un tab

package main

import "fmt"

func main() {
	fmt.Printf("Hello world %t", false, " \n"); // decide if its t/f

	fmt.Printf("hello world %b", 1025); // transform a decimal to a binary
	fmt.Printf("hello world %o", 1025); // transform a decimal to an octal
	fmt.Printf("hello world %x", 1025); // transform a decimal to a hexadecimal


	fmt.Printf("hello world %e", 1025.123456789012345678134567); // transform to scientific notation
	fmt.Printf("hello world %f", 1025.123456789); // transforma decimal clasico
	fmt.Printf("hello world %g", 1025.123456789); // no transforma

	fmt.Printf("hello world %s","Nicolas Ramirez"); // imprime el string
	fmt.Printf("hello world %q","Nicolas Ramirez"); // imprime igual pero con "

	fmt.Printf("hello world %9","Nicolas Ramirez"); // el espacio deseado"
	fmt.Printf("hello world %9.2","1025.123456789"); // aproximacion after ."
}
