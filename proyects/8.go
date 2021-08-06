/* Nicolas Ramirez
fmt.Println y fmt.Print se pueden concatenar los datos
fmt.Printf para interpolar(osea en el mismo string se dan los verbos)*/

package main

import "fmt"

func main(){
	animal1 := "dog";
	animal2 := "cat";

	fmt.Println("are you a ", animal1, "or a", animal2, "person");
	fmt.Printf("are you a %v or a %v person", animal1, animal2);

}
