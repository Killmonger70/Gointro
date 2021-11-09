/* Nicolas Ramirez
interstellar agency proyect
*/

package main

import "fmt"

func CombustibleRestante(fuel int) {
	fmt.Println("You have", fuel, "Gallons of fuel remaining");
}

func calculateFuel(planet string) int{
	var fuel int;
	switch planet{
	case "Venus":
		fuel = 300000;
	case "Mercury":
		fuel = 500000;
	case "Mars":
		fuel = 700000;
	default:
		fuel = 0;
	}
	return fuel;
}

func greetPlanet(planet string){
	fmt.Println("Welcome to planet", planet);
}

func cantFly(){
	fmt.Println("We do not have the available fuel to fly there.");
}

func main(){
	

}
