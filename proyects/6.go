/* Nicolas Ramirez
there are diferent operator with the function of replacing the  var1 = var1 (insert operand) var2
for example += -= /= *= */

package main

import "fmt"

func main(){
	var acumlator int32 = 25;
	var num1 int32 = 3;
	var divided int32 = 5;
	text :=  "result is";

	acumlator *= num1;
	acumlator /= divided;


	fmt.Println(text, acumlator);

}
