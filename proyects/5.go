/* Nicolas Ramirez
string data type in golang */

package main

import "fmt"

func main(){
	var nameofsong string = "kaikai kitan";
	var nameofartist string = "eve";
	var songdescription string;

	songdescription = nameofsong + " performed by: " + nameofartist;
	fmt.Println(songdescription);

}
