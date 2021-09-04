/* Nicolas Ramirez
conditional statements in golang
*/

package main

import "fmt"
 
func main(){
	var var1 int = 2;
	var var2 int = 3;
	var var3 int = var2 + var1;
	if  (var3 == 5){
		fmt.Println("its 5");
	} else {
		fmt.Println("it aint 5 chief");
	}
}
