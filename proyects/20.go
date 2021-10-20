/* Nicolas Ramirez
return values with functions
*/

package main

import "fmt"

func Calcular(entrada int32) int32{ //first, we declare in brackets the variables and its data types that we are going to use for the code to be executed; then, we declare the data type to be returne
	var z int32 = 0;
	z = entrada * 5;
	return z;                  //this keyword is for the data that will be returned
}

func main(){
	var x int32 = 8;
	y := Calcular(x);          //We call the function, Then we use a data as an argument for the function inside the brackets
	fmt.Println(y);        
}
