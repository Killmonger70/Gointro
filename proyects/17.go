/* Nicolas Ramirez
Randomizing with go, usoing the rand and math utilities
*/

package main

import ("fmt"
	"math/rand")

func main(){
	random := rand.Intn(13098);  // rand will randomize a value from the Intn method, with its upper limit being defined by the creator. It will randomize the same number each time the runs.
	fmt.Println("The number is: ", random);

}
