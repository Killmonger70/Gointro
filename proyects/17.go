/* Nicolas Ramirez
Randomizing with go, usoing the rand and math utilities
*/

package main

import ("fmt"
	"math/rand"
	"time")

func main(){
	random := rand.Intn(13098);  // rand will randomize a value from the Intn method, with its upper limit being defined by the creator. It will randomize the same number each time the runs.
	fmt.Println("The number is: ", random);

	//this behavior is caused in the way that go seeds, or choose a starting number (default=1)
	// This can be solvented by using a random changing seed, for example, time.
	rand.Seed(time.Now().UnixNano()); //Use a random seed, with the parameter time function now, and using the unix time schedule
	fmt.Println(rand.Intn(13098));

}
