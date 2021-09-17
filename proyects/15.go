/* Nicolas Ramirez
Else if conditional statement
*/

package main

import "fmt"

func main(){
	var Collection int = 300;

	if (Collection  < 2){
		fmt.Println("wow, they must be really valuable to you.");
	} else if (Collection == 0) {
		fmt.Println("Welcome to the hobby");
	} else {
		fmt.Println("What a nice collection you have");
	}

}
