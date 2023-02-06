package main
import (
	"fmt"
)
//CICLOS

func main(){
	i:=1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	} 

		numero:= 0
	for	{
		fmt.Println("input number:");
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var f = 0
	for f < 10 {
		fmt.Printf("\n valor de f: %d",f)
		if f == 5{
			fmt.Printf("multiplicamos por 2 \n")
			f=f*2
			continue
		}
		fmt.Printf("		Pasó por aqui \n")
		f++
	}

	var d int = 0
	// marca o etiqueta inocua e indeiferente en la ejecucion
	RUTINA:
	for d < 20 {
		fmt.Println("-----------------inicio loop-----------------")
		if d == 4 {
			d = d+2
			fmt.Println("voy a rutina sumando dos a d");
			// como el continue pero no vuelve al ciclo anterior al acutal sino que vuelve al punto indicado
			goto RUTINA
		}
		if d == 10 {
			fmt.Println("suma uno a d para no quedar en loop eterno, no llega al final del loop, salta a la siguiente iteracion");
			d++
			continue
		}
		if d == 17 {
			fmt.Println("break detiene la ejecucion del loop");
			break
		}
		fmt.Println("valor de d: ",d, "a sumar +1");
		d++
	}

}
