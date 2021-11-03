/* Nicolas Ramirez
Big example
*/


package main

import "fmt"

func calculando(){
	fmt.Println("Calculando Resultado");
}

func calculado(){
	fmt.Println("Resultado Calculado");
}

func calcular(opcion int, lado int) int{
	if opcion==1{
		calculando();
		defer calculado();
		lado = lado * 4;
		return lado;
	} else {
		calculando();
		defer calculado();
		lado = lado * lado;
		return lado;
	}
}


func main(){
	opcion := 0;
	fmt.Println("1=Perimetro, otro numero=area");
	fmt.Scanln(&opcion);
	entrada := 0;
	fmt.Println("ingrese lado");
	fmt.Scanln(&entrada);
	resultado := calcular(opcion, entrada);
	fmt.Println(resultado);
}
