package main

import (
	"fmt"

)
/*
 == igual a
 != diferente
 < menor que
 > mayor que
 >= mayor o igual que
 <= menor o igual que
 && AND concatena condiciones y da verdadera cuando todo es verdadero
 || OR concatena condiciones y da verdadero cuando uno solo es verdadero

*/
func main()  {
	x := 10
	y := 10
	if x > y {
		fmt.Printf("%d es mayor que %d\n", x,y)
	} else if x < y {
		fmt.Printf("%d es mayor que %d\n", y,x)
	} else {
		fmt.Println("Son iguales")
	}
}