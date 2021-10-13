/* Nicolas Ramirez
Random heist simulator
*/
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano());
  var isHeistOn bool = true;
  var eludedGuards int = rand.Intn(100);
if (eludedGuards>=50){
  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.");
}  else {
  isHeistOn = false;
  fmt.Println("Plan a better disguise next time?");
}
  var openedVault int = rand.Intn(100);
if (isHeistOn && (openedVault >= 70)){
  fmt.Println("Grab and GO!");
} else if (isHeistOn && (openedVault <= 70)) {
  isHeistOn = false;
  fmt.Println("the vault can't be opened");
}
leftSafely := rand.Intn(5);
if isHeistOn{
  switch leftSafely{
    case 0: 
    isHeistOn = false;
    fmt.Println("Heist Failed");
   case 1: 
    isHeistOn = false;
    fmt.Println("Heist Failed, they got us on camera");   
  case 2: 
    isHeistOn = false;
    fmt.Println("Heist Failed, guns dont fire from that side");
  case 3: 
    isHeistOn = false;
    fmt.Println("Heist Failed, we were supposed to take money OUT");
  default: 
    fmt.Println("Start the getaway car!");
  }
}
if isHeistOn {
  amtStolen := 10000 + rand.Intn(1000000);
  fmt.Println(amtStolen);
}
fmt.Println(isHeistOn);
}
