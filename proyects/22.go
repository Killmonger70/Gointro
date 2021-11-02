/* Nicolas Ramirez
Defer in golang
when we want to call a function at the end of the execution of another funtion, we use the keyword defer
*/

package main

import "fmt"

func calcularimpuestos(x int) int{
	defer calculado();
	verificar();
	x = (x * 20)/100;
	return x;
}

func verificar(){
	fmt.Println("verificando impuestos");
}

func calculado(){
	fmt.Println("impuestos calculados");
}

func main(){
	var imptuestos int = 32;
	resultado := calcularimpuestos(imptuestos);
	fmt.Println(resultado)
}
