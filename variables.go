// Nicolas Ramirez
// basic introduction to variables and data types in golang
// introduccion a las variables y los tipos de datos en golang

package main

import "fmt"

var Nombre string = "Nico"
var Edad int = 17
var Pueblo string

func main() {
	// Nueva forma de referirse a una variable / New way to refer to a variable
	Pueblo := "chia"
	fmt.Println("hola , me llamo", Nombre, ",tengo", Edad, "años y vivo en", Pueblo, ":)")
}
