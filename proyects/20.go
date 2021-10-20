/* Nicolas Ramirez
return values with functions
*/

package main

import "fmt"

func Calcular(entrada int32) int32{
	var z int32 = 0;
	z = entrada * 5;
	return z;
}

func main(){
	var x int32 = 8;
	y := Calcular(x);
	fmt.Println(y);
}
