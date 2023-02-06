package main
import (
	"fmt"
	"os"
	"bufio"

)
//MOSTRAR Y ACEPTAR DATOS

var numero1 int
var numero2 int
var resultado int
var leyenda string


func main(){
	fmt.Println("ingrese numero1 :");
	fmt.Scanf("%d",&numero1);

	fmt.Println("ingrese numero 2 :");
	fmt.Scanf("%d",&numero2);
	fmt.Println(numero1+numero2);
	//
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		leyenda = scanner.Text()
		fmt.Println(scanner);
		fmt.Println(leyenda);
	}
}