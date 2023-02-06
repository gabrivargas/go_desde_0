package main

import (
	"fmt" 
)

//CONDICIONALES

var num int
var estado bool
var estado2 bool
func main(){
	estado = true
	// se puede declarar valores dentro del if
	if a, num := 2 , 1 ; num == 1{
		// las variables DECLARADAS en el condicional solo pueden ser usadas adentro del mismo
		fmt.Println("* asignacion de valor a variable dentro del condicional if:",num," y ",a);
	}
	// ERROR fmt.Println(a);

	//una sola condicion no requiere parentesis
	if estado {
		fmt.Println("* variable bool condicional automatico toma: ",estado);
	}
	// para mas de una condicion es requerido los () se niega una declaracion con ! al inicio de la variable
	if num2 :=2 ; (!estado2 && estado){
		fmt.Println("* para mas de una condicion hace falta() ", estado2)
		fmt.Println("* variable declarada dentro de un if con mas de una condicion",num2)

	}
	fmt.Println("-----------------------SWITCHITO---------------------------")
	switch num3:=45 ; num3 {
	case 1:
		fmt.Println(num3)
	case 2:
		fmt.Println(num3)
	case 3:
		fmt.Println(num3)
	default:
		fmt.Println(num3)
	}
	fmt.Println("---------------------END SWITCH-----------------------------")


}