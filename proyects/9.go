/* Nicolas Ramirez
more fmt utilities like %v %T %f 
*/

package main

import "fmt"

func main(){
	// %T is used so the compiler tells us what data type a variable is
	numeros := 1234;
	letras :=  "abcd";
	// %f is used so the compiler aproximates the number of decimals in a float
	aprox := 12.03445;
	fmt.Printf("the data types are %T and %T \n", numeros, letras)
	fmt.Printf("aprox by .2 is %.2f \n", aprox) 
}
