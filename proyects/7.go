/* Nicolas Ramirez
comic book store management software proyects */

package main

import "fmt"

func main(){
	
	// we are going to declare different variables that would be used in a comic book categorization
	var  age, pageNumber, cost int32;
	var title, author, artist, publisher string;
	var grade float32;

	title = "100 AÑOS DE SOLEDAD";
	author = "Gabriel Garcia Marquez";
	artist = "NONE";
	publisher = "NORMA";
	age = 1982;
	pageNumber = 1000;
	grade = 10.0;

	fmt.Println(title, "by ", author, "drawn by ", artist, "published by ", publisher)
	fmt.Println("Year", age, "pages ",pageNumber,  "rated ", grade)
	fmt.Println("price", cost)

}
