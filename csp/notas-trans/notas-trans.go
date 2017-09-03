package main

import (
	"fmt"
)

func conversor( notas []string , done chan int){
	for index, i := range notas {
		switch (i){
			case "Do" : notas[index] = "C"
			case "Re" : notas[index] = "D"
			case "Mi" : notas[index] = "E"
			case "Fa" : notas[index] = "F"
			case "Sol" : notas[index] = "G"
			case "La" : notas[index] = "A"
			case "Si" : notas[index] = "B"
		}
	}
	
	done <- 1
}

func my_main(notas []string, paralel_number int){
	if(paralel_number > len(notas)){ 
		fmt.Println("[ERRO] : Você não pode paralelizar com um valor maior do que a quantidade de elementos da sua entrada!!")
	} else {  
		done := make( chan int )
		for i := 0 ; i < paralel_number; i++{
			aux_i := (i)*(len(notas)/paralel_number)
			aux_j := (i+1)*(len(notas)/paralel_number) 
			if(aux_j == len(notas) -1 ){ aux_j++}
			//fmt.Println(  aux_i, aux_j  )
			go conversor(notas[aux_i:aux_j], done)
		}
		i := 0
		for{
			if(i == paralel_number){ break }
			<- done
			i++
		}
	}
}


func main() {
	notas := [7]string{"Sol","Fa","Sol","Do","Re","Mi","Mi"}
	paralel_number := 5
	
	fmt.Println(notas)
	
	my_main(notas[:], paralel_number)
	
	fmt.Println(notas)
	fmt.Println("FIM!!")
}

