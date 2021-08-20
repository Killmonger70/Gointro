/* Nicolas Ramirez
more fmt utilities like %v %T  %d 
*/

package main

import "fmt"

func main(){
	// %T is used so the compiler tells us what data type a variable is
	numeros := 1234;
	letras :=  "abcd";
	fmt.Printf("the data types are %T and %T", numeros, letras)

}
