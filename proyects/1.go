// Nicolas Ramirez
// nueva manera de usar comentarios con /* */
/*
fmt permite usar utilidades de consola como println, similar a como lo hacen librerias de c++
we are going to add the time package
se puede hacer un alias para el uso de los packetes importados
*/

package main

import (
	p1 "fmt"
	t "time"
)

func main() {
	p1.Println(t.Now())
}
