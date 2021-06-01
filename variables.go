// Nicolas Ramirez
// basic introduction to variables and data types in golang
// introduccion a las variables y los tipos de datos en golang
// recomnedable investigar mas sobre los tipos mas especificos de dato para mejor manejo de emoria en programas
// recommended to investigate more of the specific data types for better memory efficient programs

package main

import "fmt"

var Nombre string = "Nico"
var Edad uint8 = 21
var Pueblo string

func main() {
	// Nueva forma de referirse a una variable / New way to refer to a variable
	Pueblo := "chia"
	Edad := Edad - 4
	fmt.Println("hola , me llamo", Nombre, ",tengo", Edad, "años y vivo en", Pueblo, ":)")
}
