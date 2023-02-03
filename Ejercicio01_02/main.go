package main

import (
	"fmt" 
	"strconv"
)

//varianle publicas visible a otros paquetes
var Numero int
//varianles locales privativas
var texto string
var status bool = true

func main(){
	fmt.Println("* Variable status declara por fuera y cnovocada dentro de funcion: ",status)
	//forma 1 de declara variable por defecto numero esta en 0, string en "" y boolean en false
	var num1 int
	fmt.Println("* Valor por defecto de una variable entera: ",num1)
	//forma 2 de declara variable y reasignar valor
	num2 := 0
	num2 = 13
	fmt.Println("* Valor de una variable reasignada: ",num2)

	//declaracion masiva de variables
	num4, num5, num6 , text , status := 1, 2 ,3 , "hello", false;
	fmt.Println("--------------------------------------------------")
	fmt.Println("* Valores de variables masivanente declaradas y asignadas")
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(text)
	//toma la variable local por encima de la exterior
	fmt.Println("--------------------------------------------------")
	fmt.Println("* variable local con el mismo nombre de la exterior",status)
	showStatus()

	//lenguaje estrictamente tipado
	var num7 int64 = 1;
	//cannot use num7 (variable of type int64) as type int in assignment
	//Error: num2 = num7;
	num2 = int(num7) //funcion de casteo
	//Error: texto = string(num7)
	texto = fmt.Sprint(num7)
	texto = fmt.Sprintf("%d",num7)
	fmt.Println("* variable int64 convertida a string",texto)
	texto = strconv.Itoa(int(num2))
	fmt.Println("* variable int64 convertida a string con otra libreria",texto)

}
//funcion local no puede ser accedida por otros paquetes debido a que empieza por minuscula
func showStatus(){
	//toma la variable del paquete
	fmt.Println("* variable declarada en todo el paquete",status)
}