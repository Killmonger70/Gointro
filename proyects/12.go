/* 
User input with scan in the fmt package
In this case we can only recieve the first ection of the answer, it will recieve everything until the user 
input a space character
*/


package main

import "fmt"

func main(){
	var Input string;
	fmt.Println("Tell me something(use only one word):");
	//scan expect an adress when we use it, thats why we need the & char before a variable
	fmt.Scan(&Input);
	output := fmt.Sprintf("did you say: %v. ?", Input);
	fmt.Println(output);

}
