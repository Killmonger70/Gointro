/* Nicolas Ramirez
adresses in golang are accessed using the & operator next to a variable
*/


package main

import "fmt"

func main(){
	var help string = "necesito ayuda";
	fmt.Println(help); //prints the data inside the variable
	fmt.Println(&help); //prints the adress of the variable

}
