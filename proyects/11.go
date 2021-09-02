/* Nicolas Ramirez
Sprintf is like Sprint with added verbose support
*/


package main

import "fmt"

func main(){
	var pet string = "dog";
	text := "i wish i had a %v";

	Example := fmt.Sprintf(text, pet);
	fmt.Println(Example);

}
