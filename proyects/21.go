/* Nicolas Ramirez
functions arguments and parameters
*/

package main

import "fmt"

func multiplicar(lol1 int32, lol2 int32) int32{
	lol3 := lol1 * lol2;
	return lol3;
}


func main(){
	var num1 int32;
	fmt.Println("Ingrese el primer numero");
	fmt.Scan(&num1);
	var num2 int32;
	fmt.Println("Ingrese el segundo numero");
	fmt.Scan(&num2);
	resultado := multiplicar(num1, num2)
	fmt.Println(resultado);

}
