/* Nicolas Ramirez
Switch conditional statement
*/

package main

import "fmt"

func main(){
	var Size int32;
	fmt.Println("INGRESE SU TALLA DE ZAPATO (INCH)");
	fmt.Scan(&Size);
	switch Size{   //we declare the switch statement with switch, and also the variable that we want to use
		case 4: // We use case to declare the code block that will be executed once the condition is fullfilled, we need to declare the variable in the correct way that the data type is defined, followed by : to declare the beginninng of the copde to be executed
			fmt.Println("Hallway A2"); // the code to be executed
		case 5:
			fmt.Println("Hallway B1");
		case 6: 
			fmt.Println("Hallway B2");
		default: // If none of the cases are true, then this will be the one to  be executed
			fmt.Println("NO HAY STOCK");
		}

}
