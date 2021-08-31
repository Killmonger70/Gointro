/* Nicolas Ramirez
Sprintln allows us format and concatenate multiple variables, without inmideately printing it
*/

package main

import "fmt"

func main(){
	var string1 string = "hello";
	string2 := "world";
	salute := fmt.Sprintln(string1, string2);

	fmt.Println(salute);


}
