/* Nicolas Ramirez
Switch conditional statement
*/

package main

import "fmt"

func main(){
	var Size int32;
	fmt.Println("INGRESE SU TALLA DE ZAPATO (INCH)");
	fmt.Scan(&Size);
	switch Size{
		case 4:
			fmt.Println("Hallway A2");
		case 5:
			fmt.Println("Hallway B1");
		case 6: 
			fmt.Println("Hallway B2");
		default:
			fmt.Println("NO HAY STOCK");
		}

}
