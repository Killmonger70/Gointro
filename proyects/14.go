/* Nicolas Ramirez
user input if else example
*/

package main

import "fmt"

func main(){
	var edad int;
	fmt.Println("cual es la edad");
	//to scan data we use Scan, but it can only scan anythinng before the user inputs a space char
	// also, we need to use a & exactly before the variable that we want to asign the value
	fmt.Scan(&edad);
	if (edad >= 18){
		fmt.Println("you can watch the movie");
	} else {
		fmt.Println("zoomer");
	}

}
