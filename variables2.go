// Nicolas Ramirez
// different methods for defining variables
// there are default values for each data type when it is udeclared but not defined

package main

import "fmt"

var Name string ="Nico"; //EXPLICIT VARIABLES
var Number = 1234; //implicit variables
//Grado := 7; this type of declaration can only be used inside a function


func main() {
	Grado := 7; // implicit variable declaration
	fmt.Printf("%T", Name);
	fmt.Printf("%T", Number);
	fmt.Printf("%T", Grado);
}
