/* Nicolas Ramirez
pointers in golang
pointers are variables especifically designed to store addresses
*/

package main

import "fmt"

func main(){
	var pointer *int  //with the * operator, we declare that we are going to store an address
			 // then we write the data type of said address
	direccion := 18000000;
	pointer = &direccion;//pointer being assigned an address
	fmt.Println(pointer);

}
